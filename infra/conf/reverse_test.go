package conf_test

import (
	"testing"

	"github.com/stereomonk/xray-core-awg/app/reverse"
	"github.com/stereomonk/xray-core-awg/infra/conf"
)

func TestReverseConfig(t *testing.T) {
	creator := func() conf.Buildable {
		return new(conf.ReverseConfig)
	}

	runMultiTestCase(t, []TestCase{
		{
			Input: `{
				"bridges": [{
					"tag": "test",
					"domain": "test.example.com"
				}]
			}`,
			Parser: loadJSON(creator),
			Output: &reverse.Config{
				BridgeConfig: []*reverse.BridgeConfig{
					{Tag: "test", Domain: "test.example.com"},
				},
			},
		},
		{
			Input: `{
				"portals": [{
					"tag": "test",
					"domain": "test.example.com"
				}]
			}`,
			Parser: loadJSON(creator),
			Output: &reverse.Config{
				PortalConfig: []*reverse.PortalConfig{
					{Tag: "test", Domain: "test.example.com"},
				},
			},
		},
	})
}
