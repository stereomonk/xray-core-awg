package conf_test

import (
	"testing"

	"github.com/stereomonk/xray-core-awg/common/net"
	"github.com/stereomonk/xray-core-awg/common/protocol"
	. "github.com/stereomonk/xray-core-awg/infra/conf"
	"github.com/stereomonk/xray-core-awg/proxy/freedom"
	"github.com/stereomonk/xray-core-awg/transport/internet"
)

func TestFreedomConfig(t *testing.T) {
	creator := func() Buildable {
		return new(FreedomConfig)
	}

	runMultiTestCase(t, []TestCase{
		{
			Input: `{
				"domainStrategy": "AsIs",
				"redirect": "127.0.0.1:3366",
				"userLevel": 1
			}`,
			Parser: loadJSON(creator),
			Output: &freedom.Config{
				DomainStrategy: internet.DomainStrategy_AS_IS,
				DestinationOverride: &freedom.DestinationOverride{
					Server: &protocol.ServerEndpoint{
						Address: &net.IPOrDomain{
							Address: &net.IPOrDomain_Ip{
								Ip: []byte{127, 0, 0, 1},
							},
						},
						Port: 3366,
					},
				},
				UserLevel: 1,
			},
		},
	})
}
