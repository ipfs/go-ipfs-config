package config

type DNS struct {
	// Resolver: multi-address for the DNS resolver to use. If this is empty, it will use the system's built-in resolver. Examples:
	// /ip4/1.1.1.1/udp/53 -> Sets up a DNS resolver using the standard (UDP, plain-text) DNS protocol
	// /ip4/1.1.1.1/tcp/853/tls -> DNS-over-TLS
	// /ip4/1.1.1.1/tcp/443/https -> DNS-over-HTTPS
	// /ip4/1.1.1.1/tcp/443/https/cloudflare-dns.com -> DNS-over-HTTPS, with explicit hostname (recommended)
	// /dns4/cloudflare-dns.com/tcp/443/https -> DNS-over-HTTPS using a hostname (this will first resolve the hostname using the system's built-in resolver)
	Resolver string
}
