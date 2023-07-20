package resources_test

import (
	"encoding/json"
	"testing"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestItemID = "item_id"

func TestItemAll(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.ITEM_URL
	teardown, fixture := utils.StartMockServer(url, "item_collection")
	defer teardown()
	body, err := utils.Client.Item.All(nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestItemFetch(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.ITEM_URL + "/" + TestItemID
	teardown, fixture := utils.StartMockServer(url, "fake_item")
	defer teardown()
	body, err := utils.Client.Item.Fetch(TestItemID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestItemCreate(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.ITEM_URL
	teardown, fixture := utils.StartMockServer(url, "fake_item")
	defer teardown()
	params := map[string]interface{}{
		"name":        "Book / English August",
		"description": "An indian story, Booker prize winner.",
		"amount":      20000,
		"currency":    "INR",
	}
	body, err := utils.Client.Item.Create(params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestItemUpdate(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.ITEM_URL + "/" + TestItemID
	teardown, fixture := utils.StartMockServer(url, "fake_item")
	defer teardown()
	data := map[string]interface{}{
		"name":        "Book / Ignited Minds - Updated name!",
		"description": "New descirption too.",
	}
	body, err := utils.Client.Item.Update(TestItemID, data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
