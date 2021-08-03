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

// InternalBitswap contains knobs for tuning bitswap resource utilization.
// The knobs (below) document how their value should related to each other.
// Whether their values should be raised or lowered should be determined
// based on the metrics active_tasks, pending_tasks, pending_block_tasks and active_block_tasks
// reported by bitswap.
//
// The value of active_tasks is capped by EngineTaskWorkerCount.
//
// The value of pending_tasks is generally capped by the knobs below,
// however its exact maximum value is hard to predict as it depends on task sizes
// as well as number of requesting peers. However, as a rule of thumb,
// during healthy operation this value should oscillate around a "typical" low value
// (without hitting a plateau continuously).
//
// If pending_tasks grows and eventually reaches a plateau,
// while at the same time active_tasks is at its maximum,
// the node has reached its resource limits and requests are being dropped (or not serviced).
// Raising resource limits (using the knobs below) could help, assuming the hardware can support the new limits.
//
// The value of active_block_tasks is capped by EngineBlockstoreWorkerCount.
//
// The value of pending_block_tasks is indirectly capped by active_tasks, but cannot be predicted
// as it depends on the number of blocks involved in a peer task which can vary.
//
// If the value of pending_block_tasks is observed to grow,
// while active_block_tasks is at its maximum, there is indication that the number of
// available block tasks is creating a bottleneck (either due to high-latency block operations,
// or due to high number of block operations per bitswap peer task).
// In such cases, try increasing the EngineBlockstoreWorkerCount.
type InternalBitswap struct {
	// number of threads sending outgoing messages.
	// used to thottle the number of concurrent send operations.
	TaskWorkerCount int

	// number of threads for blockstore operations.
	// used to throttle the number of concurrent requests to the block store.
	// this number should generally be a low multiple (e.g. 4) of the TaskWorkerCount,
	// but its optimal value can be informed by the metrics pending_block_tasks and active_block_tasks.
	EngineBlockstoreWorkerCount int

	// number of worker threads for decision engine task worker.
	// used to throttle the number of (send) tasks scheduled in parallel, therefore
	// this number should generally be equal to TaskWorkerCount.
	EngineTaskWorkerCount int

	// maximum number of bytes (across all tasks) pending to be processed and sent to any individual peer.
	// this number controls fairness and can very from 250Kb (very fair) to 10Mb (less fair, with more work
	// dedicated to peers who ask for more). Values below 250Kb could cause thrashing.
	// Values above 10Mb open the potential for aggressively-wanting peers to consume all resources and
	// deteriorate the quality provided to less aggressively-wanting peers.
	MaxOutstandingBytesPerPeer int
}
