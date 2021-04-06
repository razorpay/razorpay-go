package resources

import (
"fmt"
"github.com/razorpay/razorpay-go/constants"
"github.com/razorpay/razorpay-go/requests"
)

//FundAccount ...
type FundAccount struct {
	Request *requests.Request
}

// Create creates a new FundAccount for the given data.
func (fa *FundAccount) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return fa.Request.Post(constants.FUND_ACCOUNT_URL, data, extraHeaders)
}

// All fetches multiple FundAccounts for the given query params.
func (fa *FundAccount) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return fa.Request.Get(constants.FUND_ACCOUNT_URL, queryParams, extraHeaders)
}

// Fetch fetches a FundAccount having the given fundAccountID.
func (fa *FundAccount) Fetch(fundAccountID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.FUND_ACCOUNT_URL, fundAccountID)
	return fa.Request.Get(url, queryParams, extraHeaders)
}

