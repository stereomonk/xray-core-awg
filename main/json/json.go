package json

import (
	"context"
	"io"

	"github.com/stereomonk/xray-core-awg/common"
	"github.com/stereomonk/xray-core-awg/common/cmdarg"
	"github.com/stereomonk/xray-core-awg/common/errors"
	"github.com/stereomonk/xray-core-awg/core"
	"github.com/stereomonk/xray-core-awg/infra/conf"
	"github.com/stereomonk/xray-core-awg/infra/conf/serial"
	"github.com/stereomonk/xray-core-awg/main/confloader"
)

func init() {
	common.Must(core.RegisterConfigLoader(&core.ConfigFormat{
		Name:      "JSON",
		Extension: []string{"json"},
		Loader: func(input interface{}) (*core.Config, error) {
			switch v := input.(type) {
			case cmdarg.Arg:
				cf := &conf.Config{}
				for i, arg := range v {
					errors.LogInfo(context.Background(), "Reading config: ", arg)
					r, err := confloader.LoadConfig(arg)
					if err != nil {
						return nil, errors.New("failed to read config: ", arg).Base(err)
					}
					c, err := serial.DecodeJSONConfig(r)
					if err != nil {
						return nil, errors.New("failed to decode config: ", arg).Base(err)
					}
					if i == 0 {
						// This ensure even if the muti-json parser do not support a setting,
						// It is still respected automatically for the first configure file
						*cf = *c
						continue
					}
					cf.Override(c, arg)
				}
				return cf.Build()
			case io.Reader:
				if serial.UseStrictJSON {
					cfg, err := serial.DecodeJSONConfigStrict(v)
					if err != nil {
						return nil, err
					}
					return cfg.Build()
				}
				return serial.LoadJSONConfig(v)
			default:
				return nil, errors.New("unknown type")
			}
		},
	}))
}
