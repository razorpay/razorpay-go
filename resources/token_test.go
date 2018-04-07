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
	url := fmt.Sprintf("%s/%s/tokens", constants.CUSTOMER_URL, TestCustomerID)
	teardown, fixture := utils.StartMockServer(url, "token_collection")
	defer teardown()
	body, err := utils.Client.Token.All(TestCustomerID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestTokenFetch(t *testing.T) {
	url := fmt.Sprintf("%s/%s/tokens/%s", constants.CUSTOMER_URL, TestCustomerID, TestTokenID)
	teardown, fixture := utils.StartMockServer(url, "token_collection")
	defer teardown()
	body, err := utils.Client.Token.Fetch(TestCustomerID, TestTokenID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestTokenDelete(t *testing.T) {
	url := fmt.Sprintf("%s/%s/tokens/%s", constants.CUSTOMER_URL, TestCustomerID, TestTokenID)
	teardown, fixture := utils.StartMockServer(url, "token_delete")
	defer teardown()
	body, err := utils.Client.Token.Delete(TestCustomerID, TestTokenID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
