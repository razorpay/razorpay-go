package resources

import (
	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

// Payout represents the payout resource
type Payout struct {
	Request *requests.Request
}

// All fetches multiple payouts for the given query params
func (payout *Payout) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := "/" + constants.VERSION_V1 + constants.PAYOUT_URL
	return payout.Request.Get(url, queryParams, extraHeaders)
}

// Fetch fetches a payout having the given payoutID
func (payout *Payout) Fetch(id string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := "/" + constants.VERSION_V1 + constants.PAYOUT_URL + "/" + id
	return payout.Request.Get(url, queryParams, extraHeaders)
}

// Create creates a new payout for the given data
func (payout *Payout) Create(req *PayoutRequest, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := "/" + constants.VERSION_V1 + constants.PAYOUT_URL

	// Convert PayoutRequest to map[string]interface{}
	dataMap := make(map[string]interface{})
	if req != nil {
		jsonData, err := json.Marshal(req)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(jsonData, &dataMap); err != nil {
			return nil, err
		}
	}
	return payout.Request.Post(url, dataMap, extraHeaders)
}
