package dns_test

import (
	"context"
	"testing"
	"time"

	. "github.com/stereomonk/xray-core-awg/app/dns"
	"github.com/stereomonk/xray-core-awg/common"
	"github.com/stereomonk/xray-core-awg/features/dns"
)

func TestLocalNameServer(t *testing.T) {
	s := NewLocalNameServer()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	ips, _, err := s.QueryIP(ctx, "google.com", dns.IPOption{
		IPv4Enable: true,
		IPv6Enable: true,
		FakeEnable: false,
	})
	cancel()
	common.Must(err)
	if len(ips) == 0 {
		t.Error("expect some ips, but got 0")
	}
}
