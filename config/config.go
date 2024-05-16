package config

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	// URL where the SDK config is hosted.
	// format:
	// ```json
	// {
	//   "dns": {
	//     "SG": "<host1>",
	//     "US": "<host2>",
	//   }
	// }
	//```
	//  TODO: update the URL
	URL     = ""
	Timeout = time.Second
)

// Config holds the configuration needed by the SDK to operate optimally
type Config struct {
	DNS DNSConfig `json:"dns"`
}

// new returns Config struct with default values
func new() Config {
	return Config{
		DNS: newDNSConfig(),
	}
}

// New returns Config struct loaded with the remote configuration available in URL specified
func New() Config {
	cfg := new()
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		log.Println("failed to create config fetch request:", err)
		return cfg
	}

	client := &http.Client{Timeout: Timeout}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("failed to fetch remote config:", err)
		return cfg
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("failed to read the config response:", err)
		return cfg
	}

	if err = json.Unmarshal(body, &cfg); err != nil {
		log.Println("failed to unmarshal the config:", err)
	}
	return cfg
}

// GetDNS given the key it returns with the DNS address
func (c Config) GetDNS(key string) string {
	return c.DNS.get(key)
}
