package resources

import (
	"encoding/json"

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

// Payout represents the payout resource
type Payout struct {
	Request *requests.Request
}

// All fetches multiple payouts for the given query params
func (payout *Payout) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := "/" + constants.VERSION_V1 + constants.PAYOUT_URL
	return payout.Request.Get(url, queryParams, extraHeaders)
}

// Fetch fetches a payout having the given payoutID
func (payout *Payout) Fetch(id string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := "/" + constants.VERSION_V1 + constants.PAYOUT_URL + "/" + id
	return payout.Request.Get(url, queryParams, extraHeaders)
}

// Create creates a new payout for the given data
func (payout *Payout) Create(req *PayoutRequest, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := "/" + constants.VERSION_V1 + constants.PAYOUT_URL

	// Convert PayoutRequest to map[string]interface{}
	dataMap := make(map[string]interface{})
	if req != nil {
		jsonData, err := json.Marshal(req)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(jsonData, &dataMap); err != nil {
			return nil, err
		}
	}
	return payout.Request.Post(url, dataMap, extraHeaders)
}
