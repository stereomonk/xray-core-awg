package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/stereomonk/xray-core-awg/app/dispatcher"
	_ "github.com/stereomonk/xray-core-awg/app/proxyman/inbound"
	_ "github.com/stereomonk/xray-core-awg/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/stereomonk/xray-core-awg/app/commander"
	_ "github.com/stereomonk/xray-core-awg/app/log/command"
	_ "github.com/stereomonk/xray-core-awg/app/proxyman/command"
	_ "github.com/stereomonk/xray-core-awg/app/stats/command"

	// Developer preview services
	_ "github.com/stereomonk/xray-core-awg/app/observatory/command"

	// Other optional features.
	_ "github.com/stereomonk/xray-core-awg/app/dns"
	_ "github.com/stereomonk/xray-core-awg/app/dns/fakedns"
	_ "github.com/stereomonk/xray-core-awg/app/log"
	_ "github.com/stereomonk/xray-core-awg/app/metrics"
	_ "github.com/stereomonk/xray-core-awg/app/policy"
	_ "github.com/stereomonk/xray-core-awg/app/reverse"
	_ "github.com/stereomonk/xray-core-awg/app/router"
	_ "github.com/stereomonk/xray-core-awg/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/stereomonk/xray-core-awg/transport/internet/tagged/taggedimpl"

	// Developer preview features
	_ "github.com/stereomonk/xray-core-awg/app/observatory"

	// Inbound and outbound proxies.
	_ "github.com/stereomonk/xray-core-awg/proxy/blackhole"
	_ "github.com/stereomonk/xray-core-awg/proxy/dns"
	_ "github.com/stereomonk/xray-core-awg/proxy/dokodemo"
	_ "github.com/stereomonk/xray-core-awg/proxy/freedom"
	_ "github.com/stereomonk/xray-core-awg/proxy/http"
	_ "github.com/stereomonk/xray-core-awg/proxy/loopback"
	_ "github.com/stereomonk/xray-core-awg/proxy/shadowsocks"
	_ "github.com/stereomonk/xray-core-awg/proxy/socks"
	_ "github.com/stereomonk/xray-core-awg/proxy/trojan"
	_ "github.com/stereomonk/xray-core-awg/proxy/vless/inbound"
	_ "github.com/stereomonk/xray-core-awg/proxy/vless/outbound"
	_ "github.com/stereomonk/xray-core-awg/proxy/vmess/inbound"
	_ "github.com/stereomonk/xray-core-awg/proxy/vmess/outbound"
	_ "github.com/stereomonk/xray-core-awg/proxy/wireguard"

	// Transports
	_ "github.com/stereomonk/xray-core-awg/transport/internet/grpc"
	_ "github.com/stereomonk/xray-core-awg/transport/internet/httpupgrade"
	_ "github.com/stereomonk/xray-core-awg/transport/internet/kcp"
	_ "github.com/stereomonk/xray-core-awg/transport/internet/reality"
	_ "github.com/stereomonk/xray-core-awg/transport/internet/splithttp"
	_ "github.com/stereomonk/xray-core-awg/transport/internet/tcp"
	_ "github.com/stereomonk/xray-core-awg/transport/internet/tls"
	_ "github.com/stereomonk/xray-core-awg/transport/internet/udp"
	_ "github.com/stereomonk/xray-core-awg/transport/internet/websocket"

	// Transport headers
	_ "github.com/stereomonk/xray-core-awg/transport/internet/headers/http"
	_ "github.com/stereomonk/xray-core-awg/transport/internet/headers/noop"

	// JSON & TOML & YAML
	_ "github.com/stereomonk/xray-core-awg/main/json"
	_ "github.com/stereomonk/xray-core-awg/main/toml"
	_ "github.com/stereomonk/xray-core-awg/main/yaml"

	// Load config from file or http(s)
	_ "github.com/stereomonk/xray-core-awg/main/confloader/external"

	// Commands
	_ "github.com/stereomonk/xray-core-awg/main/commands/all"
)
