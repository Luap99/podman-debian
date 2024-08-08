// Code generated by go generate; DO NOT EDIT.
package images

import (
	"net/url"

	"github.com/containers/podman/v5/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *LoadOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *LoadOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithReference set field Reference to given value
func (o *LoadOptions) WithReference(value string) *LoadOptions {
	o.Reference = &value
	return o
}

// GetReference returns value of field Reference
func (o *LoadOptions) GetReference() string {
	if o.Reference == nil {
		var z string
		return z
	}
	return *o.Reference
}
