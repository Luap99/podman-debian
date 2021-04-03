package compat

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/containers/image/v5/types"
	"github.com/containers/podman/v3/libpod"
	"github.com/containers/podman/v3/pkg/api/handlers/utils"
	"github.com/containers/podman/v3/pkg/auth"
	"github.com/containers/podman/v3/pkg/channel"
	"github.com/containers/podman/v3/pkg/domain/entities"
	"github.com/containers/podman/v3/pkg/domain/infra/abi"
	"github.com/containers/storage"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/gorilla/schema"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// PushImage is the handler for the compat http endpoint for pushing images.
func PushImage(w http.ResponseWriter, r *http.Request) {
	decoder := r.Context().Value("decoder").(*schema.Decoder)
	runtime := r.Context().Value("runtime").(*libpod.Runtime)
	// Now use the ABI implementation to prevent us from having duplicate
	// code.
	imageEngine := abi.ImageEngine{Libpod: runtime}

	digestFile, err := ioutil.TempFile("", "digest.txt")
	if err != nil {
		utils.Error(w, "unable to create digest tempfile", http.StatusInternalServerError, errors.Wrap(err, "unable to create tempfile"))
		return
	}
	defer digestFile.Close()

	// Now use the ABI implementation to prevent us from having duplicate
	// code.
	imageEngine := abi.ImageEngine{Libpod: runtime}

	query := struct {
		All         bool   `schema:"all"`
		Compress    bool   `schema:"compress"`
		Destination string `schema:"destination"`
		Format      string `schema:"format"`
		TLSVerify   bool   `schema:"tlsVerify"`
		Tag         string `schema:"tag"`
	}{
		// This is where you can override the golang default value for one of fields
		TLSVerify: true,
	}

	if err := decoder.Decode(&query, r.URL.Query()); err != nil {
		utils.Error(w, "Something went wrong.", http.StatusBadRequest, errors.Wrapf(err, "failed to parse parameters for %s", r.URL.String()))
		return
	}

	// Note that Docker's docs state "Image name or ID" to be in the path
	// parameter but it really must be a name as Docker does not allow for
	// pushing an image by ID.
	imageName := strings.TrimSuffix(utils.GetName(r), "/push") // GetName returns the entire path
	if query.Tag != "" {
		imageName += ":" + query.Tag
	}
	if _, err := utils.ParseStorageReference(imageName); err != nil {
		utils.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest,
			errors.Wrapf(err, "image source %q is not a containers-storage-transport reference", imageName))
		return
	}

	authconf, authfile, key, err := auth.GetCredentials(r)
	if err != nil {
		utils.Error(w, "Something went wrong.", http.StatusBadRequest, errors.Wrapf(err, "failed to parse %q header for %s", key, r.URL.String()))
		return
	}
	defer auth.RemoveAuthfile(authfile)
	var username, password string
	if authconf != nil {
		username = authconf.Username
		password = authconf.Password
	}
	options := entities.ImagePushOptions{
		All:        query.All,
		Authfile:   authfile,
		Compress:   query.Compress,
		Format:     query.Format,
		Password:   password,
		Username:   username,
		DigestFile: digestFile.Name(),
		Quiet:      true,
		Progress:   make(chan types.ProgressProperties),
	}
	if _, found := r.URL.Query()["tlsVerify"]; found {
		options.SkipTLSVerify = types.NewOptionalBool(!query.TLSVerify)
	}

	var destination string
	if _, found := r.URL.Query()["destination"]; found {
		destination = query.Destination
	} else {
		destination = imageName
	}
	if err := imageEngine.Push(context.Background(), imageName, query.Destination, options); err != nil {
		if errors.Cause(err) != storage.ErrImageUnknown {
			utils.ImageNotFound(w, imageName, errors.Wrapf(err, "failed to find image %s", imageName))
			return
		}

	errorWriter := channel.NewWriter(make(chan []byte))
	defer errorWriter.Close()

	statusWriter := channel.NewWriter(make(chan []byte))
	defer statusWriter.Close()

	runCtx, cancel := context.WithCancel(context.Background())
	var failed bool

	go func() {
		defer cancel()

		statusWriter.Write([]byte(fmt.Sprintf("The push refers to repository [%s]", imageName)))

		err := imageEngine.Push(runCtx, imageName, destination, options)
		if err != nil {
			if errors.Cause(err) != storage.ErrImageUnknown {
				errorWriter.Write([]byte("An image does not exist locally with the tag: " + imageName))
			} else {
				errorWriter.Write([]byte(err.Error()))
			}
		}
	}()

	flush := func() {
		if flusher, ok := w.(http.Flusher); ok {
			flusher.Flush()
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	flush()

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)

loop: // break out of for/select infinite loop
	for {
		var report jsonmessage.JSONMessage

		select {
		case e := <-options.Progress:
			switch e.Event {
			case types.ProgressEventNewArtifact:
				report.Status = "Preparing"
			case types.ProgressEventRead:
				report.Status = "Pushing"
				report.Progress = &jsonmessage.JSONProgress{
					Current: int64(e.Offset),
					Total:   e.Artifact.Size,
				}
			case types.ProgressEventSkipped:
				report.Status = "Layer already exists"
			case types.ProgressEventDone:
				report.Status = "Pushed"
			}
			report.ID = e.Artifact.Digest.Encoded()[0:12]
			if err := enc.Encode(report); err != nil {
				errorWriter.Write([]byte(err.Error()))
			}
			flush()
		case e := <-statusWriter.Chan():
			report.Status = string(e)
			if err := enc.Encode(report); err != nil {
				errorWriter.Write([]byte(err.Error()))
			}
			flush()
		case e := <-errorWriter.Chan():
			failed = true
			report.Error = &jsonmessage.JSONError{
				Message: string(e),
			}
			report.ErrorMessage = string(e)
			if err := enc.Encode(report); err != nil {
				logrus.Warnf("Failed to json encode error %q", err.Error())
			}
			flush()
		case <-runCtx.Done():
			if !failed {
				digestBytes, err := ioutil.ReadAll(digestFile)
				if err == nil {
					tag := query.Tag
					if tag == "" {
						tag = "latest"
					}
					report.Status = fmt.Sprintf("%s: digest: %s", tag, string(digestBytes))
					if err := enc.Encode(report); err != nil {
						logrus.Warnf("Failed to json encode error %q", err.Error())
					}
					flush()
				}
			}
			break loop // break out of for/select infinite loop
		case <-r.Context().Done():
			// Client has closed connection
			break loop // break out of for/select infinite loop
		}
	}
}
