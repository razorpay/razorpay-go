package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

// Payout represents the payout resource and provides methods to interact with Razorpay's Payout API
// Payout contains a Request object that handles the actual HTTP requests to the Razorpay API

// Refer documents/payout.md for more details

type Payout struct {
	// Request is a pointer to the requests.Request struct that handles HTTP communication
	Request *requests.Request
}

// Create is a function to create a new payout in Razorpay
// Returns map[string]interface{} containing payout details and error if any
func (payout *Payout) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.PAYOUT_URL)
	return payout.Request.Post(url, data, extraHeaders)
}

// All is a function to fetch all payouts from Razorpay
// Returns map[string]interface{} containing list of payouts and error if any
func (payout *Payout) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.PAYOUT_URL)
	return payout.Request.Get(url, queryParams, extraHeaders)
}

// Fetch is a function to get details of a specific payout from Razorpay
// Returns map[string]interface{} containing payout details and error if any
func (payout *Payout) Fetch(payoutID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.PAYOUT_URL, payoutID)
	return payout.Request.Get(url, queryParams, extraHeaders)
}
