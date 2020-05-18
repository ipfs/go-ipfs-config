package config

type Experiments struct {
	FilestoreEnabled     bool
	UrlstoreEnabled      bool
	ShardingEnabled      bool
	GraphsyncEnabled     bool
	Libp2pStreamMounting bool
	P2pHttpProxy         bool
	QUIC                 bool
	StrategicProviding   bool
	PreferNoise          bool // PreferNoise security transport over secio
}
