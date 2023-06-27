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
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.VIRTUAL_ACCOUNT_URL)
	return v.Request.Get(url, queryParams, extraHeaders)
}

// Fetch fetches a virtual account for the given virtualAccID.
func (v *VirtualAccount) Fetch(virtualAccID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.VIRTUAL_ACCOUNT_URL, virtualAccID)
	return v.Request.Get(url, queryParams, extraHeaders)
}

// Create creates a new virtual account for the given data.
func (v *VirtualAccount) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.VIRTUAL_ACCOUNT_URL)
	return v.Request.Post(url, data, extraHeaders)
}

// Close closes the virtual account having the given virtualAccID.
func (v *VirtualAccount) Close(virtualAccID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/close", constants.VERSION_V1, constants.VIRTUAL_ACCOUNT_URL, virtualAccID)
	if data == nil {
		data = make(map[string]interface{})
	}
	return v.Request.Post(url, data, extraHeaders)
}

// Payments fetches a collection of payments associated with the virtual account having the given virtualAccID.
func (v *VirtualAccount) Payments(virtualAccID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/payments", constants.VERSION_V1, constants.VIRTUAL_ACCOUNT_URL, virtualAccID)
	return v.Request.Get(url, queryParams, extraHeaders)
}

// Create add receiver to an existing virtual account by virtual Id and data
func (v *VirtualAccount) AddReceiver(virtualAccID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/receivers", constants.VERSION_V1, constants.VIRTUAL_ACCOUNT_URL, virtualAccID)
	return v.Request.Post(url, data, extraHeaders)
}

// Allowedpayer add allowed payers account detail by virtualId and data
func (v *VirtualAccount) AllowedPayer(virtualAccID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/allowed_payers", constants.VERSION_V1, constants.VIRTUAL_ACCOUNT_URL, virtualAccID)
	return v.Request.Post(url, data, extraHeaders)
}

// Delete delete allowed payers account details by virtualId and allowed Payer id
func (v *VirtualAccount) DeleteAllowedPayer(virtualAccID string, PayerId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/allowed_payers/%s", constants.VERSION_V1, constants.VIRTUAL_ACCOUNT_URL, virtualAccID, PayerId)
	return v.Request.Delete(url, queryParams, extraHeaders)
}
