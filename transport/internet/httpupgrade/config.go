package httpupgrade

import (
	"github.com/stereomonk/xray-core-awg/common"
	"github.com/stereomonk/xray-core-awg/transport/internet"
)

func (c *Config) GetNormalizedPath() string {
	path := c.Path
	if path == "" {
		return "/"
	}
	if path[0] != '/' {
		return "/" + path
	}
	return path
}

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
