package config

const (
	PinningTag             = "Pinning"
	RemoteServicesTag      = "RemoteServices"
	RemoteServicesSelector = PinningTag + "." + RemoteServicesTag
)

type Pinning struct {
	MFSRepinInterval string // in ns, us, ms, s, m, h
	RemoteServices   map[string]RemotePinningService
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
	// PinMFS enables watching for changes in MFS and re-pinning the MFS root cid whenever a change occurs.
	PinMFS *bool
}
