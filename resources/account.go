package resources

import (
	"fmt"
	"net/url"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Account ...
type Account struct {
	Request *requests.Request
}

// Create creates a new account for the given data.
func (acc *Account) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V2, constants.ACCOUNT_URL)
	return acc.Request.Post(url, data, extraHeaders)
}

// Fetch fetches account having the given accountId.
func (acc *Account) Fetch(accountId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId))
	return acc.Request.Get(url, queryParams, extraHeaders)
}

// Edit updates the account having the given accountId.
func (acc *Account) Edit(accountId string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId))
	return acc.Request.Patch(url, data, extraHeaders)
}

// Delete deletes the account having the given accountId.
func (acc *Account) Delete(accountId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId))
	return acc.Request.Delete(url, queryParams, extraHeaders)
}

// UploadAccountDoc upload the account documents.
func (acc *Account) UploadAccountDoc(accountId string, params requests.FileUploadParams, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/documents", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId))
	return acc.Request.File(url, params, extraHeaders)
}

// FetchAccountDoc fetches the account documents.
func (acc *Account) FetchAccountDoc(accountId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/documents", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId))
	return acc.Request.Get(url, queryParams, extraHeaders)
}
