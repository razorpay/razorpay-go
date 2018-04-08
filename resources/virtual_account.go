package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//VirtualAccount : Struct for handling the razorpay virtual account API
type VirtualAccount struct {
	Request *requests.Request
}

//All : Fetch all virtual accounts for a given merchant
func (v *VirtualAccount) All(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return v.Request.Get(constants.VIRTUAL_ACCOUNT_URL, data, options)
}

//Fetch : Fetch a specific virtual account specified by the ID
func (v *VirtualAccount) Fetch(id string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.VIRTUAL_ACCOUNT_URL, id)
	return v.Request.Get(url, data, options)
}

//Create : Create a virtual account
func (v *VirtualAccount) Create(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return v.Request.Post(constants.VIRTUAL_ACCOUNT_URL, data, options)
}

//Close : Close a virtual account specified by ID
func (v *VirtualAccount) Close(id string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.VIRTUAL_ACCOUNT_URL, id)
	if data == nil {
		data = make(map[string]interface{})
	}
	data["status"] = "closed"
	return v.Request.Patch(url, data, options)
}

//Payments : Fetch the list of payments for a given virtual account
func (v *VirtualAccount) Payments(id string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/payments", constants.VIRTUAL_ACCOUNT_URL, id)
	return v.Request.Get(url, data, options)
}
