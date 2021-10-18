package config

type SwarmConfig struct {
	// AddrFilters specifies a set libp2p addresses that we should never
	// dial or receive connections from.
	AddrFilters []string

	// DisableBandwidthMetrics disables recording of bandwidth metrics for a
	// slight reduction in memory usage. You probably don't need to set this
	// flag.
	DisableBandwidthMetrics bool

	// DisableNatPortMap turns off NAT port mapping (UPnP, etc.).
	DisableNatPortMap bool

	// DisableRelay explicitly disables the relay transport.
	//
	// Deprecated: This flag is deprecated and is overridden by
	// `Transports.Relay` if specified.
	DisableRelay bool `json:",omitempty"`

	// DisableRelayService disables the limited relay (circuit v2 relay).
	DisableRelayService bool

	RelayServiceOpts RelayResources

	// EnableAutoRelay enables the "auto relay" feature.
	//
	// When both EnableAutoRelay and EnableRelayHop are set, this go-ipfs node
	// will advertise itself as a public relay. Otherwise it will find and use
	// advertised public relays when it determines that it's not reachable
	// from the public internet.
	EnableAutoRelay bool

	// Transports contains flags to enable/disable libp2p transports.
	Transports Transports

	// ConnMgr configures the connection manager.
	ConnMgr ConnMgr
}

// RelayResources configures the resources of the relay.
// For any value set to 0, a reasonable default will be used.
type RelayResources struct {
	// Limit is the (optional) relayed connection limits.
	Limit RelayLimit

	// ReservationTTL is the duration of a new (or refreshed reservation).
	ReservationTTL Duration

	// MaxReservations is the maximum number of active relay slots.
	MaxReservations int
	// MaxCircuits is the maximum number of open relay connections for each peer; defaults to 16.
	MaxCircuits int
	// BufferSize is the size of the relayed connection buffers.
	BufferSize int

	// MaxReservationsPerPeer is the maximum number of reservations originating from the same peer.
	MaxReservationsPerPeer int
	// MaxReservationsPerIP is the maximum number of reservations originating from the same IP address.
	MaxReservationsPerIP int
	// MaxReservationsPerASN is the maximum number of reservations origination from the same ASN.
	MaxReservationsPerASN int
}

// RelayLimit are the per relayed connection resource limits.
type RelayLimit struct {
	// Duration is the time limit before resetting a relayed connection.
	Duration Duration
	// Data is the limit of data relayed (on each direction) before resetting the connection.
	Data int64
}

type Transports struct {
	// Network specifies the base transports we'll use for dialing. To
	// listen on a transport, add the transport to your Addresses.Swarm.
	Network struct {
		// All default to on.
		QUIC      Flag `json:",omitempty"`
		TCP       Flag `json:",omitempty"`
		Websocket Flag `json:",omitempty"`
		Relay     Flag `json:",omitempty"`
	}

	// Security specifies the transports used to encrypt insecure network
	// transports.
	Security struct {
		// Defaults to 100.
		TLS Priority `json:",omitempty"`
		// Defaults to 200.
		SECIO Priority `json:",omitempty"`
		// Defaults to 300.
		Noise Priority `json:",omitempty"`
	}

	// Multiplexers specifies the transports used to multiplex multiple
	// connections over a single duplex connection.
	Multiplexers struct {
		// Defaults to 100.
		Yamux Priority `json:",omitempty"`
		// Defaults to 200.
		Mplex Priority `json:",omitempty"`
	}
}

// ConnMgr defines configuration options for the libp2p connection manager
type ConnMgr struct {
	Type        string
	LowWater    int
	HighWater   int
	GracePeriod string
}
