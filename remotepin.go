package config

const (
	PinningTag             = "Pinning"
	RemoteServicesTag      = "RemoteServices"
	RemoteServicesSelector = PinningTag + "." + RemoteServicesTag
)

type Pinning struct {
	RemoteServices  map[string]RemotePinningService
	DefaultPolicies RemotePinningServicePolicies
}

type RemotePinningService struct {
	Api      RemotePinningServiceApi
	Policies RemotePinningServicePolicies
}

type RemotePinningServiceApi struct {
	Endpoint string
	Key      string
}

type RemotePinningServicePolicies struct {
	MFS RemotePinningServiceMFSPolicy
}

type RemotePinningServiceMFSPolicy struct {
	// Enable enables watching for changes in MFS and re-pinning the MFS root cid whenever a change occurs.
	Enable bool
	// MFSRepinInterval determines the repin interval when the policy is enabled. In ns, us, ms, s, m, h.
	MFSRepinInterval string
}
