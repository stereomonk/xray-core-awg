//go:build linux
// +build linux

package tcp_test

import (
	"context"
	"strings"
	"testing"

	"github.com/stereomonk/xray-core-awg/common"
	"github.com/stereomonk/xray-core-awg/testing/servers/tcp"
	"github.com/stereomonk/xray-core-awg/transport/internet"
	. "github.com/stereomonk/xray-core-awg/transport/internet/tcp"
)

func TestGetOriginalDestination(t *testing.T) {
	tcpServer := tcp.Server{}
	dest, err := tcpServer.Start()
	common.Must(err)
	defer tcpServer.Close()

	config, err := internet.ToMemoryStreamConfig(nil)
	common.Must(err)
	conn, err := Dial(context.Background(), dest, config)
	common.Must(err)
	defer conn.Close()

	originalDest, err := GetOriginalDestination(conn)
	if !(dest == originalDest || strings.Contains(err.Error(), "failed to call getsockopt")) {
		t.Error("unexpected state")
	}
}
