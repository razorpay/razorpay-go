package resources

import (
	"fmt"
	"net/url"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Product ...
type Product struct {
	Request *requests.Request
}

func (prod *Product) RequestProductConfiguration(accountId string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), constants.PRODUCT_URL)
	return prod.Request.Post(url, data, extraHeaders)
}

func (prod *Product) Fetch(accountId string, productId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s%s/%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), constants.PRODUCT_URL, url.PathEscape(productId))
	return prod.Request.Get(url, queryParams, extraHeaders)
}

func (prod *Product) Edit(accountId string, productId string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s%s/%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), constants.PRODUCT_URL, url.PathEscape(productId))
	return prod.Request.Patch(url, data, extraHeaders)
}

func (prod *Product) FetchTnc(productname string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s%s", constants.VERSION_V2, constants.PRODUCT_URL, url.PathEscape(productname), constants.TNC)
	return prod.Request.Get(url, queryParams, extraHeaders)
}
