package resources

import (
	"fmt"

	"gitlab.com/venkatvghub/razorpay-go/razorpay/constants"
	"gitlab.com/venkatvghub/razorpay-go/razorpay/requests"
)

//Customers ...
type Customers struct {
	Request *requests.Request
}

func (cust *Customers) fetch(customerID string, data map[string]interface{}, options map[string]string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", constants.CUSTOMER_URL, customerID)
	return cust.Request.Get(url, data, options)
}

func (cust *Customers) create(data map[string]interface{}, options map[string]string) ([]byte, error) {
	return cust.Request.Post(constants.CUSTOMER_URL, data, options)
}

func (cust *Customers) edit(customerID string, data map[string]interface{}, options map[string]string) ([]byte, error) {

	url := fmt.Sprintf("%s/%s", constants.CUSTOMER_URL, customerID)

	return cust.Request.Post(url, data, options)
}
