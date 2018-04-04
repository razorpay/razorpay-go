package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Refunds ...
type Refunds struct {
	Request *requests.Request
}

//All ...
func (refunds *Refunds) All(data map[string]interface{}, options map[string]string) ([]byte, error) {
	return refunds.Request.Get(constants.REFUND_URL, data, options)
}

//Fetch ...
func (refunds *Refunds) Fetch(refundID string, data map[string]interface{}, options map[string]string) ([]byte, error) {

	url := fmt.Sprintf("%s/%s", constants.REFUND_URL, refundID)
	return refunds.Request.Get(url, data, options)
}

//Create ...
func (refunds *Refunds) Create(data map[string]interface{}, options map[string]string) ([]byte, error) {
	return refunds.Request.Post(constants.REFUND_URL, data, options)
}
