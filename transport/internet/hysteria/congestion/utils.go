package congestion

import (
	"github.com/apernet/quic-go"
	"github.com/stereomonk/xray-core-awg/transport/internet/hysteria/congestion/bbr"
	"github.com/stereomonk/xray-core-awg/transport/internet/hysteria/congestion/brutal"
)

func UseBBR(conn *quic.Conn) {
	conn.SetCongestionControl(bbr.NewBbrSender(
		bbr.DefaultClock{},
		bbr.GetInitialPacketSize(conn.RemoteAddr()),
	))
}

func UseBrutal(conn *quic.Conn, tx uint64) {
	conn.SetCongestionControl(brutal.NewBrutalSender(tx))
}
