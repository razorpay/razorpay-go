package resources

import (
	"fmt"
	"net/url"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Product ...
type Webhook struct {
	Request *requests.Request
}

// Webhook creates a new webhook for the given data.
func (wh *Webhook) Create(accountId string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	if accountId != "" {
		url := fmt.Sprintf("/%s%s/%s%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), constants.WEBHOOK)
		return wh.Request.Post(url, data, extraHeaders)
	}
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.WEBHOOK)
	return wh.Request.Post(url, data, extraHeaders)
}

// Fetch retrieve and view the details of a webhook.
func (wh *Webhook) Fetch(webhookId string, accountId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s%s/%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), constants.WEBHOOK, url.PathEscape(webhookId))
	return wh.Request.Get(url, queryParams, extraHeaders)
}

// Edit updates the details of a webhook.
func (wh *Webhook) Edit(webhookId string, accountId string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	if accountId != "" {
		url := fmt.Sprintf("/%s%s/%s%s/%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), constants.WEBHOOK, url.PathEscape(webhookId))
		return wh.Request.Patch(url, data, extraHeaders)
	}
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.WEBHOOK, url.PathEscape(accountId))
	return wh.Request.Put(url, data, extraHeaders)
}

// All fetches collection of webhook.
func (wh *Webhook) All(accountId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	if accountId != "" {
		url := fmt.Sprintf("/%s%s/%s%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), constants.WEBHOOK)
		return wh.Request.Get(url, queryParams, extraHeaders)
	}
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.WEBHOOK)
	return wh.Request.Get(url, queryParams, extraHeaders)
}

// Delete deletes a webhook having the given invoiceID.
func (wh *Webhook) Delete(webhookId string, accountId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s%s/%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), constants.WEBHOOK, url.PathEscape(webhookId))
	return wh.Request.Delete(url, queryParams, extraHeaders)
}
