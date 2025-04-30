package resources_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/razorpay/razorpay-go/v2/constants"
	"github.com/razorpay/razorpay-go/v2/resources"
	"github.com/razorpay/razorpay-go/v2/utils"
	"github.com/stretchr/testify/assert"
)

const TestOrderID = "fake_order_id"

func TestOrderAll(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.ORDER_URL)
	teardown, fixture := utils.StartMockServer(url, "order_collection_with_one_order")
	defer teardown()
	body, err := utils.Client.Order.All(nil, nil)
	jsonByteArray, _ := json.Marshal(body)

	var expectedResponse *resources.OrderListResponse
	err = json.Unmarshal([]byte(fixture), &expectedResponse)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestOrderWithOptions(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.ORDER_URL)
	teardown, fixture := utils.StartMockServer(url, "order_collection_with_one_order")
	defer teardown()
	queryParams := &resources.OrderList{
		Count: 1,
	}
	body, err := utils.Client.Order.All(queryParams, nil)
	jsonByteArray, _ := json.Marshal(body)

	var expectedResponse *resources.OrderListResponse
	err = json.Unmarshal([]byte(fixture), &expectedResponse)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestOrderFetch(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.ORDER_URL + "/" + TestOrderID
	teardown, fixture := utils.StartMockServer(url, "fake_order")
	defer teardown()
	body, err := utils.Client.Order.Fetch(TestOrderID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	var expectedResponse *resources.OrderResponse
	err = json.Unmarshal([]byte(fixture), &expectedResponse)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
func TestOrderPayments(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.ORDER_URL + "/" + TestOrderID + "/payments"
	teardown, fixture := utils.StartMockServer(url, "fake_order_payments")
	defer teardown()
	body, err := utils.Client.Order.Payments(TestOrderID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)

	var expectedResponse *resources.OrderPaymentsResponse
	err = json.Unmarshal([]byte(fixture), &expectedResponse)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, body.Count, expectedResponse.Count)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestOrderCreate(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.ORDER_URL)
	teardown, fixture := utils.StartMockServer(url, "fake_order")
	defer teardown()
	params := &resources.OrderRequest{
		Amount:   100,
		Currency: "INR",
		Receipt:  "dummy",
	}
	body, err := utils.Client.Order.Create(params, nil)
	jsonByteArray, _ := json.Marshal(body)

	var expectedResponse *resources.OrderResponse
	err = json.Unmarshal([]byte(fixture), &expectedResponse)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestOrderUpdate(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.ORDER_URL + "/" + TestOrderID
	teardown, fixture := utils.StartMockServer(url, "fake_order")
	defer teardown()
	params := &resources.OrderUpdateRequest{
		Notes: resources.Notes{
			"key1": "key1",
			"key2": "key",
		},
	}
	body, err := utils.Client.Order.Update(TestOrderID, params, nil)
	jsonByteArray, _ := json.Marshal(body)

	var expectedResponse *resources.OrderResponse
	err = json.Unmarshal([]byte(fixture), &expectedResponse)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, body.ID, expectedResponse.ID)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestViewRtoReview(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s/rto_review", constants.VERSION_V1, constants.ORDER_URL, TestOrderID)
	teardown, fixture := utils.StartMockServer(url, "fake_rto_review")
	defer teardown()
	body, err := utils.Client.Order.ViewRtoReview(TestOrderID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestEditFulfillment(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s/fulfillment", constants.VERSION_V1, constants.ORDER_URL, TestOrderID)
	teardown, fixture := utils.StartMockServer(url, "fake_fulfillment")
	defer teardown()
	params := map[string]interface{}{
		"payment_method": "upi",
		"shipping": map[string]interface{}{
			"waybill":  "123456789",
			"status":   "rto",
			"provider": "Bluedart",
		},
	}

	body, err := utils.Client.Order.EditFulfillment(TestOrderID, params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
