package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Customers ...
type Customers struct {
	Request *requests.Request
}

//Fetch ...
func (cust *Customers) Fetch(customerID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.CUSTOMER_URL, customerID)
	return cust.Request.Get(url, data, options)
}

//Create ...
func (cust *Customers) Create(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return cust.Request.Post(constants.CUSTOMER_URL, data, options)
}

//Edit ...
func (cust *Customers) Edit(customerID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s", constants.CUSTOMER_URL, customerID)

	return cust.Request.Post(url, data, options)
}
