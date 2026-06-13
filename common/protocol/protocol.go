package protocol // import "github.com/stereomonk/xray-core-awg/common/protocol"

import (
	"errors"
)

var ErrProtoNeedMoreData = errors.New("protocol matches, but need more data to complete sniffing")
