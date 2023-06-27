package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Invoice ...
type Invoice struct {
	Request *requests.Request
}

// All fetches multiple invoices for the given queryParams.
func (inv *Invoice) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.INVOICE_URL)
	return inv.Request.Get(url, queryParams, extraHeaders)
}

// Fetch fetches an invoice having the given invoiceID.
func (inv *Invoice) Fetch(invoiceID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.INVOICE_URL, invoiceID)

	return inv.Request.Get(url, queryParams, extraHeaders)
}

// Create creates a new invoice for the given data.
func (inv *Invoice) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.INVOICE_URL)
	return inv.Request.Post(url, data, extraHeaders)
}

// Cancel cancels a invoice having the given invoiceID.
func (s *Invoice) Cancel(invoiceID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/cancel", constants.VERSION_V1, constants.INVOICE_URL, invoiceID)
	return s.Request.Post(url, queryParams, extraHeaders)
}

// Delete delete a draft invoice having the given invoiceID.
func (s *Invoice) Delete(invoiceID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.INVOICE_URL, invoiceID)
	return s.Request.Delete(url, queryParams, extraHeaders)
}

// Update updates the invoice having the given invoiceID.
func (s *Invoice) Update(invoiceID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.INVOICE_URL, invoiceID)
	return s.Request.Patch(url, data, extraHeaders)
}

// Issue issue an invoice having the given invoiceID.
func (s *Invoice) Issue(invoiceID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/issue", constants.VERSION_V1, constants.INVOICE_URL, invoiceID)
	return s.Request.Post(url, data, extraHeaders)
}

//Notify notify to customer via email or sms by the given invoiceID.
func (s *Invoice) Notify(invoiceID string, medium string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/notify_by/%s", constants.VERSION_V1, constants.INVOICE_URL, invoiceID, medium)
	return s.Request.Post(url, data, extraHeaders)
}

// CreateRegistrationLink creates a registration link
func (s *Invoice) CreateRegistrationLink(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s/subscription_registration/auth_links", constants.VERSION_V1)
	return s.Request.Post(url, data, extraHeaders)
}
