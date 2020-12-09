package config

const (
	PinningTag             = "Pinning"
	RemoteServicesTag      = "RemoteServices"
	RemoteServicesSelector = PinningTag + "." + RemoteServicesTag
)

type Pinning struct {
	RemoteServices  map[string]RemotePinningService `json:",omitempty"`
	DefaultPolicies RemotePinningServicePolicies    `json:",omitempty"`
}

type RemotePinningService struct {
	Api      RemotePinningServiceApi
	Policies RemotePinningServicePolicies `json:",omitempty"`
}

type RemotePinningServiceApi struct {
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
	Name string
	// RepinInterval determines the repin interval when the policy is enabled. In ns, us, ms, s, m, h.
	RepinInterval string `json:",omitempty"`
}
