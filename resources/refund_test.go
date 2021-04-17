package resources_test

import (
	"encoding/json"
	"testing"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestRefundID = "fake_refund_id"

func TestRefundAll(t *testing.T) {
	teardown, fixture := utils.StartMockServer(constants.RefundURL, "refund_collection")
	defer teardown()
	body, err := utils.Client.Refund.All(nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestRefundFetch(t *testing.T) {
	url := constants.RefundURL + "/" + TestRefundID
	teardown, fixture := utils.StartMockServer(url, "fake_refund")
	defer teardown()
	body, err := utils.Client.Refund.Fetch(TestRefundID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestRefundCreate(t *testing.T) {
	teardown, fixture := utils.StartMockServer(constants.RefundURL, "fake_refund")
	defer teardown()
	params := map[string]interface{}{
		"payment_id": TestPaymentID,
	}
	body, err := utils.Client.Refund.Create(params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
