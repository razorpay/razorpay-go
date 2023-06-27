package resources_test

import (
	"encoding/json"
	"testing"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestTransferID = "fake_transfer_id"

func TestTransferAll(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.TRANSFER_URL
	teardown, fixture := utils.StartMockServer(url, "transfer_collection")
	defer teardown()
	data := map[string]interface{}{
		"recipient_settlement_id": TestSettlementID,
	}
	body, err := utils.Client.Transfer.All(data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestTransferFetch(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.TRANSFER_URL + "/" + TestTransferID
	teardown, fixture := utils.StartMockServer(url, "fake_transfer")
	defer teardown()
	body, err := utils.Client.Transfer.Fetch(TestTransferID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestTransferCreate(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.TRANSFER_URL
	teardown, fixture := utils.StartMockServer(url, "fake_transfer")
	defer teardown()
	data := map[string]interface{}{
		"account":  "acc_HjVXbtpSCIxENR",
		"amount":   100,
		"currency": "INR",
	}
	body, err := utils.Client.Transfer.Create(data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestTransferEdit(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.TRANSFER_URL + "/" + TestTransferID
	teardown, fixture := utils.StartMockServer(url, "fake_transfer")
	defer teardown()
	data := map[string]interface{}{
		"on_hold":       1,
		"on_hold_until": 1679691505,
	}
	body, err := utils.Client.Transfer.Edit(TestTransferID, data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestTransferReverse(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.TRANSFER_URL + "/" + TestTransferID + "/reversals"
	teardown, fixture := utils.StartMockServer(url, "fake_transfer")
	defer teardown()
	data := map[string]interface{}{
		"amount": 100,
	}
	body, err := utils.Client.Transfer.Reverse(TestTransferID, data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestTransferReversals(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.TRANSFER_URL + "/" + TestTransferID + "/reversals"
	teardown, fixture := utils.StartMockServer(url, "reversals_collection")
	defer teardown()
	data := map[string]interface{}{
		"count": 1,
	}
	body, err := utils.Client.Transfer.Reversals(TestTransferID, data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
