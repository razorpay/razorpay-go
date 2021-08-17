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

func TestVirtualAll(t *testing.T) {
	url := constants.VirtualAccountURL
	teardown, fixture := utils.StartMockServer(url, "fake_virtual_collection")
	defer teardown()
	body, err := utils.Client.VirtualAccount.All(nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestVirtualFetch(t *testing.T) {
	url := constants.VirtualAccountURL + "/" + TestVirtualID
	teardown, fixture := utils.StartMockServer(url, "fake_virtual")
	defer teardown()
	body, err := utils.Client.VirtualAccount.Fetch(TestVirtualID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestVirtualCreate(t *testing.T) {
	url := constants.VirtualAccountURL
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
	url := fmt.Sprintf("%s/%s/close", constants.VirtualAccountURL, TestVirtualID)
	teardown, fixture := utils.StartMockServer(url, "virtual_collection")
	defer teardown()
	body, err := utils.Client.VirtualAccount.Close(TestVirtualID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

