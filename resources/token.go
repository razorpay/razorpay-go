package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Token ...
type Token struct {
	Request *requests.Request
}

// Create creates a new token for the given data.
func (t *Token) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s/tokens", constants.VERSION_V1)
	return t.Request.Post(url, data, extraHeaders)
}

// Fetch fetches a token having the given tokenID associated with a customer having the given customerID.
func (t *Token) Fetch(customerID string, tokenID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/tokens/%s", constants.VERSION_V1, constants.CUSTOMER_URL, customerID, tokenID)
	return t.Request.Get(url, queryParams, extraHeaders)
}

// All fetches collection of tokens for associated with a customer having the given customerID.
func (t *Token) All(customerID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/tokens", constants.VERSION_V1, constants.CUSTOMER_URL, customerID)
	return t.Request.Get(url, queryParams, extraHeaders)
}

// Delete deletes a token having the given tokenID associated with a customer having the given customerID.
func (t *Token) Delete(customerID string, tokenID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/tokens/%s", constants.VERSION_V1, constants.CUSTOMER_URL, customerID, tokenID)
	return t.Request.Delete(url, queryParams, extraHeaders)
}

// FetchCardPropertiesByToken fetches card properties of an existing token.
func (t *Token) FetchCardPropertiesByToken(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s/tokens/fetch", constants.VERSION_V1)
	return t.Request.Post(url, data, extraHeaders)
}

// DeleteToken deletes a token.
func (t *Token) DeleteToken(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s/tokens/delete", constants.VERSION_V1)
	return t.Request.Post(url, data, extraHeaders)
}

// ProcessPaymentOnAlternatePAorPG process a payment on the tokenised card on another PA/PG with token
// created on razorpay
func (t *Token) ProcessPaymentOnAlternatePAorPG(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s/tokens/service_provider_tokens/token_transactional_data", constants.VERSION_V1)
	return t.Request.Post(url, data, extraHeaders)
}
