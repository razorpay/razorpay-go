package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Orders ...
type Orders struct {
	Request *requests.Request
}

//All ...
func (orders *Orders) All(data map[string]interface{}, options map[string]string) ([]byte, error) {
	return orders.Request.Get(constants.ORDER_URL, data, options)
}

//Fetch ...
func (orders *Orders) Fetch(orderID string, data map[string]interface{}, options map[string]string) ([]byte, error) {

	url := fmt.Sprintf("%s/%s", constants.ORDER_URL, orderID)
	return orders.Request.Get(url, data, options)
}

//Create ...
func (orders *Orders) Create(data map[string]interface{}, options map[string]string) ([]byte, error) {
	return orders.Request.Post(constants.ORDER_URL, data, options)
}

//Edit ...
func (orders *Orders) Edit(orderID string, data map[string]interface{}, options map[string]string) ([]byte, error) {

	url := fmt.Sprintf("%s/%s", constants.ORDER_URL, orderID)
	return orders.Request.Put(url, data, options)
}

//Payments ...
func (orders *Orders) Payments(orderID string, data map[string]interface{}, options map[string]string) ([]byte, error) {

	url := fmt.Sprintf("%s/%s/payments", constants.ORDER_URL, orderID)
	return orders.Request.Get(url, data, options)
}
