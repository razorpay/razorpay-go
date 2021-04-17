package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Payment ...
type Payment struct {
	Request *requests.Request
}

// All fetches multiple payment entities for the given queryParams.
func (p *Payment) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return p.Request.Get(constants.PaymentURL, queryParams, extraHeaders)
}

// Fetch fetches the payment entity for the given paymentID.
func (p *Payment) Fetch(paymentID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s", constants.PaymentURL, paymentID)

	return p.Request.Get(url, queryParams, extraHeaders)
}

// Capture captures the payment having the given paymentID.
func (p *Payment) Capture(paymentID string, amount int, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s/capture", constants.PaymentURL, paymentID)
	// Amount should be in paisa
	if data == nil {
		data = make(map[string]interface{})
	}
	data["amount"] = amount

	return p.Request.Post(url, data, extraHeaders)
}

// Refund initiates a refund for the given paymentID.
func (p *Payment) Refund(paymentID string, amount int, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s/refund", constants.PaymentURL, paymentID)
	// Amount should be in paisa
	if data == nil {
		data = make(map[string]interface{})
	}
	data["amount"] = amount

	return p.Request.Post(url, data, extraHeaders)
}

// Transfer creates a transfer of the payment having the given paymentID.
func (p *Payment) Transfer(paymentID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/transfers", constants.PaymentURL, paymentID)
	return p.Request.Post(url, data, extraHeaders)
}

// Transfers fetches collection of all transfers associated with the given paymentID.
func (p *Payment) Transfers(paymentID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/transfers", constants.PaymentURL, paymentID)
	return p.Request.Get(url, queryParams, extraHeaders)
}

// BankTransfer fetches BankTransfer associated with the given paymentID.
func (p *Payment) BankTransfer(paymentID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/bank_transfer", constants.PaymentURL, paymentID)
	return p.Request.Get(url, queryParams, extraHeaders)
}
