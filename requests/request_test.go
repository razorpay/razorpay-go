package requests

import (
	"testing"
)

func TestBuildURLWithParams(t *testing.T) {
	tests := []struct {
		name       string
		requestURL string
		data       map[string]interface{}
		expected   string
	}{
		{
			name:       "basic URL with simple parameters",
			requestURL: "https://api.razorpay.com/v1/payments",
			data: map[string]interface{}{
				"count": 10,
				"skip":  20,
			},
			expected: "https://api.razorpay.com/v1/payments?count=10&skip=20",
		},
		{
			name:       "URL with existing query parameters",
			requestURL: "https://api.razorpay.com/v1/payments?expand=1",
			data: map[string]interface{}{
				"count": 10,
			},
			expected: "https://api.razorpay.com/v1/payments?count=10",
		},
		{
			name:       "parameters with special characters",
			requestURL: "https://api.razorpay.com/v1/payments",
			data: map[string]interface{}{
				"merchant": "ABC & Co.",
			},
			expected: "https://api.razorpay.com/v1/payments?merchant=ABC+%26+Co.",
		},
		{
			name:       "parameters with string array values",
			requestURL: "https://api.razorpay.com/v1/payments",
			data: map[string]interface{}{
				"ids[]": []string{"pay_123", "pay_456", "pay_789"},
			},
			expected: "https://api.razorpay.com/v1/payments?ids%5B%5D=pay_123&ids%5B%5D=pay_456&ids%5B%5D=pay_789",
		},
		{
			name:       "empty parameters",
			requestURL: "https://api.razorpay.com/v1/payments",
			data:       map[string]interface{}{},
			expected:   "https://api.razorpay.com/v1/payments",
		},
		{
			name:       "nil parameters handled as empty",
			requestURL: "https://api.razorpay.com/v1/payments",
			data:       nil,
			expected:   "https://api.razorpay.com/v1/payments",
		},
		{
			name:       "mixed parameter types",
			requestURL: "https://api.razorpay.com/v1/payments",
			data: map[string]interface{}{
				"count":    10,
				"active":   true,
				"keywords": []string{"success", "authorized"},
				"amount":   1999.99,
			},
			expected: "https://api.razorpay.com/v1/payments?active=true&amount=1999.99&count=10&keywords=success&keywords=authorized",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildURLWithParams(tt.requestURL, tt.data)
			if result != tt.expected {
				t.Errorf("buildURLWithParams() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestGetUserAgentHeaderValue(t *testing.T) {
	tests := []struct {
		name     string
		request  Request
		expected string
	}{
		{
			name: "empty user agent",
			request: Request{
				SDKName:   "razorpay-go",
				Version:   "1.3.4",
				userAgent: "",
			},
			expected: "razorpay-go/1.3.4",
		},
		{
			name: "custom user agent",
			request: Request{
				SDKName:   "razorpay-go",
				Version:   "1.3.4",
				userAgent: "CustomApp/2.0",
			},
			expected: "CustomApp/2.0 (razorpay-go/1.3.4)",
		},
		{
			name: "user agent with spaces",
			request: Request{
				SDKName:   "razorpay-go",
				Version:   "1.3.4",
				userAgent: "Custom Application 2.0",
			},
			expected: "Custom Application 2.0 (razorpay-go/1.3.4)",
		},
		{
			name: "different SDK version",
			request: Request{
				SDKName:   "razorpay-go",
				Version:   "2.0.0",
				userAgent: "CustomApp/2.0",
			},
			expected: "CustomApp/2.0 (razorpay-go/2.0.0)",
		},
		{
			name: "different SDK name",
			request: Request{
				SDKName:   "razorpay-custom",
				Version:   "1.3.4",
				userAgent: "CustomApp/2.0",
			},
			expected: "CustomApp/2.0 (razorpay-custom/1.3.4)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.request.getUserAgentHeaderValue()
			if result != tt.expected {
				t.Errorf("getUserAgentHeaderValue() = %v, want %v", result, tt.expected)
			}
		})
	}
}
