package razorpay

import (
	"net/http"
	"time"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
	"github.com/razorpay/razorpay-go/resources"
)

// Client provides various helper methods to make HTTP requests to Razorpay's APIs.
type Client struct {
	*requests.Request
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
	Document       *resources.Document
	Dispute        *resources.Dispute
	Payout         *resources.Payout
}

// NewClient creates and returns a new Razorpay client. key and secret
// are used to authenticate the requests made to Razorpay's APIs.
func NewClient(key string, secret string) *Client {
	auth := requests.Auth{Key: key, Secret: secret}
	httpClient := &http.Client{Timeout: requests.TIMEOUT * time.Second}
	request := &requests.Request{
		Auth:       auth,
		AuthType:   requests.BasicAuth,
		HTTPClient: httpClient,
		Version:    getVersion(),
		SDKName:    getSDKName(),
		BaseURL:    constants.BASE_URL,
		Headers:    make(map[string]string),
	}

	account := resources.Account{Request: request}
	addon := resources.Addon{Request: request}
	card := resources.Card{Request: request}
	customer := resources.Customer{Request: request}
	invoice := resources.Invoice{Request: request}
	paymentLink := resources.PaymentLink{Request: request}
	order := resources.Order{Request: request}
	payment := resources.Payment{Request: request}
	product := resources.Product{Request: request}
	plan := resources.Plan{Request: request}
	refund := resources.Refund{Request: request}
	subscription := resources.Subscription{Request: request}
	token := resources.Token{Request: request}
	transfer := resources.Transfer{Request: request}
	va := resources.VirtualAccount{Request: request}
	qrCode := resources.QrCode{Request: request}
	fundAccount := resources.FundAccount{Request: request}
	settlement := resources.Settlement{Request: request}
	stakeholder := resources.Stakeholder{Request: request}
	item := resources.Item{Request: request}
	iin := resources.Iin{Request: request}
	webhook := resources.Webhook{Request: request}
	document := resources.Document{Request: request}
	dispute := resources.Dispute{Request: request}
	payout := resources.Payout{Request: request}

	client := Client{
		Request:        request,
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
		Document:       &document,
		Dispute:        &dispute,
		Payout:         &payout,
	}
	return &client
}


func NewClientOAuth(oauthToken string) *Client {
	auth := requests.Auth{Token: oauthToken}
	httpClient := &http.Client{Timeout: requests.TIMEOUT * time.Second}
	request := &requests.Request{
		Auth:       auth,
		AuthType:   requests.OAuthAuth,
		HTTPClient: httpClient,
		Version:    getVersion(),
		SDKName:    getSDKName(),
		BaseURL:    constants.BASE_URL,
		Headers:    make(map[string]string),
	}

	account := resources.Account{Request: request}
	addon := resources.Addon{Request: request}
	card := resources.Card{Request: request}
	customer := resources.Customer{Request: request}
	invoice := resources.Invoice{Request: request}
	paymentLink := resources.PaymentLink{Request: request}
	order := resources.Order{Request: request}
	payment := resources.Payment{Request: request}
	product := resources.Product{Request: request}
	plan := resources.Plan{Request: request}
	refund := resources.Refund{Request: request}
	subscription := resources.Subscription{Request: request}
	token := resources.Token{Request: request}
	transfer := resources.Transfer{Request: request}
	va := resources.VirtualAccount{Request: request}
	qrCode := resources.QrCode{Request: request}
	fundAccount := resources.FundAccount{Request: request}
	settlement := resources.Settlement{Request: request}
	stakeholder := resources.Stakeholder{Request: request}
	item := resources.Item{Request: request}
	iin := resources.Iin{Request: request}
	webhook := resources.Webhook{Request: request}
	document := resources.Document{Request: request}
	dispute := resources.Dispute{Request: request}
	payout := resources.Payout{Request: request}

	client := Client{
		Request:        request,
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
		Document:       &document,
		Dispute:        &dispute,
		Payout:         &payout,
	}
	return &client
}

// AddHeaders adds additional headers to Razorpay's client. All requests
// made using the client will contain these additional headers in the HTTP request.
func (client *Client) AddHeaders(headers map[string]string) {
	client.Request.AddHeaders(headers)
}

// SetTimeout sets the timeout of Razorpay's Client. The default timeout will
// be overridden for all HTTP requests made using this client.
func (client *Client) SetTimeout(timeout int16) {
	client.Request.SetTimeout(timeout)
}

func (client *Client) SetUserAgent(userAgent string) {
	client.Request.SetUserAgent(userAgent)
}

func getVersion() string {
	return SDKVersion
}

func getSDKName() string {
	return SDKName
}
