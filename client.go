package razorpay

import (
	"net/http"
	"time"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
	"github.com/razorpay/razorpay-go/resources"
)

var request *requests.Request

//Client The base Razorpay SDK Client Struct. This is a composition of all the invidual client apis
type Client struct {
	Addon          *resources.Addon
	Card           *resources.Card
	Customer       *resources.Customer
	Invoice        *resources.Invoice
	Order          *resources.Order
	Payment        *resources.Payment
	Plan           *resources.Plan
	Refund         *resources.Refund
	Subscription   *resources.Subscription
	Token          *resources.Token
	Transfer       *resources.Transfer
	VirtualAccount *resources.VirtualAccount
}

func getVersion() string {
	return constants.SDKVersion
}

func getSDKName() string {
	return constants.SDKName
}

//NewClient Constructor to the Client Struct
func NewClient(key string, secret string) *Client {
	auth := requests.Auth{Key: key, Secret: secret}
	httpClient := &http.Client{Timeout: requests.TIMEOUT * time.Second}
	request = &requests.Request{Auth: auth, HTTPClient: httpClient,
		Version: getVersion(), SDKName: getSDKName(),
		BaseURL: constants.BASE_URL}

	addon := resources.Addon{Request: request}
	card := resources.Card{Request: request}
	customer := resources.Customer{Request: request}
	invoice := resources.Invoice{Request: request}
	order := resources.Order{Request: request}
	payment := resources.Payment{Request: request}
	plan := resources.Plan{Request: request}
	refund := resources.Refund{Request: request}
	subscription := resources.Subscription{Request: request}
	token := resources.Token{Request: request}
	transfer := resources.Transfer{Request: request}
	va := resources.VirtualAccount{Request: request}
	client := Client{
		Addon:          &addon,
		Card:           &card,
		Customer:       &customer,
		Invoice:        &invoice,
		Order:          &order,
		Payment:        &payment,
		Plan:           &plan,
		Refund:         &refund,
		Subscription:   &subscription,
		Token:          &token,
		Transfer:       &transfer,
		VirtualAccount: &va,
	}
	return &client
}

//AddHeaders method to add additional headers. Note that User-Agent and Content-Type cannot be overwritten at the moment
func (client *Client) AddHeaders(headers map[string]string) {
	request.AddHeaders(headers)
}

//SetTimeout method to set the client timeout
func (client *Client) SetTimeout(timeout int16) {
	request.SetTimeout(timeout)
}

//SetBaseURL method to set an additional BASE URL. This should be used for testing the golang SDK alone
func (client *Client) SetBaseURL(baseURL string) {
	request.BaseURL = baseURL
}
