package resources

import (
	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//FundAccount ...
type FundAccount struct {
	Request *requests.Request
}

// All fetches collection of fund accounts for the given queryParams.
func (fa *FundAccount) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return fa.Request.Get(constants.FUND_ACCOUNT_URL, queryParams, extraHeaders)
}

// Create creates a fund account for the given data.
func (fa *FundAccount) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return fa.Request.Post(constants.FUND_ACCOUNT_URL, data, extraHeaders)
}
