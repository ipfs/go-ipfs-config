package config

var (
	RemoteServicesPath     = "Pinning.RemoteServices"
	PinningConcealSelector = []string{"Pinning", "RemoteServices", "*", "API", "Key"}
)

type Pinning struct {
	RemoteServices map[string]RemotePinningService `json:",omitempty"`
}

type RemotePinningService struct {
	API      RemotePinningServiceAPI
	Policies RemotePinningServicePolicies `json:",omitempty"`
}

type RemotePinningServiceAPI struct {
	Endpoint string
	Key      string
}

type RemotePinningServicePolicies struct {
	MFS RemotePinningServiceMFSPolicy `json:",omitempty"`
}

type RemotePinningServiceMFSPolicy struct {
	// Enable enables watching for changes in MFS and re-pinning the MFS root cid whenever a change occurs.
	Enable bool
	// Name is the pin name for MFS.
	PinName string `json:",omitempty"`
	// RepinInterval determines the repin interval when the policy is enabled. In ns, us, ms, s, m, h.
	RepinInterval string `json:",omitempty"`
}
