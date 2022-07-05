package resources

import (
	"fmt"
    "net/url"
	
	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Payment ...
type Payment struct {
	Request *requests.Request
}

// All fetches multiple payment entities for the given queryParams.
func (p *Payment) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return p.Request.Get(constants.PAYMENT_URL, queryParams, extraHeaders)
}

// Fetch fetches the payment entity for the given paymentID.
func (p *Payment) Fetch(paymentID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s", constants.PAYMENT_URL, url.PathEscape(paymentID))

	return p.Request.Get(url, queryParams, extraHeaders)
}

// Capture captures the payment having the given paymentID.
func (p *Payment) Capture(paymentID string, amount int, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s/capture", constants.PAYMENT_URL, url.PathEscape(paymentID))
	// Amount should be in paisa
	if data == nil {
		data = make(map[string]interface{})
	}
	data["amount"] = amount

	return p.Request.Post(url, data, extraHeaders)
}

// Refund initiates a refund for the given paymentID.
func (p *Payment) Refund(paymentID string, amount int, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s/refund", constants.PAYMENT_URL, url.PathEscape(paymentID))
	// Amount should be in paisa
	if data == nil {
		data = make(map[string]interface{})
	}
	data["amount"] = amount

	return p.Request.Post(url, data, extraHeaders)
}

// Transfer creates a transfer of the payment having the given paymentID.
func (p *Payment) Transfer(paymentID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/transfers", constants.PAYMENT_URL, url.PathEscape(paymentID))
	return p.Request.Post(url, data, extraHeaders)
}

// Transfers fetches collection of all transfers associated with the given paymentID.
func (p *Payment) Transfers(paymentID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/transfers", constants.PAYMENT_URL, url.PathEscape(paymentID))
	return p.Request.Get(url, queryParams, extraHeaders)
}

// BankTransfer fetches BankTransfer associated with the given paymentID.
func (p *Payment) BankTransfer(paymentID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/bank_transfer", constants.PAYMENT_URL, url.PathEscape(paymentID))
	return p.Request.Get(url, queryParams, extraHeaders)
}

// CreatePaymentJson creates a json payment for the given data.
func (p *Payment) CreatePaymentJson(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/create/json", constants.PAYMENT_URL)
	return p.Request.Post(url, data, extraHeaders)
}

// CreateRecurringPayment creates a recurring payment for the given data.
func (p *Payment) CreateRecurringPayment(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/create/recurring", constants.PAYMENT_URL)
	return p.Request.Post(url, data, extraHeaders)
}

// Edit updates the payment having the given paymentID.
func (p *Payment) Edit(paymentID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
    url := fmt.Sprintf("%s/%s", constants.PAYMENT_URL, url.PathEscape(paymentID))
    return p.Request.Patch(url, data, extraHeaders)
}

// FetchCardDetails fetches card details with the given paymentID.
func (p *Payment) FetchCardDetails(paymentID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/card", constants.PAYMENT_URL, url.PathEscape(paymentID))
	return p.Request.Get(url, queryParams, extraHeaders)
}

// FetchPaymentDowntime fetches downtime details.
func (p *Payment) FetchPaymentDowntime(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/downtimes", constants.PAYMENT_URL)
	return p.Request.Get(url, queryParams, extraHeaders)
}

// FetchPaymentDowntimeById fetches downtime details with the given downtimeID.
func (p *Payment) FetchPaymentDowntimeById(downtimeId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/downtimes/%s", constants.PAYMENT_URL, url.PathEscape(downtimeId))
	return p.Request.Get(url, queryParams, extraHeaders)
}

// FetchMultipleRefund fetches multiple refunds details with the given paymentID.
func (p *Payment) FetchMultipleRefund(paymentId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/refunds", constants.PAYMENT_URL, url.PathEscape(paymentId))
	return p.Request.Get(url, queryParams, extraHeaders)
}

// FetchRefund fetches refund detail with the given paymentId and refundId
func (p *Payment) FetchRefund(paymentId string,refundId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/refunds/%s", constants.PAYMENT_URL, url.PathEscape(paymentId), url.PathEscape(refundId))
	return p.Request.Get(url, queryParams, extraHeaders)
}

func (p *Payment) CreateUpi(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/create/upi", constants.PAYMENT_URL)
	return p.Request.Post(url, data, extraHeaders)
}

func (p *Payment) ValidateVpa(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/validate/vpa", constants.PAYMENT_URL)
	return p.Request.Post(url, data, extraHeaders)
}

func (p *Payment) FetchMethods(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s", constants.METHODS_URL)
	return p.Request.Get(url, data, extraHeaders)
}
