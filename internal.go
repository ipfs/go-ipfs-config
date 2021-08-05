package config

type Internal struct {
	Bitswap InternalBitswap
}

type InternalBitswap struct {
	TaskWorkerCount             int
	EngineBlockstoreWorkerCount int
	EngineTaskWorkerCount       int
	MaxOutstandingBytesPerPeer  int
}
