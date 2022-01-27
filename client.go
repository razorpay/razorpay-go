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

//Client provides various helper methods to make HTTP requests to Razorpay's APIs.
type Client struct {
	Addon          *resources.Addon
	Card           *resources.Card
	Customer       *resources.Customer
	Invoice        *resources.Invoice
	PaymentLink    *resources.PaymentLink
	Order          *resources.Order
	Payment        *resources.Payment
	Plan           *resources.Plan
	Refund         *resources.Refund
	Subscription   *resources.Subscription
	Token          *resources.Token
	Transfer       *resources.Transfer
	VirtualAccount *resources.VirtualAccount
}

// NewClient creates and returns a new Razorpay client. key and secret
// are used to authenticate the requests made to Razorpay's APIs.
func NewClient(key string, secret string) *Client {
	auth := requests.Auth{Key: key, Secret: secret}
	httpClient := &http.Client{Timeout: requests.TIMEOUT * time.Second}
	Request = &requests.Request{Auth: auth, HTTPClient: httpClient,
		Version: getVersion(), SDKName: getSDKName(),
		BaseURL: constants.BaseURL}

	addon := resources.Addon{Request: Request}
	card := resources.Card{Request: Request}
	customer := resources.Customer{Request: Request}
	invoice := resources.Invoice{Request: Request}
	paymentLink := resources.PaymentLink{Request: Request}
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
		PaymentLink:    &paymentLink,
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

// AddHeaders adds additional headers to Razorpay's client. All requests
// made using the client will contain these additional headers in the HTTP request.
func (client *Client) AddHeaders(headers map[string]string) {
	Request.AddHeaders(headers)
}

// SetTimeout sets the timeout of Razorpay's Client. The default timeout will
// be overridden for all HTTP requests made using this client.
func (client *Client) SetTimeout(timeout int16) {
	Request.SetTimeout(timeout)
}

func getVersion() string {
	return SDKVersion
}

func getSDKName() string {
	return SDKName
}