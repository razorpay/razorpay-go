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

//All ...
func (v *VirtualAccount) All(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return v.Request.Get(constants.VIRTUAL_ACCOUNT_URL, data, options)
}

//Fetch ...
func (v *VirtualAccount) Fetch(id string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.VIRTUAL_ACCOUNT_URL, id)
	return v.Request.Get(url, data, options)
}

//Create ...
func (v *VirtualAccount) Create(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return v.Request.Post(constants.VIRTUAL_ACCOUNT_URL, data, options)
}

//Close ...
func (v *VirtualAccount) Close(id string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.VIRTUAL_ACCOUNT_URL, id)
	if data == nil {
		data = make(map[string]interface{})
	}
	data["status"] = "closed"
	return v.Request.Patch(url, data, options)
}

//Payments ...
func (v *VirtualAccount) Payments(id string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/payments", constants.VIRTUAL_ACCOUNT_URL, id)
	return v.Request.Get(url, data, options)
}
