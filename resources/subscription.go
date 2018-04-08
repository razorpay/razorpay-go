package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Subscription : Struct for handling razorpay subscription api
type Subscription struct {
	Request *requests.Request
}

//All : Fetch all subscriptions for a merchant
func (s *Subscription) All(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return s.Request.Get(constants.SUBSCRIPTION_URL, data, options)
}

//Fetch : Fetch a subscription by ID
func (s *Subscription) Fetch(id string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.SUBSCRIPTION_URL, id)
	return s.Request.Get(url, data, options)
}

//Create : Create a subscription
func (s *Subscription) Create(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return s.Request.Post(constants.SUBSCRIPTION_URL, data, options)
}

//Cancel : Cancel a subscription
func (s *Subscription) Cancel(id string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/cancel", constants.SUBSCRIPTION_URL, id)
	return s.Request.Post(url, data, options)
}

//CreateAddon : Create an addon for a given subscription
func (s *Subscription) CreateAddon(id string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/addons", constants.SUBSCRIPTION_URL, id)
	return s.Request.Post(url, data, options)
}
