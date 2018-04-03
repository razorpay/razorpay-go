package resources

import (
	"fmt"

	"gitlab.com/venkatvghub/razorpay-go/razorpay/constants"
	"gitlab.com/venkatvghub/razorpay-go/razorpay/requests"
)

//Orders ...
type Orders struct {
	Request *requests.Request
}

func (orders *Orders) all(data map[string]interface{}, options map[string]string) ([]byte, error) {
	return orders.Request.Get(constants.ORDER_URL, data, options)
}

func (orders *Orders) fetch(orderID string, data map[string]interface{}, options map[string]string) ([]byte, error) {

	url := fmt.Sprintf("%s/%s", constants.ORDER_URL, orderID)
	return orders.Request.Get(url, data, options)
}

func (orders *Orders) create(data map[string]interface{}, options map[string]string) ([]byte, error) {
	return orders.Request.Post(constants.ORDER_URL, data, options)
}

func (orders *Orders) edit(orderID string, data map[string]interface{}, options map[string]string) ([]byte, error) {

	url := fmt.Sprintf("%s/%s", constants.ORDER_URL, orderID)
	return orders.Request.Put(url, data, options)
}

func (orders *Orders) payments(orderID string, data map[string]interface{}, options map[string]string) ([]byte, error) {

	url := fmt.Sprintf("%s/%s/payments", constants.ORDER_URL, orderID)
	return orders.Request.Get(url, data, options)
}
