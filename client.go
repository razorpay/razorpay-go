package razorpay

import (
	"net/http"
	"time"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
	"github.com/razorpay/razorpay-go/resources"
)

//Request ...
var Request *requests.Request

//Client ...
type Client struct {
	Payment  *resources.Payment
	Order    *resources.Order
	Customer *resources.Customer
	Refund   *resources.Refund
}

func getVersion() string {
	return SDKVersion
}

func getSDKName() string {
	return SDKName
}

//NewClient ...
func NewClient(key string, secret string) *Client {
	auth := requests.Auth{Key: key, Secret: secret}
	httpClient := &http.Client{Timeout: requests.TIMEOUT * time.Second}
	Request = &requests.Request{Auth: auth, HTTPClient: httpClient,
		Version: getVersion(), SDKName: getSDKName(),
		BaseURL: constants.BASE_URL}
	payment := resources.Payment{Request: Request}
	order := resources.Order{Request: Request}
	customer := resources.Customer{Request: Request}
	refund := resources.Refund{Request: Request}
	client := Client{
		Payment:  &payment,
		Order:    &order,
		Customer: &customer,
		Refund:   &refund,
	}
	return &client
}

//AddHeaders ...
func (client *Client) AddHeaders(headers map[string]string) {
	Request.AddHeaders(headers)
}

//SetTimeout ...
func (client *Client) SetTimeout(timeout int16) {
	Request.SetTimeout(timeout)
}
