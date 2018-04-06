package resources_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestPaymentID = "fake_payment_id"
const TestCaptureAmount = 500
const TestRefundAmount = 2000

func testResponse(json_body string, fixture string, t *testing.T) {
	status, err := utils.JSONCompare(fixture, string(json_body))
	if err == nil {
		assert.Equal(t, status, true)
	} else {
		panic(err.Error())
	}
}

func TestPaymentAll(t *testing.T) {
	teardown := utils.TestSetup()
	defer teardown()
	fixture := utils.GetFixture("payment_collection")
	utils.Mux.HandleFunc(constants.PAYMENT_URL, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, fixture)
	})
	body, err := utils.Client.Payment.All(nil, nil)
	json_body, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	testResponse(string(json_body), fixture, t)
}

func TestPaymentAllWithOptions(t *testing.T) {
	teardown := utils.TestSetup()
	defer teardown()
	fixture := utils.GetFixture("payment_collection_with_one_payment")
	utils.Mux.HandleFunc(constants.PAYMENT_URL, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, fixture)
	})
	queryParams := map[string]interface{}{
		"count": 1,
	}
	body, err := utils.Client.Payment.All(queryParams, nil)
	json_body, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	testResponse(string(json_body), fixture, t)
}

func TestPaymentFetch(t *testing.T) {
	teardown := utils.TestSetup()
	defer teardown()
	fixture := utils.GetFixture("fake_payment")
	url := constants.PAYMENT_URL + "/" + TestPaymentID
	utils.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, fixture)
	})
	body, err := utils.Client.Payment.Fetch(TestPaymentID, nil, nil)
	json_body, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	testResponse(string(json_body), fixture, t)
}

func TestPaymentCapture(t *testing.T) {
	teardown := utils.TestSetup()
	defer teardown()
	fixture := utils.GetFixture("fake_captured_payment")
	url := constants.PAYMENT_URL + "/" + TestPaymentID + "/capture"
	utils.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, fixture)
	})
	body, err := utils.Client.Payment.Capture(TestPaymentID, TestCaptureAmount, nil, nil)
	json_body, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	testResponse(string(json_body), fixture, t)
}

func TestPaymentRefundCreate(t *testing.T) {
	teardown := utils.TestSetup()
	defer teardown()
	fixture := utils.GetFixture("fake_refund")
	url := constants.PAYMENT_URL + "/" + TestPaymentID + "/refund"
	utils.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, fixture)
	})
	body, err := utils.Client.Payment.Refund(TestPaymentID, TestRefundAmount, nil, nil)
	json_body, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	testResponse(string(json_body), fixture, t)
}
func TestPaymentTransfer(t *testing.T) {
	teardown := utils.TestSetup()
	defer teardown()
	url := constants.PAYMENT_URL + "/" + TestPaymentID + "/transfers"
	fixture := utils.GetFixture("transfers_collection_with_payment_id")
	utils.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, fixture)
	})
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
	json_body, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	testResponse(string(json_body), fixture, t)
}

func TestPaymentTransferFetch(t *testing.T) {
	teardown := utils.TestSetup()
	defer teardown()
	url := constants.PAYMENT_URL + "/" + TestPaymentID + "/transfers"
	fixture := utils.GetFixture("transfers_collection_with_payment_id")
	utils.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, fixture)
	})
	body, err := utils.Client.Payment.Transfers(TestPaymentID, nil, nil)
	json_body, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	testResponse(string(json_body), fixture, t)

}

func TestPaymentBankTransferFetch(t *testing.T) {
	teardown := utils.TestSetup()
	defer teardown()
	url := constants.PAYMENT_URL + "/" + TestPaymentID + "/bank_transfer"
	fixture := utils.GetFixture("fake_bank_transfer")
	utils.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, fixture)
	})
	body, err := utils.Client.Payment.BankTransfer(TestPaymentID, nil, nil)
	json_body, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	testResponse(string(json_body), fixture, t)
}
