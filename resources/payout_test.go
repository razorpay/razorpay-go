package resources_test

import (
	"encoding/json"
	"testing"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

// TestPayoutID is a constant representing a test payout ID
const TestPayoutID = "pout_00000000000001"

// TestPayoutAll is a function to test fetching all payouts
func TestPayoutAll(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PAYOUT_URL
	teardown, fixture := utils.StartMockServer(url, "payout_collection")
	defer teardown()
	queryParams := map[string]interface{}{
		"account_number": "7878780080316316",
	}
	body, err := utils.Client.Payout.All(queryParams, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

// TestPayoutFetch is a function to test fetching a specific payout
func TestPayoutFetch(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PAYOUT_URL + "/" + TestPayoutID
	teardown, fixture := utils.StartMockServer(url, "fake_payout")
	defer teardown()
	body, err := utils.Client.Payout.Fetch(TestPayoutID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

// TestPayoutCreate is a function to test creating a new payout
func TestPayoutCreate(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PAYOUT_URL
	teardown, fixture := utils.StartMockServer(url, "fake_payout")
	defer teardown()
	params := map[string]interface{}{
		"account_number":       "7878780080316316",
		"fund_account_id":      "fa_00000000000001",
		"amount":               1000000,
		"currency":             "INR",
		"mode":                 "IMPS",
		"purpose":              "refund",
		"queue_if_low_balance": true,
		"reference_id":         "Acme Transaction ID 12345",
		"narration":            "Acme Corp Fund Transfer",
		"notes": map[string]interface{}{
			"notes_key_1": "Tea, Earl Grey, Hot",
			"notes_key_2": "Tea, Earl Grey… decaf.",
		},
	}
	body, err := utils.Client.Payout.Create(params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
