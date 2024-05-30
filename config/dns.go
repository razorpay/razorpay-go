package config

import "strings"

const (
	DefaultDNS         = "https://api.razorpay.com"
	DefaultCountryCode = "IN"
)

// DNSConfig holds the DNS related configuration used by the client
type DNSConfig map[string]string

// newDNSConfig creates new DNS configuration struct with defaults
func newDNSConfig() DNSConfig {
	return DNSConfig{}
}

// get returns the dns which is optimal for the given key
// if no match is found, then it returns the default DNS
func (dns DNSConfig) get(key string) string {
	countryCode := deriveCountryCode(key)
	if val, ok := dns[countryCode]; ok {
		return val
	}

	return DefaultDNS
}

// deriveCountryCode derives country code embedded in the key
// if no match is found, then it returns default country code
func deriveCountryCode(key string) string {
	tokens := strings.Split(key, "_")

	if len(tokens) < 4 {
		return DefaultCountryCode
	}

	countryCode := tokens[len(tokens)-2]
	// check for the default format and then return the default country code
	if len(countryCode) != 2 {
		return DefaultCountryCode
	}

	return strings.ToUpper(countryCode)
}
