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
	Account        *resources.Account
	Card           *resources.Card
	Customer       *resources.Customer
	Invoice        *resources.Invoice
	PaymentLink    *resources.PaymentLink
	Order          *resources.Order
	Payment        *resources.Payment
	Plan           *resources.Plan
	Product        *resources.Product
	Refund         *resources.Refund
	Subscription   *resources.Subscription
	Token          *resources.Token
	Transfer       *resources.Transfer
	VirtualAccount *resources.VirtualAccount
	QrCode         *resources.QrCode
	FundAccount    *resources.FundAccount
	Settlement     *resources.Settlement
	Stakeholder    *resources.Stakeholder
	Item           *resources.Item
	Iin            *resources.Iin
	Webhook        *resources.Webhook
}

// NewClient creates and returns a new Razorpay client. key and secret
// are used to authenticate the requests made to Razorpay's APIs.
func NewClient(key string, secret string) *Client {
	auth := requests.Auth{Key: key, Secret: secret}
	httpClient := &http.Client{Timeout: requests.TIMEOUT * time.Second}
	Request = &requests.Request{Auth: auth, HTTPClient: httpClient,
		Version: getVersion(), SDKName: getSDKName(),
		BaseURL: constants.BASE_URL}

	account := resources.Account{Request: Request}
	addon := resources.Addon{Request: Request}
	card := resources.Card{Request: Request}
	customer := resources.Customer{Request: Request}
	invoice := resources.Invoice{Request: Request}
	paymentLink := resources.PaymentLink{Request: Request}
	order := resources.Order{Request: Request}
	payment := resources.Payment{Request: Request}
	product := resources.Product{Request: Request}
	plan := resources.Plan{Request: Request}
	refund := resources.Refund{Request: Request}
	subscription := resources.Subscription{Request: Request}
	token := resources.Token{Request: Request}
	transfer := resources.Transfer{Request: Request}
	va := resources.VirtualAccount{Request: Request}
	qrCode := resources.QrCode{Request: Request}
	fundAccount := resources.FundAccount{Request: Request}
	settlement := resources.Settlement{Request: Request}
	stakeholder := resources.Stakeholder{Request: Request}
	item := resources.Item{Request: Request}
	iin := resources.Iin{Request: Request}
	webhook := resources.Webhook{Request: Request}
	client := Client{
		Account:        &account,
		Addon:          &addon,
		Card:           &card,
		Customer:       &customer,
		Invoice:        &invoice,
		PaymentLink:    &paymentLink,
		Order:          &order,
		Payment:        &payment,
		Plan:           &plan,
		Product:        &product,
		Refund:         &refund,
		Subscription:   &subscription,
		Token:          &token,
		Transfer:       &transfer,
		VirtualAccount: &va,
		QrCode:         &qrCode,
		FundAccount:    &fundAccount,
		Settlement:     &settlement,
		Stakeholder:    &stakeholder,
		Item:           &item,
		Iin:            &iin,
		Webhook:        &webhook,
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
