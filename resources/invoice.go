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
	return inv.Request.Get(constants.InvoiceURL, queryParams, extraHeaders)
}

// Fetch fetches an invoice having the given invoiceID.
func (inv *Invoice) Fetch(invoiceID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s", constants.InvoiceURL, invoiceID)

	return inv.Request.Get(url, queryParams, extraHeaders)
}

// Create creates a new invoice for the given data.
func (inv *Invoice) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return inv.Request.Post(constants.InvoiceURL, data, extraHeaders)
}
