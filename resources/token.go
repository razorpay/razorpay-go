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

// Fetch fetches a token having the given tokenID associated with a customer having the given customerID.
func (t *Token) Fetch(customerID string, tokenID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/tokens/%s", constants.CUSTOMER_URL, customerID, tokenID)
	return t.Request.Get(url, queryParams, extraHeaders)
}

// All fetches collection of tokens for associated with a customer having the given customerID.
func (t *Token) All(customerID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/tokens", constants.CUSTOMER_URL, customerID)
	return t.Request.Get(url, queryParams, extraHeaders)
}

// Delete deletes a token having the given tokenID associated with a customer having the given customerID.
func (t *Token) Delete(customerID string, tokenID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/tokens/%s", constants.CUSTOMER_URL, customerID, tokenID)
	return t.Request.Delete(url, queryParams, extraHeaders)
}
