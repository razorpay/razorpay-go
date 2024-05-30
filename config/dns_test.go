package config

import (
	"testing"
)

func TestDNS(t *testing.T) {
	tests := []struct {
		name   string
		key    string
		dns    DNSConfig
		expect string
	}{
		{
			name:   "invalid key",
			key:    "test_key",
			dns:    newDNSConfig(),
			expect: DefaultDNS,
		},
		{
			name:   "valid key with default format",
			key:    "rzp_test_qwertyuioplkjh",
			dns:    newDNSConfig(),
			expect: DefaultDNS,
		},
		{
			name:   "valid key with default format with url override",
			key:    "rzp_test_qwertyuioplkjh",
			dns:    DNSConfig{"IN": "https://api-in.razorpay.com"},
			expect: "https://api-in.razorpay.com",
		},
		{
			key: "rzp_test_sg_qwertyuioplkjh",
			dns: DNSConfig{
				"IN": "https://api.razorpay.com",
				"SG": "https://api-sg.razorpay.com",
			},
			expect: "https://api-sg.razorpay.com",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if val := test.dns.get(test.key); val != test.expect {
				t.Errorf("Expected %s, got %s", test.expect, val)
			}
		})
	}
}
