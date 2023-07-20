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

// RequestProductConfiguration creates a request for a product configuration for the given data.
func (prod *Product) RequestProductConfiguration(accountId string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), constants.PRODUCT_URL)
	return prod.Request.Post(url, data, extraHeaders)
}

// Fetch fetches the details of a product.
func (prod *Product) Fetch(accountId string, productId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s%s/%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), constants.PRODUCT_URL, url.PathEscape(productId))
	return prod.Request.Get(url, queryParams, extraHeaders)
}

// Edit updates the product's configuration.
func (prod *Product) Edit(accountId string, productId string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s%s/%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), constants.PRODUCT_URL, url.PathEscape(productId))
	return prod.Request.Patch(url, data, extraHeaders)
}

// FetchTnc fetches the terms and conditions.
func (prod *Product) FetchTnc(productname string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s%s", constants.VERSION_V2, constants.PRODUCT_URL, url.PathEscape(productname), constants.TNC)
	return prod.Request.Get(url, queryParams, extraHeaders)
}
