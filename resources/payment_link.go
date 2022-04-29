package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//PaymentLink ...
type PaymentLink struct {
	Request *requests.Request
}

// All fetches multiple PaymentLink for the given queryParams.
func (pl *PaymentLink) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return pl.Request.Get(constants.PaymentLink_URL, queryParams, extraHeaders)
}

// Fetch fetches an PaymentLink having the given paymentLinkID.
func (pl *PaymentLink) Fetch(paymentLinkID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s", constants.PaymentLink_URL, paymentLinkID)

	return pl.Request.Get(url, queryParams, extraHeaders)
}

// Create creates a new PaymentLink  for the given data.
func (pl *PaymentLink) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return pl.Request.Post(constants.PaymentLink_URL, data, extraHeaders)
}

// Cancel cancels paymentLink
func (pl *PaymentLink) Cancel(paymentLinkID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/cancel", constants.PaymentLink_URL, paymentLinkID)
	return pl.Request.Post(url, data, extraHeaders)
}

// NotifyBy Send/re-send notification by given medium
func (pl *PaymentLink) NotifyBy(paymentLinkID string, medium string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/notify_by/%s", constants.PaymentLink_URL, paymentLinkID, medium)
	return pl.Request.Post(url, data, extraHeaders)
}

// Edit updates the paymentLink for the given data.
func (pl *PaymentLink) Update(paymentLinkID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.PaymentLink_URL, paymentLinkID)
	return pl.Request.Patch(url, data, extraHeaders)
}
