package resources

import (
	"fmt"
	"net/url"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Refund ...
type Refund struct {
	Request *requests.Request
}

// All fetches colelction of Refund for the given queryParams
func (refund *Refund) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.REFUND_URL)
	return refund.Request.Get(url, queryParams, extraHeaders)
}

// Fetch fetches the Refund having the given refundID.
func (refund *Refund) Fetch(refundID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.REFUND_URL, url.PathEscape(refundID))
	return refund.Request.Get(url, queryParams, extraHeaders)
}

// Create creates a new Refund for the given data.
func (refund *Refund) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.REFUND_URL)
	return refund.Request.Post(url, data, extraHeaders)
}

// Update updates the Refund having the given refundID.
func (refund *Refund) Update(refundID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.REFUND_URL, url.PathEscape(refundID))
	return refund.Request.Patch(url, queryParams, extraHeaders)
}
