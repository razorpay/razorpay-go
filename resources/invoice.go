package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Invoice ...
type Invoice struct {
	Request *requests.Request
}

//All ...
func (inv *Invoice) All(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return inv.Request.Get(constants.INVOICE_URL, data, options)
}

//Fetch ...
func (inv *Invoice) Fetch(id string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s", constants.INVOICE_URL, id)

	return inv.Request.Get(url, data, options)
}

//Create ...
func (inv *Invoice) Create(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return inv.Request.Post(constants.INVOICE_URL, data, options)
}
