package resources_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestVirtualID = "fake_virtual_id"
const PayerID = "fake_payer_id"

func TestVirtualAll(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.VIRTUAL_ACCOUNT_URL
	teardown, fixture := utils.StartMockServer(url, "fake_virtual_collection")
	defer teardown()
	body, err := utils.Client.VirtualAccount.All(nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestVirtualFetch(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.VIRTUAL_ACCOUNT_URL + "/" + TestVirtualID
	teardown, fixture := utils.StartMockServer(url, "fake_virtual")
	defer teardown()
	body, err := utils.Client.VirtualAccount.Fetch(TestVirtualID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestVirtualCreate(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.VIRTUAL_ACCOUNT_URL
	teardown, fixture := utils.StartMockServer(url, "fake_virtual")
	defer teardown()
	line_item := map[string]interface{}{
		"name":   "name",
		"amount": 1000,
	}
	lineItems := []map[string]interface{}{line_item}
	data := map[string]interface{}{
		"type":        "bank_account",
		"decsription": "test",
		"line_items":  lineItems,
	}
	body, err := utils.Client.VirtualAccount.Create(data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestVirtualClose(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s/close", constants.VERSION_V1, constants.VIRTUAL_ACCOUNT_URL, TestVirtualID)
	teardown, fixture := utils.StartMockServer(url, "virtual_collection")
	defer teardown()
	body, err := utils.Client.VirtualAccount.Close(TestVirtualID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestVirtualAddReceiver(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s/receivers", constants.VERSION_V1, constants.VIRTUAL_ACCOUNT_URL, TestVirtualID)
	teardown, fixture := utils.StartMockServer(url, "fake_receiver")
	defer teardown()
	line_item := make(map[string]interface{})
	line_item["0"] = "vpa"

	data := map[string]interface{}{
		"types": line_item,
		"vpa": map[string]interface{}{
			"descriptor": "gauravkumar",
		},
	}
	body, err := utils.Client.VirtualAccount.AddReceiver(TestVirtualID, data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestVirtualAllowedPayer(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s/allowed_payers", constants.VERSION_V1, constants.VIRTUAL_ACCOUNT_URL, TestVirtualID)
	teardown, fixture := utils.StartMockServer(url, "fake_allowed_payer")
	defer teardown()
	data := map[string]interface{}{
		"type": "bank_account",
		"bank_account": map[string]interface{}{
			"ifsc":           "UTIB0000013",
			"account_number": 914211112235679,
		},
	}
	body, err := utils.Client.VirtualAccount.AllowedPayer(TestVirtualID, data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestVirtualDeleteAllowedPayer(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s/allowed_payers/%s", constants.VERSION_V1, constants.VIRTUAL_ACCOUNT_URL, TestVirtualID, PayerID)
	teardown, fixture := utils.StartMockServer(url, "fake_allowed_payer")
	defer teardown()
	body, err := utils.Client.VirtualAccount.DeleteAllowedPayer(TestVirtualID, PayerID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
