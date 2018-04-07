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

//Fetch ...
func (cust *Customer) Fetch(customerID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.CUSTOMER_URL, customerID)
	return cust.Request.Get(url, data, options)
}

//Create ...
func (cust *Customer) Create(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return cust.Request.Post(constants.CUSTOMER_URL, data, options)
}

//Edit ...
func (cust *Customer) Edit(customerID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s", constants.CUSTOMER_URL, customerID)

	return cust.Request.Post(url, data, options)
}
