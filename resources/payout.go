package resources

import (
	"fmt"
	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Contact ...
type Payout struct {
	Request *requests.Request
}

// Create creates a new payout for the given data.
func (p *Payout) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return p.Request.Post(constants.PAYOUT_URL, data, extraHeaders)
}

// All fetches multiple payouts for the given query params.
func (p *Payout) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return p.Request.Get(constants.PAYOUT_URL, queryParams, extraHeaders)
}

// Fetch fetches a payout having the given payoutID.
func (p *Payout) Fetch(payoutID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.PAYOUT_URL, payoutID)
	return p.Request.Get(url, queryParams, extraHeaders)
}

