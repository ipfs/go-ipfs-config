package config

// Pinning configures the pinning services.
type Pinning struct {
	// Services lists the pinning services
	Services []PinningServices
}

type PinningServices struct {
	Name        string
	ApiEndpoint string
	ApiKey      string
}
