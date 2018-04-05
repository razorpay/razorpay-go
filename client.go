package razorpay

import (
	"net/http"
	"time"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
	"github.com/razorpay/razorpay-go/resources"
)

var request *requests.Request

//Client ...
type Client struct {
	payment *resources.Payment
	order   *resources.Order
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
	request = &requests.Request{Auth: auth, HTTPClient: httpClient,
		Version: getVersion(), SDKName: getSDKName(),
		BaseURL: constants.BASE_URL}
	payment := resources.Payment{Request: request}
	order := resources.Order{Request: request}
	client := Client{&payment, &order}
	return &client
}

//AddHeaders ...
func (client *Client) AddHeaders(headers map[string]string) {
	request.AddHeaders(headers)
}

//SetTimeout ...
func (client *Client) SetTimeout(timeout int16) {
	request.SetTimeout(timeout)
}
