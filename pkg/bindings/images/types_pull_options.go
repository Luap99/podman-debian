// Code generated by go generate; DO NOT EDIT.
package images

import (
	"io"
	"net/url"

	"github.com/containers/podman/v5/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *PullOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *PullOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithAllTags set field AllTags to given value
func (o *PullOptions) WithAllTags(value bool) *PullOptions {
	o.AllTags = &value
	return o
}

// GetAllTags returns value of field AllTags
func (o *PullOptions) GetAllTags() bool {
	if o.AllTags == nil {
		var z bool
		return z
	}
	return *o.AllTags
}

// WithArch set field Arch to given value
func (o *PullOptions) WithArch(value string) *PullOptions {
	o.Arch = &value
	return o
}

// GetArch returns value of field Arch
func (o *PullOptions) GetArch() string {
	if o.Arch == nil {
		var z string
		return z
	}
	return *o.Arch
}

// WithAuthfile set field Authfile to given value
func (o *PullOptions) WithAuthfile(value string) *PullOptions {
	o.Authfile = &value
	return o
}

// GetAuthfile returns value of field Authfile
func (o *PullOptions) GetAuthfile() string {
	if o.Authfile == nil {
		var z string
		return z
	}
	return *o.Authfile
}

// WithOS set field OS to given value
func (o *PullOptions) WithOS(value string) *PullOptions {
	o.OS = &value
	return o
}

// GetOS returns value of field OS
func (o *PullOptions) GetOS() string {
	if o.OS == nil {
		var z string
		return z
	}
	return *o.OS
}

// WithPolicy set field Policy to given value
func (o *PullOptions) WithPolicy(value string) *PullOptions {
	o.Policy = &value
	return o
}

// GetPolicy returns value of field Policy
func (o *PullOptions) GetPolicy() string {
	if o.Policy == nil {
		var z string
		return z
	}
	return *o.Policy
}

// WithPassword set field Password to given value
func (o *PullOptions) WithPassword(value string) *PullOptions {
	o.Password = &value
	return o
}

// GetPassword returns value of field Password
func (o *PullOptions) GetPassword() string {
	if o.Password == nil {
		var z string
		return z
	}
	return *o.Password
}

// WithProgressWriter set field ProgressWriter to given value
func (o *PullOptions) WithProgressWriter(value io.Writer) *PullOptions {
	o.ProgressWriter = &value
	return o
}

// GetProgressWriter returns value of field ProgressWriter
func (o *PullOptions) GetProgressWriter() io.Writer {
	if o.ProgressWriter == nil {
		var z io.Writer
		return z
	}
	return *o.ProgressWriter
}

// WithQuiet set field Quiet to given value
func (o *PullOptions) WithQuiet(value bool) *PullOptions {
	o.Quiet = &value
	return o
}

// GetQuiet returns value of field Quiet
func (o *PullOptions) GetQuiet() bool {
	if o.Quiet == nil {
		var z bool
		return z
	}
	return *o.Quiet
}

// WithRetry set field Retry to given value
func (o *PullOptions) WithRetry(value uint) *PullOptions {
	o.Retry = &value
	return o
}

// GetRetry returns value of field Retry
func (o *PullOptions) GetRetry() uint {
	if o.Retry == nil {
		var z uint
		return z
	}
	return *o.Retry
}

// WithRetryDelay set field RetryDelay to given value
func (o *PullOptions) WithRetryDelay(value string) *PullOptions {
	o.RetryDelay = &value
	return o
}

// GetRetryDelay returns value of field RetryDelay
func (o *PullOptions) GetRetryDelay() string {
	if o.RetryDelay == nil {
		var z string
		return z
	}
	return *o.RetryDelay
}

// WithSkipTLSVerify set field SkipTLSVerify to given value
func (o *PullOptions) WithSkipTLSVerify(value bool) *PullOptions {
	o.SkipTLSVerify = &value
	return o
}

// GetSkipTLSVerify returns value of field SkipTLSVerify
func (o *PullOptions) GetSkipTLSVerify() bool {
	if o.SkipTLSVerify == nil {
		var z bool
		return z
	}
	return *o.SkipTLSVerify
}

// WithUsername set field Username to given value
func (o *PullOptions) WithUsername(value string) *PullOptions {
	o.Username = &value
	return o
}

// GetUsername returns value of field Username
func (o *PullOptions) GetUsername() string {
	if o.Username == nil {
		var z string
		return z
	}
	return *o.Username
}

// WithVariant set field Variant to given value
func (o *PullOptions) WithVariant(value string) *PullOptions {
	o.Variant = &value
	return o
}

// GetVariant returns value of field Variant
func (o *PullOptions) GetVariant() string {
	if o.Variant == nil {
		var z string
		return z
	}
	return *o.Variant
}
