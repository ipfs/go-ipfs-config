package config

// P2P defines the p2p configuration for listen and forward
type P2P struct {
	Listen  map[string]string
	Forward map[string]ForwardConfig
}

// ForwardDetail defines the forward configuration
type ForwardConfig struct {
	Peer     string
	Protocol string
}
