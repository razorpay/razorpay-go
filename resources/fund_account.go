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

// All fetches collection of fund accounts for the given queryParams.
func (fa *FundAccount) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.FUND_ACCOUNT_URL)
	return fa.Request.Get(url, queryParams, extraHeaders)
}

// Create creates a fund account for the given data.
func (fa *FundAccount) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.FUND_ACCOUNT_URL)
	return fa.Request.Post(url, data, extraHeaders)
}
