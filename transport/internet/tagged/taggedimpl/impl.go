package taggedimpl

import (
	"context"

	"github.com/stereomonk/xray-core-awg/common/errors"
	"github.com/stereomonk/xray-core-awg/common/net"
	"github.com/stereomonk/xray-core-awg/common/net/cnc"
	"github.com/stereomonk/xray-core-awg/common/session"
	"github.com/stereomonk/xray-core-awg/core"
	"github.com/stereomonk/xray-core-awg/features/routing"
	"github.com/stereomonk/xray-core-awg/transport/internet/tagged"
)

func DialTaggedOutbound(ctx context.Context, dispatcher routing.Dispatcher, dest net.Destination, tag string) (net.Conn, error) {
	if core.FromContext(ctx) == nil {
		return nil, errors.New("Instance context variable is not in context, dial denied. ")
	}
	content := new(session.Content)
	content.SkipDNSResolve = true

	ctx = session.ContextWithContent(ctx, content)
	ctx = session.SetForcedOutboundTagToContext(ctx, tag)

	r, err := dispatcher.Dispatch(ctx, dest)
	if err != nil {
		return nil, err
	}
	var readerOpt cnc.ConnectionOption
	if dest.Network == net.Network_TCP {
		readerOpt = cnc.ConnectionOutputMulti(r.Reader)
	} else {
		readerOpt = cnc.ConnectionOutputMultiUDP(r.Reader)
	}
	return cnc.NewConnection(cnc.ConnectionInputMulti(r.Writer), readerOpt), nil
}

func init() {
	tagged.Dialer = DialTaggedOutbound
}
