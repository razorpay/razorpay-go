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
	return inv.Request.Get(constants.INVOICE_URL, queryParams, extraHeaders)
}

// Fetch fetches an invoice having the given invoiceID.
func (inv *Invoice) Fetch(invoiceID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s", constants.INVOICE_URL, invoiceID)

	return inv.Request.Get(url, queryParams, extraHeaders)
}

// Create creates a new invoice for the given data.
func (inv *Invoice) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return inv.Request.Post(constants.INVOICE_URL, data, extraHeaders)
}

// Cancel cancels a invoice having the given invoiceID.
func (s *Invoice) Cancel(invoiceID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/cancel", constants.INVOICE_URL, invoiceID)
	return s.Request.Post(url, queryParams, extraHeaders)
}

// Delete delete a draft invoice having the given invoiceID.
func (s *Invoice) Delete(invoiceID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.INVOICE_URL, invoiceID)
	return s.Request.Delete(url, queryParams, extraHeaders)
}

// Update updates the invoice having the given invoiceID.
func (s *Invoice) Update(invoiceID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
    url := fmt.Sprintf("%s/%s", constants.INVOICE_URL, invoiceID)
    return s.Request.Patch(url, data, extraHeaders)
}

// Issue issue an invoice having the given invoiceID.
func (s *Invoice) Issue(invoiceID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
    url := fmt.Sprintf("%s/%s/issue", constants.INVOICE_URL, invoiceID)
    return s.Request.Post(url, data, extraHeaders)
}

//Notify notify to customer via email or sms by the given invoiceID.
func (s *Invoice) Notify(invoiceID string, medium string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
    url := fmt.Sprintf("%s/%s/notify_by/%s", constants.INVOICE_URL, invoiceID, medium)
	return s.Request.Post(url, data, extraHeaders)
}

// CreateRegistrationLink creates a registration link
func (s *Invoice) CreateRegistrationLink(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := "/subscription_registration/auth_links"
	return s.Request.Post(url, data, extraHeaders)
}
