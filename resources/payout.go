package resources

import (
	"fmt"
	"net/http"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

// PayoutRequest represents the request structure for creating a payout to a bank account
type PayoutRequest struct {
	// AccountNumber is the account from which you want to make the payout
	AccountNumber string `json:"account_number"`

	// FundAccountID is the unique identifier linked to a fund account
	FundAccountID string `json:"fund_account_id"`

	// Amount is the payout amount in paise (e.g., 1000000 for ₹10,000)
	Amount int `json:"amount"`

	// Currency is the payout currency (default: INR)
	Currency string `json:"currency"`

	// Mode is the mode to be used for payout (NEFT, RTGS, IMPS, card)
	Mode string `json:"mode"`

	// Purpose is the purpose of the payout (refund, cashback, payout, salary, utility bill, vendor bill)
	Purpose string `json:"purpose"`

	// QueueIfLowBalance determines if payout should be queued when balance is low
	QueueIfLowBalance bool `json:"queue_if_low_balance"`

	// ReferenceID is a user-generated reference for the payout (max 40 chars)
	ReferenceID string `json:"reference_id"`

	// Narration is a custom note that appears on bank statement (max 30 chars)
	Narration string `json:"narration"`

	// Notes contains additional information about the payout (max 15 key-value pairs)
	Notes map[string]string `json:"notes,omitempty"`
}

// FundAccountResponse represents the fund account details in the payout response
type FundAccountResponse struct {
	ID          string              `json:"id"`
	Entity      string              `json:"entity"`
	ContactID   string              `json:"contact_id"`
	AccountType string              `json:"account_type"`
	BankAccount BankAccountResponse `json:"bank_account"`
	Active      bool                `json:"active"`
}

// BankAccountResponse represents the bank account details in the fund account response
type BankAccountResponse struct {
	Name          string `json:"name"`
	IFSC          string `json:"ifsc"`
	AccountNumber string `json:"account_number"`
}

// PayoutResponse represents the response structure returned by the Payout API
type PayoutResponse struct {
	// ID is the unique identifier of the payout
	ID string `json:"id"`

	// Entity represents the type of entity (payout)
	Entity string `json:"entity"`

	// FundAccountID is the unique identifier linked to the fund account
	FundAccountID string `json:"fund_account_id"`

	// Amount is the payout amount in paise
	Amount int `json:"amount"`

	// Currency is the payout currency
	Currency string `json:"currency"`

	// Notes contains additional information about the payout
	Notes map[string]string `json:"notes,omitempty"`

	// Fees is the fees charged for the payout
	Fees int `json:"fees"`

	// Tax is the tax charged on the fees
	Tax int `json:"tax"`

	// Status is the current status of the payout
	Status string `json:"status"`

	// UTR is the unique transaction reference number
	UTR string `json:"utr"`

	// Mode is the mode used for the payout
	Mode string `json:"mode"`

	// Purpose is the purpose of the payout
	Purpose string `json:"purpose"`

	// ReferenceID is the user-generated reference for the payout
	ReferenceID string `json:"reference_id"`

	// Narration is the custom note that appears on bank statement
	Narration string `json:"narration"`

	// BatchID is returned if contact was created as part of bulk upload
	BatchID string `json:"batch_id"`

	// StatusDetails contains details about the current status
	StatusDetails StatusDetails `json:"status_details"`

	// CreatedAt is the Unix timestamp when the payout was created
	CreatedAt int `json:"created_at"`

	// FeeType indicates the type of fee charged (e.g., free_payout)
	FeeType string `json:"fee_type"`
}

// StatusDetails represents the status details of a payout
type StatusDetails struct {
	// Description provides details about the current status
	Description string `json:"description"`

	// Source indicates the source of the status
	Source string `json:"source"`

	// Reason provides the reason for the current status
	Reason string `json:"reason"`
}

// Payout represents the payout resource
type Payout struct {
	Request *requests.Request
}

// All fetches multiple payouts for the given query params
func (payout *Payout) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.PAYOUT_URL)
	var result map[string]interface{}
	err := payout.Request.Call(http.MethodGet, url, nil, queryParams, &result)
	return result, err
}

// Fetch fetches a payout having the given payoutID
func (payout *Payout) Fetch(payoutID string, queryParams map[string]interface{}, extraHeaders map[string]string) (*PayoutResponse, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.PAYOUT_URL, payoutID)
	var result *PayoutResponse
	err := payout.Request.Call(http.MethodGet, url, nil, queryParams, &result)
	return result, err
}

// Create creates a new payout for the given data
func (payout *Payout) Create(data *PayoutRequest, extraHeaders map[string]string) (*PayoutResponse, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.PAYOUT_URL)
	var result *PayoutResponse
	err := payout.Request.Call(http.MethodPost, url, data, nil, &result)
	return result, err
}
