package config

type DNS struct {
    // CustomResolver: when true, use a custom DNS server specified in the config, rather than the machine's default DNS resolver
    CustomResolver bool

    // Address is the address of the custom DNS server (e.g. 1.1.1.1)
    Address string

    // Protocol to use: "udp" (or "" - default), "dns-over-https", "dns-over-tls"
    // Optional: default is "" (= "udp")
    Protocol string

    // DNSoverHTTPSHost is the value for the "Host" header used when making requests via DNS-over-HTTPS
    // Optional: use if necessary
    DNSoverHTTPSHost string

    // Port is the port the DNS server listens to
    // Optional: default value is 53 (udp), 443 (dns-over-https) or 853 (dns-over-tls) depending on the protocol used
    Port uint
}

