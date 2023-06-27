package resources

import (
	"fmt"
	"net/url"
	"os"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Account ...
type Account struct {
	Request *requests.Request
}

type FileUpload struct {
	File   *os.File
	Fields map[string]string
}

func (acc *Account) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V2, constants.ACCOUNT_URL)
	return acc.Request.Post(url, data, extraHeaders)
}

func (acc *Account) Fetch(accountId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId))
	return acc.Request.Get(url, queryParams, extraHeaders)
}

func (acc *Account) Edit(accountId string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId))
	return acc.Request.Patch(url, data, extraHeaders)
}

func (acc *Account) Delete(accountId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId))
	return acc.Request.Delete(url, queryParams, extraHeaders)
}

func (acc *Account) UploadAccountDoc(accountId string, params requests.FileUploadParams) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/documents", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId))
	return acc.Request.File(url, params)
}

func (acc *Account) FetchAccountDoc(accountId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/documents", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId))
	return acc.Request.Get(url, queryParams, extraHeaders)
}
