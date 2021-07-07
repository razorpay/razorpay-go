package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//VirtualAccount ...
type VirtualAccount struct {
	Request *requests.Request
}

// All fetches collection of virtual account for the given queryParams.
func (v *VirtualAccount) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return v.Request.Get(constants.VIRTUAL_ACCOUNT_URL, queryParams, extraHeaders)
}

// Fetch fetches a virtual account for the given virtualAccID.
func (v *VirtualAccount) Fetch(virtualAccID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.VIRTUAL_ACCOUNT_URL, virtualAccID)
	return v.Request.Get(url, queryParams, extraHeaders)
}

// Create creates a new virtual account for the given data.
func (v *VirtualAccount) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return v.Request.Post(constants.VIRTUAL_ACCOUNT_URL, data, extraHeaders)
}

// Close closes the virtual account having the given virtualAccID.
func (v *VirtualAccount) Close(virtualAccID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/close", constants.VIRTUAL_ACCOUNT_URL, virtualAccID)
	if data == nil {
		data = make(map[string]interface{})
	}
	return v.Request.Post(url, data, extraHeaders)
}

// Payments fetches a collection of payments associated with the virtual account having the given virtualAccID.
func (v *VirtualAccount) Payments(virtualAccID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/payments", constants.VIRTUAL_ACCOUNT_URL, virtualAccID)
	return v.Request.Get(url, queryParams, extraHeaders)
}
