package config

type Internal struct {
	Bitswap InternalBitswap
}

type InternalBitswap struct {
	TaskWorkerCount            int
	EngineTaskWorkerCount      int
	MaxOutstandingBytesPerPeer int
}
