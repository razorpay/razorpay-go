package resources

import (
	"fmt"

	"gitlab.com/venkatvghub/razorpay-go/razorpay/constants"
	"gitlab.com/venkatvghub/razorpay-go/razorpay/requests"
)

//Refunds ...
type Refunds struct {
	Request *requests.Request
}

func (refunds *Refunds) all(data map[string]interface{}, options map[string]string) ([]byte, error) {
	return refunds.Request.Get(constants.REFUND_URL, data, options)
}

func (refunds *Refunds) fetch(refundID string, data map[string]interface{}, options map[string]string) ([]byte, error) {

	url := fmt.Sprintf("%s/%s", constants.REFUND_URL, refundID)
	return refunds.Request.Get(url, data, options)
}

func (refunds *Refunds) create(data map[string]interface{}, options map[string]string) ([]byte, error) {
	return refunds.Request.Post(constants.REFUND_URL, data, options)
}
