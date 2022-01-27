package resources_test

import (
	"encoding/json"
	"testing"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestPaymentLinkID = "payment_id"

func TestPaymentLinkAll(t *testing.T) {
	url := constants.PaymentLinkURL
	teardown, fixture := utils.StartMockServer(url, "paymentlink_collection")
	defer teardown()
	body, err := utils.Client.PaymentLink.All(nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentLinkFetch(t *testing.T) {
	url := constants.PaymentLinkURL + "/" + TestPaymentLinkID
	teardown, fixture := utils.StartMockServer(url, "paymentlink_collection")
	defer teardown()
	body, err := utils.Client.PaymentLink.Fetch(TestPaymentLinkID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentLinkCreate(t *testing.T) {
	url := constants.PaymentLinkURL
	teardown, fixture := utils.StartMockServer(url, "fake_paymentlink")
	defer teardown()
	line_item := map[string]interface{}{
		"name":   "name",
		"amount": 1000,
	}
	lineItems := []map[string]interface{}{line_item}
	data := map[string]interface{}{
		"type":        "link",
		"decsription": "test",
		"line_items":  lineItems,
	}
	body, err := utils.Client.PaymentLink.Create(data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

