package resources

import (
	"fmt"

	"gitlab.com/venkatvghub/razorpay-go/razorpay/constants"
	"gitlab.com/venkatvghub/razorpay-go/razorpay/requests"
)

//Payments ...
type Payments struct {
	Request *requests.Request
}

//All ...
func (p *Payments) All(data map[string]interface{}, options map[string]string) ([]byte, error) {
	return p.Request.Get(constants.PAYMENT_URL, data, options)
}

//Fetch ...
func (p *Payments) Fetch(id string, data map[string]interface{}, options map[string]string) ([]byte, error) {

	url := fmt.Sprintf("%s/%s", constants.PAYMENT_URL, id)

	return p.Request.Get(url, data, options)
}

//Capture ...
func (p *Payments) Capture(id string, amount int, data map[string]interface{}, options map[string]string) ([]byte, error) {

	url := fmt.Sprintf("%s/%s/capture", constants.PAYMENT_URL, id)
	// Amount should be in paisa

	data["amount"] = amount

	return p.Request.Post(url, data, options)
}

//Refund ...
func (p *Payments) Refund(id string, amount int, data map[string]interface{}, options map[string]string) ([]byte, error) {

	url := fmt.Sprintf("%s/%s/refund", constants.PAYMENT_URL, id)
	// Amount should be in paisa

	data["amount"] = amount

	return p.Request.Post(url, data, options)
}
