// Code generated by go generate; DO NOT EDIT.
package kube

import (
	"net/url"

	"github.com/containers/podman/v5/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *ApplyOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *ApplyOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithKubeconfig set field Kubeconfig to given value
func (o *ApplyOptions) WithKubeconfig(value string) *ApplyOptions {
	o.Kubeconfig = &value
	return o
}

// GetKubeconfig returns value of field Kubeconfig
func (o *ApplyOptions) GetKubeconfig() string {
	if o.Kubeconfig == nil {
		var z string
		return z
	}
	return *o.Kubeconfig
}

// WithNamespace set field Namespace to given value
func (o *ApplyOptions) WithNamespace(value string) *ApplyOptions {
	o.Namespace = &value
	return o
}

// GetNamespace returns value of field Namespace
func (o *ApplyOptions) GetNamespace() string {
	if o.Namespace == nil {
		var z string
		return z
	}
	return *o.Namespace
}

// WithCACertFile set field CACertFile to given value
func (o *ApplyOptions) WithCACertFile(value string) *ApplyOptions {
	o.CACertFile = &value
	return o
}

// GetCACertFile returns value of field CACertFile
func (o *ApplyOptions) GetCACertFile() string {
	if o.CACertFile == nil {
		var z string
		return z
	}
	return *o.CACertFile
}

// WithFile set field File to given value
func (o *ApplyOptions) WithFile(value string) *ApplyOptions {
	o.File = &value
	return o
}

// GetFile returns value of field File
func (o *ApplyOptions) GetFile() string {
	if o.File == nil {
		var z string
		return z
	}
	return *o.File
}

// WithService set field Service to given value
func (o *ApplyOptions) WithService(value bool) *ApplyOptions {
	o.Service = &value
	return o
}

// GetService returns value of field Service
func (o *ApplyOptions) GetService() bool {
	if o.Service == nil {
		var z bool
		return z
	}
	return *o.Service
}
