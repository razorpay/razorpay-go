package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

// Payout represents the payout resource
type Payout struct {
	Request *requests.Request
}

// Create creates a new payout for the given data
func (payout *Payout) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.PAYOUT_URL)
	return payout.Request.Post(url, data, extraHeaders)
}

// All fetches multiple payouts for the given query params
func (payout *Payout) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.PAYOUT_URL)
	return payout.Request.Get(url, queryParams, extraHeaders)
}

// Fetch fetches a payout having the given payoutID
func (payout *Payout) Fetch(payoutID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.PAYOUT_URL, payoutID)
	return payout.Request.Get(url, queryParams, extraHeaders)
}
