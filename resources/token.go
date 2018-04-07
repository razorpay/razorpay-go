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

//Fetch ...
func (t *Token) Fetch(customerID string, tokenID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/tokens/%s", constants.CUSTOMER_URL, customerID, tokenID)
	return t.Request.Get(url, data, options)
}

//All ..
func (t *Token) All(customerID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/tokens", constants.CUSTOMER_URL, customerID)
	return t.Request.Get(url, data, options)
}

//Delete ...
func (t *Token) Delete(customerID string, tokenID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/tokens/%s", constants.CUSTOMER_URL, customerID, tokenID)
	return t.Request.Delete(url, data, options)
}
