package resources_test

import (
	"encoding/json"
	"testing"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

func TestFundAccountAll(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.FUND_ACCOUNT_URL
	teardown, fixture := utils.StartMockServer(url, "fund_account_collection")
	defer teardown()
	params := map[string]interface{}{
		"customer_id": "cust_Aa000000000001",
	}
	body, err := utils.Client.FundAccount.All(params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestFundAccountCreate(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.FUND_ACCOUNT_URL
	teardown, fixture := utils.StartMockServer(url, "fake_fund_account")
	defer teardown()
	params := map[string]interface{}{
		"customer_id":  "cust_Aa000000000001",
		"account_type": "bank_account",
		"bank_account": map[string]interface{}{
			"name":           "Gaurav Kumar",
			"account_number": "11214311215411",
			"ifsc":           "HDFC0000053",
		},
	}

	body, err := utils.Client.FundAccount.Create(params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
