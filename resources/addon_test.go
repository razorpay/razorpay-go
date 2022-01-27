package resources_test

import (
	"encoding/json"
	"testing"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestAddonID = "ao_8sg8LU73Y3ieav"

func TestFetchAddon(t *testing.T) {
	url := constants.AddonURL + "/" + TestAddonID
	teardown, fixture := utils.StartMockServer(url, "fake_addon")
	defer teardown()
	body, err := utils.Client.Addon.Fetch(TestAddonID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestAddonDelete(t *testing.T) {
	url := constants.AddonURL + "/" + TestAddonID
	teardown, fixture := utils.StartMockServer(url, "fake_addon")
	defer teardown()
	body, err := utils.Client.Addon.Delete(TestAddonID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
