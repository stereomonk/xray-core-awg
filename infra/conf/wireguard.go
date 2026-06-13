package conf

import (
	"encoding/base64"
	"encoding/hex"
	"strings"

	"github.com/stereomonk/xray-core-awg/common/errors"
	"github.com/stereomonk/xray-core-awg/proxy/wireguard"
	"google.golang.org/protobuf/proto"
)

type WireGuardPeerConfig struct {
	PublicKey    string   `json:"publicKey"`
	PreSharedKey string   `json:"preSharedKey"`
	Endpoint     string   `json:"endpoint"`
	KeepAlive    uint32   `json:"keepAlive"`
	AllowedIPs   []string `json:"allowedIPs,omitempty"`
}

func (c *WireGuardPeerConfig) Build() (proto.Message, error) {
	var err error
	config := new(wireguard.PeerConfig)

	if c.PublicKey != "" {
		config.PublicKey, err = ParseWireGuardKey(c.PublicKey)
		if err != nil {
			return nil, err
		}
	}

	if c.PreSharedKey != "" {
		config.PreSharedKey, err = ParseWireGuardKey(c.PreSharedKey)
		if err != nil {
			return nil, err
		}
	}

	config.Endpoint = c.Endpoint
	// default 0
	config.KeepAlive = c.KeepAlive
	if c.AllowedIPs == nil {
		config.AllowedIps = []string{"0.0.0.0/0", "::0/0"}
	} else {
		config.AllowedIps = c.AllowedIPs
	}

	return config, nil
}

type AmneziaParameters struct {
	JunkCount int32 `json:"jc"`
	JunkMin   int32 `json:"jmin"`
	JunkMax   int32 `json:"jmax"`

	InitPadding      int32 `json:"s1"`
	ResponsePadding  int32 `json:"s2"`
	CookiePadding    int32 `json:"s3"`
	TransportPadding int32 `json:"s4"`

	InitHeader      string `json:"h1"`
	ResponseHeader  string `json:"h2"`
	CookieHeader    string `json:"h3"`
	TransportHeader string `json:"h4"`

	Signature1 string `json:"i1"`
	Signature2 string `json:"i2"`
	Signature3 string `json:"i3"`
	Signature4 string `json:"i4"`
	Signature5 string `json:"i5"`
}

type WireGuardConfig struct {
	IsClient bool `json:""`

	NoKernelTun    bool                   `json:"noKernelTun"`
	SecretKey      string                 `json:"secretKey"`
	Address        []string               `json:"address"`
	Peers          []*WireGuardPeerConfig `json:"peers"`
	MTU            int32                  `json:"mtu"`
	NumWorkers     int32                  `json:"workers"`
	Reserved       []byte                 `json:"reserved"`
	DomainStrategy string                 `json:"domainStrategy"`
	Amnezia        *AmneziaParameters     `json:"awg,omitempty"`
}

func (c *WireGuardConfig) Build() (proto.Message, error) {
	config := new(wireguard.DeviceConfig)

	var err error
	config.SecretKey, err = ParseWireGuardKey(c.SecretKey)
	if err != nil {
		return nil, errors.New("invalid WireGuard secret key: %w", err)
	}

	if c.Address == nil {
		// bogon ips
		config.Endpoint = []string{"10.0.0.1", "fd59:7153:2388:b5fd:0000:0000:0000:0001"}
	} else {
		config.Endpoint = c.Address
	}

	if c.Peers != nil {
		config.Peers = make([]*wireguard.PeerConfig, len(c.Peers))
		for i, p := range c.Peers {
			msg, err := p.Build()
			if err != nil {
				return nil, err
			}
			config.Peers[i] = msg.(*wireguard.PeerConfig)
		}
	}

	if c.MTU == 0 {
		config.Mtu = 1420
	} else {
		config.Mtu = c.MTU
	}
	// these a fallback code exists in wireguard-go code,
	// we don't need to process fallback manually
	config.NumWorkers = c.NumWorkers

	if len(c.Reserved) != 0 && len(c.Reserved) != 3 {
		return nil, errors.New(`"reserved" should be empty or 3 bytes`)
	}
	config.Reserved = c.Reserved

	switch strings.ToLower(c.DomainStrategy) {
	case "forceip", "":
		config.DomainStrategy = wireguard.DeviceConfig_FORCE_IP
	case "forceipv4":
		config.DomainStrategy = wireguard.DeviceConfig_FORCE_IP4
	case "forceipv6":
		config.DomainStrategy = wireguard.DeviceConfig_FORCE_IP6
	case "forceipv4v6":
		config.DomainStrategy = wireguard.DeviceConfig_FORCE_IP46
	case "forceipv6v4":
		config.DomainStrategy = wireguard.DeviceConfig_FORCE_IP64
	default:
		return nil, errors.New("unsupported domain strategy: ", c.DomainStrategy)
	}

	config.IsClient = c.IsClient
	config.NoKernelTun = c.NoKernelTun

	config.Amnezia = &wireguard.AmneziaParamters{
		JC:   c.Amnezia.JunkCount,
		JMin: c.Amnezia.JunkMin,
		JMax: c.Amnezia.JunkMax,

		S1: c.Amnezia.InitPadding,
		S2: c.Amnezia.ResponsePadding,
		S3: c.Amnezia.CookiePadding,
		S4: c.Amnezia.TransportPadding,

		H1: c.Amnezia.InitHeader,
		H2: c.Amnezia.ResponseHeader,
		H3: c.Amnezia.CookieHeader,
		H4: c.Amnezia.TransportHeader,

		I1: c.Amnezia.Signature1,
		I2: c.Amnezia.Signature2,
		I3: c.Amnezia.Signature3,
		I4: c.Amnezia.Signature4,
		I5: c.Amnezia.Signature5,
	}

	return config, nil
}

func ParseWireGuardKey(str string) (string, error) {
	var err error

	if str == "" {
		return "", errors.New("key must not be empty")
	}

	if len(str) == 64 {
		_, err = hex.DecodeString(str)
		if err == nil {
			return str, nil
		}
	}

	var dat []byte
	str = strings.TrimSuffix(str, "=")
	if strings.ContainsRune(str, '+') || strings.ContainsRune(str, '/') {
		dat, err = base64.RawStdEncoding.DecodeString(str)
	} else {
		dat, err = base64.RawURLEncoding.DecodeString(str)
	}
	if err == nil {
		return hex.EncodeToString(dat), nil
	}

	return "", errors.New("failed to deserialize key").Base(err)
}
