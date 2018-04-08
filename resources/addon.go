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

//Fetch ...
func (addon *Addon) Fetch(addonID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.ADDON_URL, addonID)
	return addon.Request.Get(url, data, options)
}

//Delete ...
func (addon *Addon) Delete(addonID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.ADDON_URL, addonID)
	return addon.Request.Delete(url, data, options)
}
