package razorpay

import (
	"testing"
	"time"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	key := "rzp_test_key123"
	secret := "test_secret_456"
	t.Setenv("RAZORPAY_BASE_URL", "")

	client := NewClient(key, secret)

	assert.NotNil(t, client, "Client should not be nil")
	assert.Equal(t, key, client.Request.Auth.Key, "Key should be set correctly")
	assert.Equal(t, secret, client.Request.Auth.Secret, "Secret should be set correctly")
	assert.Equal(t, requests.BasicAuth, client.Request.AuthType, "AuthType should be BasicAuth")
	assert.Equal(t, constants.DEFAULT_BASE_URL, client.Request.BaseURL, "BaseURL should default to the production api url")
}

func TestNewClient_BaseURLOverride(t *testing.T) {
	customURL := "https://api-web.ext.dev.razorpay.in"
	t.Setenv("RAZORPAY_BASE_URL", customURL)

	client := NewClient("rzp_test_key", "test_secret")

	assert.Equal(t, customURL, client.Request.BaseURL, "BaseURL should be read from RAZORPAY_BASE_URL")
}

func TestNewClient_ResourcesInitialized(t *testing.T) {
	client := NewClient("rzp_test_key", "test_secret")

	// Verify all 25 resources are initialized
	assert.NotNil(t, client.Addon, "Addon resource should be initialized")
	assert.NotNil(t, client.Account, "Account resource should be initialized")
	assert.NotNil(t, client.Card, "Card resource should be initialized")
	assert.NotNil(t, client.Customer, "Customer resource should be initialized")
	assert.NotNil(t, client.Invoice, "Invoice resource should be initialized")
	assert.NotNil(t, client.PaymentLink, "PaymentLink resource should be initialized")
	assert.NotNil(t, client.Order, "Order resource should be initialized")
	assert.NotNil(t, client.Payment, "Payment resource should be initialized")
	assert.NotNil(t, client.Plan, "Plan resource should be initialized")
	assert.NotNil(t, client.Product, "Product resource should be initialized")
	assert.NotNil(t, client.Refund, "Refund resource should be initialized")
	assert.NotNil(t, client.Subscription, "Subscription resource should be initialized")
	assert.NotNil(t, client.Token, "Token resource should be initialized")
	assert.NotNil(t, client.Transfer, "Transfer resource should be initialized")
	assert.NotNil(t, client.VirtualAccount, "VirtualAccount resource should be initialized")
	assert.NotNil(t, client.QrCode, "QrCode resource should be initialized")
	assert.NotNil(t, client.FundAccount, "FundAccount resource should be initialized")
	assert.NotNil(t, client.Settlement, "Settlement resource should be initialized")
	assert.NotNil(t, client.Stakeholder, "Stakeholder resource should be initialized")
	assert.NotNil(t, client.Item, "Item resource should be initialized")
	assert.NotNil(t, client.Iin, "Iin resource should be initialized")
	assert.NotNil(t, client.Webhook, "Webhook resource should be initialized")
	assert.NotNil(t, client.Document, "Document resource should be initialized")
	assert.NotNil(t, client.Dispute, "Dispute resource should be initialized")
	assert.NotNil(t, client.Payout, "Payout resource should be initialized")
}

func TestNewClientOAuth(t *testing.T) {
	token := "oauth_token_abc123"
	t.Setenv("RAZORPAY_BASE_URL", "")

	client := NewClientOAuth(token)

	assert.NotNil(t, client, "OAuth client should not be nil")
	assert.Equal(t, token, client.Request.Auth.Token, "Token should be set correctly")
	assert.Equal(t, requests.OAuth, client.Request.AuthType, "AuthType should be OAuth")
	assert.Equal(t, constants.DEFAULT_BASE_URL, client.Request.BaseURL, "BaseURL should default to the production api url")
}

func TestNewClientOAuth_BaseURLOverride(t *testing.T) {
	customURL := "https://api-web.ext.dev.razorpay.in"
	t.Setenv("RAZORPAY_BASE_URL", customURL)

	client := NewClientOAuth("oauth_token_test")

	assert.Equal(t, customURL, client.Request.BaseURL, "BaseURL should be read from RAZORPAY_BASE_URL")
}

func TestNewClientOAuth_ResourcesInitialized(t *testing.T) {
	client := NewClientOAuth("oauth_token_test")

	// Verify all resources are initialized for OAuth client too
	assert.NotNil(t, client.Order, "Order resource should be initialized")
	assert.NotNil(t, client.Payment, "Payment resource should be initialized")
	assert.NotNil(t, client.Refund, "Refund resource should be initialized")
}

func TestAddHeaders(t *testing.T) {
	client := NewClient("rzp_test_key", "test_secret")

	headers := map[string]string{
		"X-Custom-Header": "custom-value",
		"X-Another":       "another-value",
	}
	client.AddHeaders(headers)

	assert.Equal(t, "custom-value", client.Request.Headers["X-Custom-Header"], "Custom header should be set")
	assert.Equal(t, "another-value", client.Request.Headers["X-Another"], "Another header should be set")
}

func TestAddHeaders_Multiple(t *testing.T) {
	client := NewClient("rzp_test_key", "test_secret")

	// Add first set of headers
	client.AddHeaders(map[string]string{"X-First": "first-value"})
	// Add second set of headers
	client.AddHeaders(map[string]string{"X-Second": "second-value"})

	assert.Equal(t, "first-value", client.Request.Headers["X-First"], "First header should persist")
	assert.Equal(t, "second-value", client.Request.Headers["X-Second"], "Second header should be added")
}

func TestSetTimeout(t *testing.T) {
	client := NewClient("rzp_test_key", "test_secret")

	client.SetTimeout(30)

	assert.NotNil(t, client.Request.HTTPClient, "HTTPClient should be set after SetTimeout")
	assert.Equal(t, 30*time.Second, client.Request.HTTPClient.Timeout, "Timeout should be set to 30 seconds")
}

func TestSetUserAgent(t *testing.T) {
	client := NewClient("rzp_test_key", "test_secret")

	userAgent := "MyApp/1.0"
	client.SetUserAgent(userAgent)

	assert.Equal(t, userAgent, client.Request.GetUserAgent(), "User agent should be set correctly")
}

func TestSetUserAgent_Empty(t *testing.T) {
	client := NewClient("rzp_test_key", "test_secret")

	client.SetUserAgent("")

	assert.Equal(t, "", client.Request.GetUserAgent(), "Empty user agent should be allowed")
}
