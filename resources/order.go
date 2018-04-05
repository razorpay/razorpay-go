package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Order ...
type Order struct {
	Request *requests.Request
}

//All ...
func (order *Order) All(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return order.Request.Get(constants.ORDER_URL, data, options)
}

//Fetch ...
func (order *Order) Fetch(orderID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s", constants.ORDER_URL, orderID)
	return order.Request.Get(url, data, options)
}

//Create ...
func (order *Order) Create(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return order.Request.Post(constants.ORDER_URL, data, options)
}

//Edit ...
func (order *Order) Edit(orderID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s", constants.ORDER_URL, orderID)
	return order.Request.Put(url, data, options)
}

//Payments ...
func (order *Order) Payments(orderID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s/payments", constants.ORDER_URL, orderID)
	return order.Request.Get(url, data, options)
}
