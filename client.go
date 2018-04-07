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

	addon := resources.Addon{Request: Request}
	card := resources.Card{Request: Request}
	customer := resources.Customer{Request: Request}
	invoice := resources.Invoice{Request: Request}
	order := resources.Order{Request: Request}
	payment := resources.Payment{Request: Request}
	plan := resources.Plan{Request: Request}
	refund := resources.Refund{Request: Request}
	subscription := resources.Subscription{Request: Request}
	token := resources.Token{Request: Request}
	transfer := resources.Transfer{Request: Request}
	va := resources.VirtualAccount{Request: Request}
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

//AddHeaders ...
func (client *Client) AddHeaders(headers map[string]string) {
	Request.AddHeaders(headers)
}

//SetTimeout ...
func (client *Client) SetTimeout(timeout int16) {
	Request.SetTimeout(timeout)
}
