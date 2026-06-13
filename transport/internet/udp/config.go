package udp

import (
	"github.com/stereomonk/xray-core-awg/common"
	"github.com/stereomonk/xray-core-awg/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
