//go:build !windows && !linux

package net

import (
	"github.com/stereomonk/xray-core-awg/common/errors"
)

func FindProcess(dest Destination) (int, string, string, error) {
	return 0, "", "", errors.New("process lookup is not supported on this platform")
}
