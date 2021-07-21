package resources

import (
	"fmt"
	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Contact ...
type Contact struct {
	Request *requests.Request
}

// Create creates a new contact for the given data.
func (contact *Contact) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return contact.Request.Post(constants.CONTACT_URL, data, extraHeaders)
}

// All fetches multiple contacts for the given query params.
func (contact *Contact) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return contact.Request.Get(constants.CONTACT_URL, queryParams, extraHeaders)
}

// Fetch fetches a contact having the given contactID.
func (contact *Contact) Fetch(contactID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.CONTACT_URL, contactID)
	return contact.Request.Get(url, queryParams, extraHeaders)
}
