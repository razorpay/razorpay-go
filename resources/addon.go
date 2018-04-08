//Package resources : holds the list of individual razorpay client apis
package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Addon : Struct for handling the Razorpay Addon API
type Addon struct {
	Request *requests.Request
}

//Fetch : method to fetch a given addon by ID
func (addon *Addon) Fetch(addonID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.ADDON_URL, addonID)
	return addon.Request.Get(url, data, options)
}

//Delete : method to delete an addon by ID
func (addon *Addon) Delete(addonID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.ADDON_URL, addonID)
	return addon.Request.Delete(url, data, options)
}
