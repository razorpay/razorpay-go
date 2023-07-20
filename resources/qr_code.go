package resources

import (
	"fmt"
	"net/url"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//QrCode ...
type QrCode struct {
	Request *requests.Request
}

// Create creates a new qrcode for the given data.
func (q *QrCode) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.QRCODE_URL)
	return q.Request.Post(url, data, extraHeaders)
}

//Fetch fetches the qrcode entity having the given QrCodeID.
func (q *QrCode) FetchPayments(QrCodeID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/payments", constants.VERSION_V1, constants.QRCODE_URL, url.PathEscape(QrCodeID))
	return q.Request.Get(url, queryParams, extraHeaders)
}

//Fetch fetches the QrCode entity having the given QrCodeID.
func (q *QrCode) Fetch(QrCodeID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.QRCODE_URL, url.PathEscape(QrCodeID))
	return q.Request.Get(url, queryParams, extraHeaders)
}

//All fetches collection of QrCode for the given queryParams.
func (q *QrCode) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.QRCODE_URL)
	return q.Request.Get(url, queryParams, extraHeaders)
}

//Close close a QrCode for the given QrCodeID.
func (q *QrCode) Close(QrCodeID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/close", constants.VERSION_V1, constants.QRCODE_URL, url.PathEscape(QrCodeID))
	return q.Request.Post(url, data, extraHeaders)
}
