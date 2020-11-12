package config

// Pinning configures the pinning services.
type Pinning struct {
	// Services lists the pinning services
	RemoteServices []RemotePinningService
}

type RemotePinningService struct {
	Name        string
	ApiEndpoint string
	ApiKey      string
}
