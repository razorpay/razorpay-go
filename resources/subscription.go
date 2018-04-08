package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Subscription ...
type Subscription struct {
	Request *requests.Request
}

//All ...
func (s *Subscription) All(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return s.Request.Get(constants.SUBSCRIPTION_URL, data, options)
}

//Fetch ...
func (s *Subscription) Fetch(id string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.SUBSCRIPTION_URL, id)
	return s.Request.Get(url, data, options)
}

//Create ...
func (s *Subscription) Create(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return s.Request.Post(constants.SUBSCRIPTION_URL, data, options)
}

//Cancel ...
func (s *Subscription) Cancel(id string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/cancel", constants.SUBSCRIPTION_URL, id)
	return s.Request.Post(url, data, options)
}

//CreateAddon ...
func (s *Subscription) CreateAddon(id string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/addons", constants.SUBSCRIPTION_URL, id)
	return s.Request.Post(url, data, options)
}
