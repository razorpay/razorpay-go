package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Customer ...
type Customer struct {
	Request *requests.Request
}

// Fetch fetches customer having the given cutomerID.
func (cust *Customer) Fetch(customerID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.CustomerURL, customerID)
	return cust.Request.Get(url, queryParams, extraHeaders)
}

// Create creates a new customer for the given data.
func (cust *Customer) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return cust.Request.Post(constants.CustomerURL, data, extraHeaders)
}

// Edit updates the customer having the given customerID.
func (cust *Customer) Edit(customerID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s", constants.CustomerURL, customerID)

	return cust.Request.Post(url, data, extraHeaders)
}
