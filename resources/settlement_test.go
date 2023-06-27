package resources_test

import (
	"encoding/json"
	"testing"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestSettlementID = "fake_settlement_id"

func TestSettlementAll(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.SETTLEMENT_URL
	teardown, fixture := utils.StartMockServer(url, "settlement_collection")
	defer teardown()
	body, err := utils.Client.Settlement.All(nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestSettlementFetch(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.SETTLEMENT_URL + "/" + TestSettlementID
	teardown, fixture := utils.StartMockServer(url, "fake_settlement")
	defer teardown()
	body, err := utils.Client.Settlement.Fetch(TestSettlementID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestSettlementReports(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.SETTLEMENT_URL + "/recon/combined"
	teardown, fixture := utils.StartMockServer(url, "settlement_report_collection")
	defer teardown()
	params := map[string]interface{}{
		"year":  2018,
		"month": 2,
	}
	body, err := utils.Client.Settlement.Reports(params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestFetchAllOnDemandSettlement(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.SETTLEMENT_URL + "/ondemand"
	teardown, fixture := utils.StartMockServer(url, "instant_settlement_collection")
	defer teardown()
	body, err := utils.Client.Settlement.FetchAllOnDemandSettlement(nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestFetchOnDemandSettlementById(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.SETTLEMENT_URL + "/ondemand/" + TestSettlementID
	teardown, fixture := utils.StartMockServer(url, "fake_instant_settlement")
	defer teardown()
	body, err := utils.Client.Settlement.FetchOnDemandSettlementById(TestSettlementID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestCreateOnDemandSettlement(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.SETTLEMENT_URL + "/ondemand/"
	teardown, fixture := utils.StartMockServer(url, "fake_instant_settlement")
	defer teardown()
	params := map[string]interface{}{
		"amount":              200000,
		"settle_full_balance": false,
		"description":         "Need this to make vendor payments.",
		"notes": map[string]interface{}{
			"notes_key_1": "Tea, Earl Grey, Hot",
			"notes_key_2": "Tea, Earl Grey… decaf.",
		},
	}

	body, err := utils.Client.Settlement.CreateOnDemandSettlement(params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
