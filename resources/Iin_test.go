package resources_test

import (
	"encoding/json"
	"testing"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestTokenIinID = "fake_token_id"

func TestFetchIin(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.IIN + "/" + TestTokenIinID
	teardown, fixture := utils.StartMockServer(url, "fake_Iin")
	defer teardown()
	body, err := utils.Client.Iin.Fetch(TestTokenIinID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
