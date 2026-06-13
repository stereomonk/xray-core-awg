package wireguard

import (
	"context"
	"errors"
	"fmt"
	"net/netip"
	"reflect"
	"strings"

	"github.com/xtls/xray-core/common"
)

func init() {
	common.Must(common.RegisterConfig((*DeviceConfig)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		deviceConfig := config.(*DeviceConfig)
		if deviceConfig.IsClient {
			return New(ctx, deviceConfig)
		} else {
			return NewServer(ctx, deviceConfig)
		}
	}))
}

// convert endpoint string to netip.Addr
func parseEndpoints(conf *DeviceConfig) ([]netip.Addr, bool, bool, error) {
	var hasIPv4, hasIPv6 bool

	endpoints := make([]netip.Addr, len(conf.Endpoint))
	for i, str := range conf.Endpoint {
		var addr netip.Addr
		if strings.Contains(str, "/") {
			prefix, err := netip.ParsePrefix(str)
			if err != nil {
				return nil, false, false, err
			}
			addr = prefix.Addr()
			if prefix.Bits() != addr.BitLen() {
				return nil, false, false, errors.New("interface address subnet should be /32 for IPv4 and /128 for IPv6")
			}
		} else {
			var err error
			addr, err = netip.ParseAddr(str)
			if err != nil {
				return nil, false, false, err
			}
		}
		endpoints[i] = addr

		if addr.Is4() {
			hasIPv4 = true
		} else if addr.Is6() {
			hasIPv6 = true
		}
	}

	return endpoints, hasIPv4, hasIPv6, nil
}

// serialize the config into an IPC request
func createIPCRequest(conf *DeviceConfig) string {
	var request strings.Builder

	request.WriteString(fmt.Sprintf("private_key=%s\n", conf.SecretKey))

	if conf.Amnezia != nil {
		// write parametr helper function
		_wp := func(name string, value any) {
			// check if not 0 or "", for safe use only Get*
			if !reflect.ValueOf(value).IsZero() {
				request.WriteString(fmt.Sprintf("%s=%v\n", name, value))
			}
		}
		awg := conf.Amnezia

		_wp("jc", awg.GetJC())
		_wp("jmin", awg.GetJMin())
		_wp("jmax", awg.GetJMax())

		_wp("s1", awg.GetS1())
		_wp("s2", awg.GetS2())
		_wp("s3", awg.GetS3())
		_wp("s4", awg.GetS4())

		_wp("h1", awg.GetH1())
		_wp("h2", awg.GetH2())
		_wp("h3", awg.GetH3())
		_wp("h4", awg.GetH4())

		_wp("i1", awg.GetI1())
		_wp("i2", awg.GetI2())
		_wp("i3", awg.GetI3())
		_wp("i4", awg.GetI4())
		_wp("i5", awg.GetI5())
	}

	if !conf.IsClient {
		// placeholder, we'll handle actual port listening on Xray
		request.WriteString("listen_port=1337\n")
	}

	for _, peer := range conf.Peers {
		if peer.PublicKey != "" {
			request.WriteString(fmt.Sprintf("public_key=%s\n", peer.PublicKey))
		}

		if peer.PreSharedKey != "" {
			request.WriteString(fmt.Sprintf("preshared_key=%s\n", peer.PreSharedKey))
		}

		if peer.Endpoint != "" {
			request.WriteString(fmt.Sprintf("endpoint=%s\n", peer.Endpoint))
		}

		for _, ip := range peer.AllowedIps {
			request.WriteString(fmt.Sprintf("allowed_ip=%s\n", ip))
		}

		if peer.KeepAlive != 0 {
			request.WriteString(fmt.Sprintf("persistent_keepalive_interval=%d\n", peer.KeepAlive))
		}
	}

	return request.String()[:request.Len()]
}
