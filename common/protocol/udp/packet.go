package udp

import (
	"github.com/stereomonk/xray-core-awg/common/buf"
	"github.com/stereomonk/xray-core-awg/common/net"
)

// Packet is a UDP packet together with its source and destination address.
type Packet struct {
	Payload *buf.Buffer
	Source  net.Destination
	Target  net.Destination
}
