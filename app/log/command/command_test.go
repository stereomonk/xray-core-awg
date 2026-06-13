package command_test

import (
	"context"
	"testing"

	"github.com/stereomonk/xray-core-awg/app/dispatcher"
	"github.com/stereomonk/xray-core-awg/app/log"
	. "github.com/stereomonk/xray-core-awg/app/log/command"
	"github.com/stereomonk/xray-core-awg/app/proxyman"
	_ "github.com/stereomonk/xray-core-awg/app/proxyman/inbound"
	_ "github.com/stereomonk/xray-core-awg/app/proxyman/outbound"
	"github.com/stereomonk/xray-core-awg/common"
	"github.com/stereomonk/xray-core-awg/common/serial"
	"github.com/stereomonk/xray-core-awg/core"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
