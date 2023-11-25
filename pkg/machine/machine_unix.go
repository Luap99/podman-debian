//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd
// +build darwin dragonfly freebsd linux netbsd openbsd

package machine

import (
	"errors"
	"strings"
)

// ParseVolumeFromPath is a oneshot parsing of a provided volume.  It follows the "rules" of
// the singular parsing functions
func ParseVolumeFromPath(v string) (source, target, options string, readonly bool, err error) {
	split := strings.SplitN(v, ":", 3)
	switch len(split) {
	case 1:
		source = split[0]
		target = split[0]
	case 2:
		source = split[0]
		target = split[1]
	case 3:
		source = split[0]
		target = split[1]
		options = split[2]
	default:
		return "", "", "", false, errors.New("invalid volume provided")
	}

	// I suppose an option not intended for read-only could interfere here but I do not see a better way
	if strings.Contains(options, "ro") {
		readonly = true
	}
	return
}
