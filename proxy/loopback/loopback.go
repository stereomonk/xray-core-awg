package loopback

import (
	"context"

	"github.com/stereomonk/xray-core-awg/common"
	"github.com/stereomonk/xray-core-awg/common/errors"
	"github.com/stereomonk/xray-core-awg/common/session"
	"github.com/stereomonk/xray-core-awg/core"
	"github.com/stereomonk/xray-core-awg/features/routing"
	"github.com/stereomonk/xray-core-awg/transport"
	"github.com/stereomonk/xray-core-awg/transport/internet"
)

type Loopback struct {
	config             *Config
	dispatcherInstance routing.Dispatcher
}

func (l *Loopback) Process(ctx context.Context, link *transport.Link, _ internet.Dialer) error {
	outbounds := session.OutboundsFromContext(ctx)
	ob := outbounds[len(outbounds)-1]
	if !ob.Target.IsValid() {
		return errors.New("target not specified.")
	}
	ob.Name = "loopback"
	destination := ob.Target

	errors.LogInfo(ctx, "opening connection to ", destination)
	content := new(session.Content)
	content.SkipDNSResolve = true

	ctx = session.ContextWithContent(ctx, content)
	inbound := &session.Inbound{}
	originInbound := session.InboundFromContext(ctx)
	if originInbound != nil {
		// get a shallow copy to avoid modifying the inbound tag in upstream context
		*inbound = *originInbound
	}
	inbound.Tag = l.config.InboundTag
	ctx = session.ContextWithInbound(ctx, inbound)

	err := l.dispatcherInstance.DispatchLink(ctx, destination, link)
	if err != nil {
		errors.New(ctx, "failed to process loopback connection").Base(err)
		return err
	}
	return nil
}

func (l *Loopback) init(config *Config, dispatcherInstance routing.Dispatcher) error {
	l.dispatcherInstance = dispatcherInstance
	l.config = config
	return nil
}

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		l := new(Loopback)
		err := core.RequireFeatures(ctx, func(dispatcherInstance routing.Dispatcher) error {
			return l.init(config.(*Config), dispatcherInstance)
		})
		return l, err
	}))
}
