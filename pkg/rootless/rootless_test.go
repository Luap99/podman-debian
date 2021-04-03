package rootless

import (
	"reflect"
	"testing"

	"github.com/opencontainers/runc/libcontainer/user"
	spec "github.com/opencontainers/runtime-spec/specs-go"
)

func TestMaybeSplitMappings(t *testing.T) {
	mappings := []spec.LinuxIDMapping{
		{
			ContainerID: 0,
			HostID:      0,
			Size:        2,
		},
	}
	desiredMappings := []spec.LinuxIDMapping{
		{
			ContainerID: 0,
			HostID:      0,
			Size:        1,
		},
		{
			ContainerID: 1,
			HostID:      1,
			Size:        1,
		},
	}
	availableMappings := []user.IDMap{
		{
			ID:       1,
			ParentID: 1000000,
			Count:    65536,
		},
		{
			ID:       0,
			ParentID: 1000,
			Count:    1,
		},
	}
	newMappings := MaybeSplitMappings(mappings, availableMappings)
	if !reflect.DeepEqual(newMappings, desiredMappings) {
		t.Fatal("wrong mappings generated")
	}

	mappings = []spec.LinuxIDMapping{
		{
			ContainerID: 0,
			HostID:      0,
			Size:        2,
		},
	}
	desiredMappings = []spec.LinuxIDMapping{
		{
			ContainerID: 0,
			HostID:      0,
			Size:        2,
		},
	}
	availableMappings = []user.IDMap{
		{
			ID:       0,
			ParentID: 1000000,
			Count:    65536,
		},
	}
	newMappings = MaybeSplitMappings(mappings, availableMappings)

	if !reflect.DeepEqual(newMappings, desiredMappings) {
		t.Fatal("wrong mappings generated")
	}

	mappings = []spec.LinuxIDMapping{
		{
			ContainerID: 0,
			HostID:      0,
			Size:        1,
		},
	}
	desiredMappings = []spec.LinuxIDMapping{
		{
			ContainerID: 0,
			HostID:      0,
			Size:        1,
		},
	}
	availableMappings = []user.IDMap{
		{
			ID:       10000,
			ParentID: 10000,
			Count:    65536,
		},
	}

	newMappings = MaybeSplitMappings(mappings, availableMappings)
	if !reflect.DeepEqual(newMappings, desiredMappings) {
		t.Fatal("wrong mappings generated")
	}
}
