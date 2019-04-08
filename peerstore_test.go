package config

import (
	"encoding/json"
	"testing"
)

func TestPeerstoreDecode(t *testing.T) {
	str := `
		{
			"Type": "datastore",
			"Datastore": {
				"Cache": {
					"Disable": true,
					"Size": 999
				},
				"GC": {
					"Disable": true,
					"PurgeIntervalMillis": 888,
					"LookaheadIntervalMillis": 777
				}
			}
		}
	`

	var cfg Peerstore
	if err := json.Unmarshal([]byte(str), &cfg); err != nil {
		t.Errorf("failed while unmarshalling peerstore struct: %s", err)
	}

	// check some sample values
	if cfg.Type != PeerstoreDatastore {
		t.Errorf("unexpected peerstore type: %s", cfg.Type)
	}

	if !cfg.Datastore.Cache.Disable {
		t.Error("cached should be disabled")
	}

	if cfg.Datastore.GC.PurgeIntervalMillis != 888 {
		t.Errorf("unexpected purge interval: %d", cfg.Datastore.GC.PurgeIntervalMillis)
	}

	if cfg.Datastore.GC.InitDelayMillis != nil {
		t.Error("expected nil value for InitDelayMillis")
	}
}
