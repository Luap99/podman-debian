// Code generated by go generate; DO NOT EDIT.
package containers

import (
	"net/url"

	"github.com/containers/podman/v5/pkg/bindings/internal/util"
)

// Changed returns true if named field has been set
func (o *CheckpointOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams formats struct fields to be passed to API service
func (o *CheckpointOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}

// WithExport set field Export to given value
func (o *CheckpointOptions) WithExport(value string) *CheckpointOptions {
	o.Export = &value
	return o
}

// GetExport returns value of field Export
func (o *CheckpointOptions) GetExport() string {
	if o.Export == nil {
		var z string
		return z
	}
	return *o.Export
}

// WithCreateImage set field CreateImage to given value
func (o *CheckpointOptions) WithCreateImage(value string) *CheckpointOptions {
	o.CreateImage = &value
	return o
}

// GetCreateImage returns value of field CreateImage
func (o *CheckpointOptions) GetCreateImage() string {
	if o.CreateImage == nil {
		var z string
		return z
	}
	return *o.CreateImage
}

// WithIgnoreRootfs set field IgnoreRootfs to given value
func (o *CheckpointOptions) WithIgnoreRootfs(value bool) *CheckpointOptions {
	o.IgnoreRootfs = &value
	return o
}

// GetIgnoreRootfs returns value of field IgnoreRootfs
func (o *CheckpointOptions) GetIgnoreRootfs() bool {
	if o.IgnoreRootfs == nil {
		var z bool
		return z
	}
	return *o.IgnoreRootfs
}

// WithKeep set field Keep to given value
func (o *CheckpointOptions) WithKeep(value bool) *CheckpointOptions {
	o.Keep = &value
	return o
}

// GetKeep returns value of field Keep
func (o *CheckpointOptions) GetKeep() bool {
	if o.Keep == nil {
		var z bool
		return z
	}
	return *o.Keep
}

// WithLeaveRunning set field LeaveRunning to given value
func (o *CheckpointOptions) WithLeaveRunning(value bool) *CheckpointOptions {
	o.LeaveRunning = &value
	return o
}

// GetLeaveRunning returns value of field LeaveRunning
func (o *CheckpointOptions) GetLeaveRunning() bool {
	if o.LeaveRunning == nil {
		var z bool
		return z
	}
	return *o.LeaveRunning
}

// WithTCPEstablished set field TCPEstablished to given value
func (o *CheckpointOptions) WithTCPEstablished(value bool) *CheckpointOptions {
	o.TCPEstablished = &value
	return o
}

// GetTCPEstablished returns value of field TCPEstablished
func (o *CheckpointOptions) GetTCPEstablished() bool {
	if o.TCPEstablished == nil {
		var z bool
		return z
	}
	return *o.TCPEstablished
}

// WithPrintStats set field PrintStats to given value
func (o *CheckpointOptions) WithPrintStats(value bool) *CheckpointOptions {
	o.PrintStats = &value
	return o
}

// GetPrintStats returns value of field PrintStats
func (o *CheckpointOptions) GetPrintStats() bool {
	if o.PrintStats == nil {
		var z bool
		return z
	}
	return *o.PrintStats
}

// WithPreCheckpoint set field PreCheckpoint to given value
func (o *CheckpointOptions) WithPreCheckpoint(value bool) *CheckpointOptions {
	o.PreCheckpoint = &value
	return o
}

// GetPreCheckpoint returns value of field PreCheckpoint
func (o *CheckpointOptions) GetPreCheckpoint() bool {
	if o.PreCheckpoint == nil {
		var z bool
		return z
	}
	return *o.PreCheckpoint
}

// WithWithPrevious set field WithPrevious to given value
func (o *CheckpointOptions) WithWithPrevious(value bool) *CheckpointOptions {
	o.WithPrevious = &value
	return o
}

// GetWithPrevious returns value of field WithPrevious
func (o *CheckpointOptions) GetWithPrevious() bool {
	if o.WithPrevious == nil {
		var z bool
		return z
	}
	return *o.WithPrevious
}

// WithFileLocks set field FileLocks to given value
func (o *CheckpointOptions) WithFileLocks(value bool) *CheckpointOptions {
	o.FileLocks = &value
	return o
}

// GetFileLocks returns value of field FileLocks
func (o *CheckpointOptions) GetFileLocks() bool {
	if o.FileLocks == nil {
		var z bool
		return z
	}
	return *o.FileLocks
}
