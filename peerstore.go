package config

// PeerstoreType is an enum to specify the peerstore backend.
type PeerstoreType string

const (
	PeerstoreMemory    PeerstoreType = "memory"
	PeerstoreDatastore PeerstoreType = "datastore"
)

type Peerstore struct {
	// The type of the peerstore implementation.
	Type PeerstoreType
	// Configuration of the datastore-backed peerstore.
	// See default values in: https://godoc.org/github.com/libp2p/go-libp2p-peerstore/pstoreds#DefaultOpts.
	Datastore struct {
		// Configuration for the in-memory ARC cache that alleviates disk access for frequently and recently used keys.
		Cache struct {
			// Disables the in-memory cache.
			Disable bool
			// Size of the in-memory cache.
			Size uint32
		}

		GC struct {
			// Disables automatic GC.
			Disable bool
			// Sweep interval to purge expired addresses from the datastore, in millis.
			PurgeIntervalMillis uint32
			// Interval to renew the GC lookahead window, in millis. Disabled by default.
			LookaheadIntervalMillis uint32
			// Initial delay before GC processes start, in millis. Intended to give the system breathing room
			// to fully boot before starting GC.
			InitDelayMillis *uint32 `json:"omitempty"`
		}
	}
}
