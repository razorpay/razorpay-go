package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Transfer ...
type Transfer struct {
	Request *requests.Request
}

// All fetches collection of transfer for the given queryParams.
func (t *Transfer) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.TRANSFER_URL)
	return t.Request.Get(url, queryParams, extraHeaders)
}

// Fetch fetches a transfer having the given transferID.
func (t *Transfer) Fetch(transferID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.TRANSFER_URL, transferID)
	return t.Request.Get(url, queryParams, extraHeaders)
}

// Create creates a new transfer for the given data.
func (t *Transfer) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.TRANSFER_URL)
	return t.Request.Post(url, data, extraHeaders)
}

// Edit edits the transfer having the given transferID.
func (t *Transfer) Edit(transferID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.TRANSFER_URL, transferID)
	return t.Request.Patch(url, data, extraHeaders)
}

// Reverse reverses the transfer having the given transferID.
func (t *Transfer) Reverse(transferID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/reversals", constants.VERSION_V1, constants.TRANSFER_URL, transferID)
	return t.Request.Post(url, data, extraHeaders)
}

// Reversals fetches a collection of transfer associated with the given transferID.
func (t *Transfer) Reversals(transferID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/reversals", constants.VERSION_V1, constants.TRANSFER_URL, transferID)
	return t.Request.Get(url, queryParams, extraHeaders)
}
