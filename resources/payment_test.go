package resources_test

import (
	"encoding/json"
	libhttp "net/http"
	"testing"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestPaymentID = "fake_payment_id"
const TestCaptureAmount = 500
const TestRefundAmount = 2000

func TestPaymentAll(t *testing.T) {
	teardown, fixture := utils.StartMockServer(constants.PAYMENT_URL, "payment_collection", libhttp.StatusOK)
	defer teardown()
	body, err := utils.Client.Payment.All(nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentAllWithOptions(t *testing.T) {
	teardown, fixture := utils.StartMockServer(constants.PAYMENT_URL, "payment_collection_with_one_payment", libhttp.StatusOK)
	defer teardown()
	queryParams := map[string]interface{}{
		"count": 1,
	}
	body, err := utils.Client.Payment.All(queryParams, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentFetch(t *testing.T) {
	url := constants.PAYMENT_URL + "/" + TestPaymentID
	teardown, fixture := utils.StartMockServer(url, "payment_collection_with_one_payment", libhttp.StatusOK)
	defer teardown()
	body, err := utils.Client.Payment.Fetch(TestPaymentID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentCapture(t *testing.T) {
	url := constants.PAYMENT_URL + "/" + TestPaymentID + "/capture"
	teardown, fixture := utils.StartMockServer(url, "fake_captured_payment", libhttp.StatusOK)
	defer teardown()
	body, err := utils.Client.Payment.Capture(TestPaymentID, TestCaptureAmount, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentRefundCreate(t *testing.T) {
	url := constants.PAYMENT_URL + "/" + TestPaymentID + "/refund"
	teardown, fixture := utils.StartMockServer(url, "fake_refund", libhttp.StatusOK)
	defer teardown()
	body, err := utils.Client.Payment.Refund(TestPaymentID, TestRefundAmount, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
func TestPaymentTransfer(t *testing.T) {
	url := constants.PAYMENT_URL + "/" + TestPaymentID + "/transfers"
	teardown, fixture := utils.StartMockServer(url, "transfers_collection_with_payment_id", libhttp.StatusOK)
	defer teardown()
	params := map[string]interface{}{
		"transfers": map[string]interface{}{
			"currency": map[string]interface{}{
				"amount":   100,
				"currency": "INR",
				"account":  "dummy_acc",
			},
		},
	}
	body, err := utils.Client.Payment.Transfer(TestPaymentID, params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentTransferFetch(t *testing.T) {
	url := constants.PAYMENT_URL + "/" + TestPaymentID + "/transfers"
	teardown, fixture := utils.StartMockServer(url, "transfers_collection_with_payment_id", libhttp.StatusOK)
	defer teardown()
	body, err := utils.Client.Payment.Transfers(TestPaymentID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestPaymentBankTransferFetch(t *testing.T) {
	url := constants.PAYMENT_URL + "/" + TestPaymentID + "/bank_transfer"
	teardown, fixture := utils.StartMockServer(url, "fake_bank_transfer", libhttp.StatusOK)
	defer teardown()
	body, err := utils.Client.Payment.BankTransfer(TestPaymentID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
