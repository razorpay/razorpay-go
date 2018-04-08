package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Token : Struct handling razorpay customer token api
type Token struct {
	Request *requests.Request
}

//Fetch : Fetch token for a given customer id and token id
func (t *Token) Fetch(customerID string, tokenID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/tokens/%s", constants.CUSTOMER_URL, customerID, tokenID)
	return t.Request.Get(url, data, options)
}

//All : Fetch all tokens for a customer specified by customer id
func (t *Token) All(customerID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/tokens", constants.CUSTOMER_URL, customerID)
	return t.Request.Get(url, data, options)
}

//Delete : Delete a specified token for a specific customer
func (t *Token) Delete(customerID string, tokenID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/tokens/%s", constants.CUSTOMER_URL, customerID, tokenID)
	return t.Request.Delete(url, data, options)
}
