package constants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBaseURL_Default(t *testing.T) {
	t.Setenv("RAZORPAY_BASE_URL", "")

	assert.Equal(t, DEFAULT_BASE_URL, GetBaseURL(), "GetBaseURL should return DEFAULT_BASE_URL when RAZORPAY_BASE_URL is unset or empty")
}

func TestGetBaseURL_EnvOverride(t *testing.T) {
	customURL := "https://api-web.ext.dev.razorpay.in"
	t.Setenv("RAZORPAY_BASE_URL", customURL)

	assert.Equal(t, customURL, GetBaseURL(), "GetBaseURL should return RAZORPAY_BASE_URL when set")
}
