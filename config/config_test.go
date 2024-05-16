package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_GetDNS(t *testing.T) {
	cfg := new()

	assert.Equal(t, DefaultDNS, cfg.GetDNS("rzp_live_1234567890123"))
	assert.Equal(t, DefaultDNS, cfg.GetDNS("rzp_live_12"))
	assert.Equal(t, DefaultDNS, cfg.GetDNS("rzp_live_sg123123123123123"))

	customINDNS := "api-in.razorpay.com"
	customSGDNS := "api-sg.razorpay.com"

	// TODO: replace with New() once we have the URL up and running
	cfg = Config{DNS: DNSConfig{
		"IN": customINDNS,
		"SG": customSGDNS,
	}}

	assert.Equal(t, customINDNS, cfg.GetDNS("rzp_live_1234567890123"))
	assert.Equal(t, customINDNS, cfg.GetDNS("rzp_live_12"))
	assert.Equal(t, customSGDNS, cfg.GetDNS("rzp_live_sg123123123123123"))
}
