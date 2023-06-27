package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Plan ...
type Plan struct {
	Request *requests.Request
}

// Create creates a new plan for the given data.
func (plan *Plan) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.PLAN_URL)
	return plan.Request.Post(url, data, extraHeaders)
}

// Fetch fetches the plan entity having the given planID.
func (plan *Plan) Fetch(planID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.PLAN_URL, planID)
	return plan.Request.Get(url, queryParams, extraHeaders)
}

// All fetches collection of plans for the given queryParams.
func (plan *Plan) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.PLAN_URL)
	return plan.Request.Get(url, queryParams, extraHeaders)
}
