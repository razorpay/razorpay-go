package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Addon ...
type Addon struct {
	Request *requests.Request
}

// Fetch fetches addon having the given addonID.
func (addon *Addon) Fetch(addonID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.ADDON_URL, addonID)
	return addon.Request.Get(url, queryParams, extraHeaders)
}

// Delete deletes the addon having the given addonID.
func (addon *Addon) Delete(addonID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.ADDON_URL, addonID)
	return addon.Request.Delete(url, queryParams, extraHeaders)
}

// All fetches collection of addon for the given queryParams.
func (addon *Addon) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.ADDON_URL)
	return addon.Request.Get(url, queryParams, extraHeaders)
}