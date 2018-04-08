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

//All ...
func (t *Transfer) All(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return t.Request.Get(constants.TRANSFER_URL, data, options)
}

//Fetch ...
func (t *Transfer) Fetch(transferID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.TRANSFER_URL, transferID)
	return t.Request.Get(url, data, options)
}

//Create ...
func (t *Token) Create(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return t.Request.Post(constants.TRANSFER_URL, data, options)
}

//Edit ...
func (t *Token) Edit(transferID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.TRANSFER_URL, transferID)
	return t.Request.Patch(url, data, options)
}

//Reverse ...
func (t *Token) Reverse(transferID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/reversals", constants.TRANSFER_URL, transferID)
	return t.Request.Post(url, data, options)
}

//Reversals ...
func (t *Token) Reversals(transferID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/reversals", constants.TRANSFER_URL, transferID)
	return t.Request.Get(url, data, options)
}
