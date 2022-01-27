package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Refund ...
type Refund struct {
	Request *requests.Request
}

// All fetches colelction of Refund for the given queryParams
func (refund *Refund) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return refund.Request.Get(constants.RefundURL, queryParams, extraHeaders)
}

// Fetch fetches the Refund having the given refundID.
func (refund *Refund) Fetch(refundID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s", constants.RefundURL, refundID)
	return refund.Request.Get(url, queryParams, extraHeaders)
}

// Create creates a new Refund for the given data.
func (refund *Refund) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return refund.Request.Post(constants.RefundURL, data, extraHeaders)
}
