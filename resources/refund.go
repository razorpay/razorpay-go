package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Refund : Struct to handle razorpay refunds api
type Refund struct {
	Request *requests.Request
}

//All : Fetch all refunds for a merchant
func (refund *Refund) All(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return refund.Request.Get(constants.REFUND_URL, data, options)
}

//Fetch : Fetch a refund by ID
func (refund *Refund) Fetch(refundID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s", constants.REFUND_URL, refundID)
	return refund.Request.Get(url, data, options)
}

//Create : Create a refund
func (refund *Refund) Create(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return refund.Request.Post(constants.REFUND_URL, data, options)
}
