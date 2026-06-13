package task

import "github.com/stereomonk/xray-core-awg/common"

// Close returns a func() that closes v.
func Close(v interface{}) func() error {
	return func() error {
		return common.Close(v)
	}
}
