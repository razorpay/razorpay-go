package resources_test

import (
	"encoding/json"
	"testing"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestCustomerID = "fake_customer_id"

func TestCustomerFetch(t *testing.T) {
	url := constants.CUSTOMER_URL + "/" + TestCustomerID
	teardown, fixture := utils.StartMockServer(url, "fake_customer")
	defer teardown()
	body, err := utils.Client.Customer.Fetch(TestCustomerID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestCustomerCreate(t *testing.T) {
	teardown, fixture := utils.StartMockServer(constants.CUSTOMER_URL, "fake_customer")
	defer teardown()
	params := map[string]interface{}{
		"name":  "test",
		"email": "test@test.com",
	}
	body, err := utils.Client.Customer.Create(params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestCustomerEdit(t *testing.T) {
	url := constants.CUSTOMER_URL + "/" + TestCustomerID
	teardown, fixture := utils.StartMockServer(url, "fake_customer")
	defer teardown()
	params := map[string]interface{}{
		"email": "test@test.com",
	}
	body, err := utils.Client.Customer.Edit(TestCustomerID, params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
