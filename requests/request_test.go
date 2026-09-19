package requests

import (
	"bytes"
	"io"
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

func TestAddHeaders(t *testing.T) {
	request := &Request{
		Headers: make(map[string]string),
	}

	headers := map[string]string{
		"X-Custom-Header": "custom-value",
		"X-Another":       "another-value",
	}
	request.AddHeaders(headers)

	if request.Headers["X-Custom-Header"] != "custom-value" {
		t.Errorf("Expected X-Custom-Header to be 'custom-value', got %s", request.Headers["X-Custom-Header"])
	}
	if request.Headers["X-Another"] != "another-value" {
		t.Errorf("Expected X-Another to be 'another-value', got %s", request.Headers["X-Another"])
	}
}

func TestAddHeaders_Overwrite(t *testing.T) {
	request := &Request{
		Headers: map[string]string{
			"X-Existing": "old-value",
		},
	}

	request.AddHeaders(map[string]string{
		"X-Existing": "new-value",
	})

	if request.Headers["X-Existing"] != "new-value" {
		t.Errorf("Expected X-Existing to be overwritten to 'new-value', got %s", request.Headers["X-Existing"])
	}
}

func TestSetTimeout(t *testing.T) {
	request := &Request{}

	request.SetTimeout(30)

	if request.HTTPClient == nil {
		t.Error("Expected HTTPClient to be set after SetTimeout")
	}
}

func TestSetUserAgent(t *testing.T) {
	request := &Request{}

	request.SetUserAgent("TestApp/1.0")

	if request.GetUserAgent() != "TestApp/1.0" {
		t.Errorf("Expected user agent to be 'TestApp/1.0', got %s", request.GetUserAgent())
	}
}

func TestGetUserAgent_Empty(t *testing.T) {
	request := &Request{}

	if request.GetUserAgent() != "" {
		t.Errorf("Expected empty user agent, got %s", request.GetUserAgent())
	}
}

// Regression test for #116: the non-2xx path used to return without closing
// response.Body, leaking the connection back to the transport pool.

// closeTrackingBody records whether Close was called.
type closeTrackingBody struct {
	reader io.Reader
	closed *bool
}

func (b *closeTrackingBody) Read(p []byte) (int, error) { return b.reader.Read(p) }

func (b *closeTrackingBody) Close() error {
	*b.closed = true
	return nil
}

// stubTransport returns a canned response with an instrumented body.
type stubTransport struct {
	statusCode int
	body       string
	closed     *bool
}

func (t *stubTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: t.statusCode,
		Status:     http.StatusText(t.statusCode),
		Body:       &closeTrackingBody{reader: bytes.NewBufferString(t.body), closed: t.closed},
		Header:     make(http.Header),
		Request:    req,
	}, nil
}

func TestResponseBodyClosedOnErrorResponse(t *testing.T) {
	cases := []struct {
		name       string
		statusCode int
		body       string
	}{
		{"well-formed razorpay error", 400, `{"error":{"internal_error_code":"BAD_REQUEST_ERROR","description":"amount required"}}`},
		{"malformed json", 400, `{invalid json`},
		{"empty body", 502, ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			closed := false
			request := &Request{
				Auth:       Auth{Key: "key", Secret: "secret"},
				HTTPClient: &http.Client{Transport: &stubTransport{statusCode: tc.statusCode, body: tc.body, closed: &closed}},
				Headers:    map[string]string{},
				BaseURL:    "https://api.razorpay.com",
			}

			_, err := request.Get("/v1/orders/order_123", nil, nil)

			if err == nil {
				t.Fatalf("expected an error for status %d", tc.statusCode)
			}
			if err.Error() == "" {
				t.Errorf("Error() must not be empty for status %d", tc.statusCode)
			}
			if !closed {
				t.Errorf("response.Body was not closed for status %d", tc.statusCode)
			}
		})
	}
}
