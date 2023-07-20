package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Item ...
type Item struct {
	Request *requests.Request
}

// All fetches collection of items for the given queryParams.
func (item *Item) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.ITEM_URL)
	return item.Request.Get(url, queryParams, extraHeaders)
}

// Fetch fetches an item having the given itemID.
func (item *Item) Fetch(itemID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.ITEM_URL, itemID)
	return item.Request.Get(url, queryParams, extraHeaders)
}

// Create creates a new item for the given data.
func (item *Item) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.ITEM_URL)
	return item.Request.Post(url, data, extraHeaders)
}

// Update updates an item for the given data.
func (item *Item) Update(itemID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.ITEM_URL, itemID)
	return item.Request.Patch(url, data, extraHeaders)
}

// Delete deletes an item having the given itemID.
func (item *Item) Delete(itemID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.ITEM_URL, itemID)
	return item.Request.Delete(url, queryParams, extraHeaders)
}
