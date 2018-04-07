package resources_test

import (
	"encoding/json"
	"testing"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestOrderID = "fake_order_id"

func TestOrderAll(t *testing.T) {
	teardown, fixture := utils.StartMockServer(constants.ORDER_URL, "order_collection")
	defer teardown()
	body, err := utils.Client.Order.All(nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestOrderWithOptions(t *testing.T) {
	teardown, fixture := utils.StartMockServer(constants.ORDER_URL, "order_collection_with_one_order")
	defer teardown()
	queryParams := map[string]interface{}{
		"count": 1,
	}
	body, err := utils.Client.Order.All(queryParams, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestOrderFetch(t *testing.T) {
	url := constants.ORDER_URL + "/" + TestOrderID
	teardown, fixture := utils.StartMockServer(url, "fake_order")
	defer teardown()
	body, err := utils.Client.Order.Fetch(TestOrderID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
func TestOrderPayments(t *testing.T) {
	url := constants.ORDER_URL + "/" + TestOrderID + "/payments"
	teardown, fixture := utils.StartMockServer(url, "fake_order")
	defer teardown()
	body, err := utils.Client.Order.Payments(TestOrderID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestOrderCreate(t *testing.T) {
	teardown, fixture := utils.StartMockServer(constants.ORDER_URL, "fake_order")
	defer teardown()
	params := map[string]interface{}{
		"amount":   100,
		"currency": "INR",
		"receipt":  "dummy",
	}
	body, err := utils.Client.Order.Create(params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestOrderEdit(t *testing.T) {
	url := constants.ORDER_URL + "/" + TestOrderID
	teardown, fixture := utils.StartMockServer(url, "fake_order")
	defer teardown()
	params := map[string]interface{}{
		"amount": 100,
	}
	body, err := utils.Client.Order.Edit(TestOrderID, params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
