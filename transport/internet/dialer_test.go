package internet_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stereomonk/xray-core-awg/common"
	"github.com/stereomonk/xray-core-awg/common/net"
	"github.com/stereomonk/xray-core-awg/testing/servers/tcp"
	. "github.com/stereomonk/xray-core-awg/transport/internet"
)

func TestDialWithLocalAddr(t *testing.T) {
	server := &tcp.Server{}
	dest, err := server.Start()
	common.Must(err)
	defer server.Close()

	conn, err := DialSystem(context.Background(), net.TCPDestination(net.LocalHostIP, dest.Port), nil)
	common.Must(err)
	if r := cmp.Diff(conn.RemoteAddr().String(), "127.0.0.1:"+dest.Port.String()); r != "" {
		t.Error(r)
	}
	conn.Close()
}
