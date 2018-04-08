package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Plan : Struct to handle the razorpay Plan API
type Plan struct {
	Request *requests.Request
}

//Create : Create a plan
func (plan *Plan) Create(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return plan.Request.Post(constants.PLAN_URL, data, options)
}

//Fetch : Fetch a plan by ID
func (plan *Plan) Fetch(id string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.PLAN_URL, id)
	return plan.Request.Get(url, data, options)
}

//All : Fetch all plans for a merchant
func (plan *Plan) All(data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	return plan.Request.Get(constants.PLAN_URL, data, options)
}
