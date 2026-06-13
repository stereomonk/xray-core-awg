package outbound

import (

	"github.com/stereomonk/xray-core-awg/common/net"
	"github.com/stereomonk/xray-core-awg/common/protocol"
)

// As a stub command consumer.
func (h *Handler) handleCommand(dest net.Destination, cmd protocol.ResponseCommand) {
	switch cmd.(type) {
	default:
	}
}
