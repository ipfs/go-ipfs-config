package config

type PubsubConfig struct {
	// Router can be either floodsub (legacy) or gossipsub (new and
	// backwards compatible).
	Router string

	// DisableSigning disables message signing. Message signing is *enabled*
	// by default.
	DisableSigning bool

	// StrictSignatureVerification is a noop. When DisableSigning is false,
	// messages are _always_ verified.
	StrictSignatureVerification bool
}
