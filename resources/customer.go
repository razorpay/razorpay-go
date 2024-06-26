package resources

import (
	"fmt"
	"net/url"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Customer ...
type Customer struct {
	Request *requests.Request
}

// Fetch fetches customer having the given cutomerID.
func (cust *Customer) Fetch(customerID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.CUSTOMER_URL, url.PathEscape(customerID))
	return cust.Request.Get(url, queryParams, extraHeaders)
}

// Create creates a new customer for the given data.
func (cust *Customer) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.CUSTOMER_URL)
	return cust.Request.Post(url, data, extraHeaders)
}

// Edit updates the customer having the given customerID.
func (cust *Customer) Edit(customerID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.CUSTOMER_URL, url.PathEscape(customerID))

	return cust.Request.Put(url, data, extraHeaders)
}

// All fetches collection of customer for the given queryParams.
func (cust *Customer) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.CUSTOMER_URL)
	return cust.Request.Get(url, queryParams, extraHeaders)
}

func (cust *Customer) AddBankAccount(customerID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/bank_account", constants.VERSION_V1, constants.CUSTOMER_URL, url.PathEscape(customerID))
	return cust.Request.Post(url, data, extraHeaders)
}

func (cust *Customer) DeleteBankAccount(customerID string, bankAccountId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/bank_account/%s", constants.VERSION_V1, constants.CUSTOMER_URL, url.PathEscape(customerID), url.PathEscape(bankAccountId))
	return cust.Request.Delete(url, queryParams, extraHeaders)
}

func (cust *Customer) RequestEligibilityCheck(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/eligibility", constants.VERSION_V1, constants.CUSTOMER_URL)
	return cust.Request.Post(url, data, extraHeaders)
}

func (cust *Customer) FetchEligibility(eligibilityId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/eligibility/%s", constants.VERSION_V1, constants.CUSTOMER_URL, url.PathEscape(eligibilityId))
	return cust.Request.Get(url, queryParams, extraHeaders)
}
