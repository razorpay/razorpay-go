package requests

import (
	"net/http"
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

func TestAuthType(t *testing.T) {
	// Test AuthType constants
	if BasicAuth != 0 {
		t.Errorf("Expected BasicAuth to be 0, got %d", BasicAuth)
	}
	if OAuth != 1 {
		t.Errorf("Expected OAuth to be 1, got %d", OAuth)
	}
}

func TestSetAuthentication(t *testing.T) {
	tests := []struct {
		name           string
		request        Request
		expectedAuth   string
		expectedHeader string
	}{
		{
			name: "BasicAuth with key and secret",
			request: Request{
				Auth:     Auth{Key: "test_key", Secret: "test_secret"},
				AuthType: BasicAuth,
			},
			expectedAuth: "Basic", // This will be set by SetBasicAuth
		},
		{
			name: "OAuth with token",
			request: Request{
				Auth:     Auth{Token: "oauth_token_123"},
				AuthType: OAuth,
			},
			expectedHeader: "Bearer oauth_token_123",
		},
		{
			name: "OAuth with empty token",
			request: Request{
				Auth:     Auth{Token: ""},
				AuthType: OAuth,
			},
			expectedHeader: "Bearer ",
		},
		{
			name: "BasicAuth with empty credentials",
			request: Request{
				Auth:     Auth{Key: "", Secret: ""},
				AuthType: BasicAuth,
			},
			expectedAuth: "Basic",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test HTTP request
			httpReq, err := http.NewRequest("GET", "https://api.razorpay.com/v1/test", nil)
			if err != nil {
				t.Fatalf("Failed to create HTTP request: %v", err)
			}

			// Call setAuthentication
			tt.request.setAuthentication(httpReq)

			// Verify the authentication was set correctly
			if tt.request.AuthType == OAuth {
				authHeader := httpReq.Header.Get("Authorization")
				if authHeader != tt.expectedHeader {
					t.Errorf("Expected Authorization header to be %s, got %s", tt.expectedHeader, authHeader)
				}
			} else if tt.request.AuthType == BasicAuth {
				authHeader := httpReq.Header.Get("Authorization")
				// For BasicAuth, just check that some Authorization header was set
				if authHeader == "" {
					t.Error("Expected BasicAuth to set Authorization header")
				}
				// Verify it starts with "Basic "
				if len(authHeader) < 6 || authHeader[:6] != "Basic " {
					t.Errorf("Expected BasicAuth header to start with 'Basic ', got %s", authHeader)
				}
			}
		})
	}
}

func TestSetAuthenticationWithDifferentTokens(t *testing.T) {
	tests := []struct {
		name     string
		token    string
		expected string
	}{
		{
			name:     "standard OAuth token",
			token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
			expected: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
		},
		{
			name:     "token with special characters",
			token:    "token-with-dashes_and_underscores.and.dots",
			expected: "Bearer token-with-dashes_and_underscores.and.dots",
		},
		{
			name:     "very long token",
			token:    "very_long_token_that_might_be_used_in_real_scenarios_with_lots_of_characters_1234567890",
			expected: "Bearer very_long_token_that_might_be_used_in_real_scenarios_with_lots_of_characters_1234567890",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := Request{
				Auth:     Auth{Token: tt.token},
				AuthType: OAuth,
			}

			httpReq, err := http.NewRequest("GET", "https://api.razorpay.com/v1/test", nil)
			if err != nil {
				t.Fatalf("Failed to create HTTP request: %v", err)
			}

			request.setAuthentication(httpReq)

			authHeader := httpReq.Header.Get("Authorization")
			if authHeader != tt.expected {
				t.Errorf("Expected Authorization header to be %s, got %s", tt.expected, authHeader)
			}
		})
	}
}

func TestAuthStruct(t *testing.T) {
	// Test Auth struct fields
	auth := Auth{
		Key:    "test_key",
		Secret: "test_secret",
		Token:  "test_token",
	}

	if auth.Key != "test_key" {
		t.Errorf("Expected Key to be 'test_key', got %s", auth.Key)
	}
	if auth.Secret != "test_secret" {
		t.Errorf("Expected Secret to be 'test_secret', got %s", auth.Secret)
	}
	if auth.Token != "test_token" {
		t.Errorf("Expected Token to be 'test_token', got %s", auth.Token)
	}
}
