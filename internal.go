package config

type Internal struct {
	Bitswap                     *InternalBitswap `json:",omitempty"` // This is omitempty since we are expecting to make changes to all subcomponents of Internal
	UnixFSShardingSizeThreshold *OptionalString  `json:",omitempty"`
}

type InternalBitswap struct {
	TaskWorkerCount             OptionalInteger
	EngineBlockstoreWorkerCount OptionalInteger
	EngineTaskWorkerCount       OptionalInteger
	MaxOutstandingBytesPerPeer  OptionalInteger
}
