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
	assert.Nil(t, err)
	assert.NotNil(t, body)

	jsonByteArray, _ := json.Marshal(body)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

// TestPayoutFetch is a function to test fetching a specific payout
func TestPayoutFetch(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.PAYOUT_URL + "/" + TestPayoutID
	teardown, fixture := utils.StartMockServer(url, "fake_payout")
	defer teardown()

	queryParams := map[string]interface{}{
		"expand[]": "fund_account",
	}
	body, err := utils.Client.Payout.Fetch(TestPayoutID, queryParams, nil)
	assert.Nil(t, err)
	assert.NotNil(t, body)

	jsonByteArray, _ := json.Marshal(body)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
