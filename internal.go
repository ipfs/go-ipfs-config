package config

type Internal struct {
	Bitswap InternalBitswap
}

const (
	DefaultBitswapEngineBlockstoreWorkerCount = 2000
	DefaultBitswapTaskWorkerCount             = 500
	DefaultBitswapEngineTaskWorkerCount       = 500
	DefaultBitswapMaxOutstandingBytesPerPeer  = 1 << 21
)

type InternalBitswap struct {
	EngineBlockstoreWorkerCount int
	TaskWorkerCount             int
	EngineTaskWorkerCount       int
	MaxOutstandingBytesPerPeer  int
}
