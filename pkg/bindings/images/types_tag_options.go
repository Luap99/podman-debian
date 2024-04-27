// Code generated by go generate; DO NOT EDIT.
package images

import (
	"net/url"

	"github.com/containers/podman/v5/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *TagOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *TagOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}
