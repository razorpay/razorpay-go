package resources_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestTokenID = "fake_token_id"

func TestTokenAll(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s/tokens", constants.VERSION_V1, constants.CUSTOMER_URL, TestCustomerID)
	teardown, fixture := utils.StartMockServer(url, "token_collection")
	defer teardown()
	body, err := utils.Client.Token.All(TestCustomerID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestTokenFetch(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s/tokens/%s", constants.VERSION_V1, constants.CUSTOMER_URL, TestCustomerID, TestTokenID)
	teardown, fixture := utils.StartMockServer(url, "token_collection")
	defer teardown()
	body, err := utils.Client.Token.Fetch(TestCustomerID, TestTokenID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestTokenDelete(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s/tokens/%s", constants.VERSION_V1, constants.CUSTOMER_URL, TestCustomerID, TestTokenID)
	teardown, fixture := utils.StartMockServer(url, "token_delete")
	defer teardown()
	body, err := utils.Client.Token.Delete(TestCustomerID, TestTokenID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestParterTokenCreate(t *testing.T) {
	url := fmt.Sprintf("/%s/tokens", constants.VERSION_V1)
	teardown, fixture := utils.StartMockServer(url, "fake_token")
	defer teardown()
	params := map[string]interface{}{
		"customer_id": "cust_1Aa00000000001",
		"method":      "card",
		"card": map[string]interface{}{
			"number":       "4111111111111111",
			"cvv":          "123",
			"expiry_month": "12",
			"expiry_year":  "21",
			"name":         "Gaurav Kumar",
		},
		"authentication": map[string]interface{}{
			"provider":                        "razorpay",
			"provider_reference_id":           "pay_123wkejnsakd",
			"authentication_reference_number": "100222021120200000000742753928",
		},
	}
	body, err := utils.Client.Token.Create(params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
