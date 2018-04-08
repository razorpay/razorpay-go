package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Payment : Struct for handling the Razorpay Payments API
type Payment struct {
	Request *requests.Request
}

//All : Fetch all payments for merchant
func (p *Payment) All(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return p.Request.Get(constants.PAYMENT_URL, data, options)
}

//Fetch : Fetch a specific payment by ID
func (p *Payment) Fetch(id string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s", constants.PAYMENT_URL, id)

	return p.Request.Get(url, data, options)
}

//Capture : Capture the payment specified by the ID
func (p *Payment) Capture(id string, amount int, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s/capture", constants.PAYMENT_URL, id)
	// Amount should be in paisa
	if data == nil {
		data = make(map[string]interface{})
	}
	data["amount"] = amount

	return p.Request.Post(url, data, options)
}

//Refund : Refund a payment specified by the ID
func (p *Payment) Refund(id string, amount int, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s/refund", constants.PAYMENT_URL, id)
	// Amount should be in paisa
	if data == nil {
		data = make(map[string]interface{})
	}
	data["amount"] = amount

	return p.Request.Post(url, data, options)
}

//Transfer : Create a transfer for a payment specified by the ID
func (p *Payment) Transfer(id string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/transfers", constants.PAYMENT_URL, id)
	return p.Request.Post(url, data, options)
}

//Transfers : Get all transfers for a given payment ID
func (p *Payment) Transfers(id string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/transfers", constants.PAYMENT_URL, id)
	return p.Request.Get(url, data, options)
}

//BankTransfer : Get banktransfer for a given ID
func (p *Payment) BankTransfer(id string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/bank_transfer", constants.PAYMENT_URL, id)
	return p.Request.Get(url, data, options)
}
