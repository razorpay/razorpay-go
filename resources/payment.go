package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Payment ...
type Payment struct {
	Request *requests.Request
}

//All ...
func (p *Payment) All(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return p.Request.Get(constants.PAYMENT_URL, data, options)
}

//Fetch ...
func (p *Payment) Fetch(id string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s", constants.PAYMENT_URL, id)

	return p.Request.Get(url, data, options)
}

//Capture ...
func (p *Payment) Capture(id string, amount int, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s/capture", constants.PAYMENT_URL, id)
	// Amount should be in paisa
	if data == nil {
		data = make(map[string]interface{})
	}
	data["amount"] = amount

	return p.Request.Post(url, data, options)
}

//Refund ...
func (p *Payment) Refund(id string, amount int, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s/refund", constants.PAYMENT_URL, id)
	// Amount should be in paisa
	if data == nil {
		data = make(map[string]interface{})
	}
	data["amount"] = amount

	return p.Request.Post(url, data, options)
}

//Transfer ...
func (p *Payment) Transfer(id string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/transfers", constants.PAYMENT_URL, id)
	return p.Request.Post(url, data, options)
}

//Transfers ...
func (p *Payment) Transfers(id string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/transfers", constants.PAYMENT_URL, id)
	return p.Request.Get(url, data, options)
}

//BankTransfer ...
func (p *Payment) BankTransfer(id string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/bank_transfer", constants.PAYMENT_URL, id)
	return p.Request.Get(url, data, options)
}
