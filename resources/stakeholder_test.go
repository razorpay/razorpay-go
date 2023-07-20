package resources_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestStakeholderId = "sth_GLGgm8fFCKc92m"
const TestStakeholderAccountId = "acc_M28vQMUgbIBo1N"

func TestStakeholderFetch(t *testing.T) {
	url := "/v2" + constants.ACCOUNT_URL + "/" + TestStakeholderAccountId + constants.STAKEHOLDER_URL + "/" + TestStakeholderId
	teardown, fixture := utils.StartMockServer(url, "fake_stakeholder")
	defer teardown()
	body, err := utils.Client.Stakeholder.Fetch(TestStakeholderAccountId, TestStakeholderId, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestStakeholderCreate(t *testing.T) {
	url := "/v2" + constants.ACCOUNT_URL + "/" + TestStakeholderAccountId + constants.STAKEHOLDER_URL

	teardown, fixture := utils.StartMockServer(url, "fake_stakeholder")
	defer teardown()
	b, err := ioutil.ReadFile("../testdata/stakeholder_data.json")
	if err != nil {
		panic(err)
	}
	fixture = strings.Replace(strings.Replace(string(b), "\n", "", -1), " ", "", -1)

	var mapData map[string]interface{}
	if err := json.Unmarshal([]byte(fixture), &mapData); err != nil {
		fmt.Println(err)
	}

	body, err := utils.Client.Stakeholder.Create("acc_M28vQMUgbIBo1N", mapData, nil)
	assert.Equal(t, err, nil)
	assert.Equal(t, body["id"], TestStakeholderId)
}

func TestStakeholderEdit(t *testing.T) {
	url := "/v2" + constants.ACCOUNT_URL + "/" + TestStakeholderAccountId + constants.STAKEHOLDER_URL + "/" + TestStakeholderId
	teardown, fixture := utils.StartMockServer(url, "fake_stakeholder")
	defer teardown()
	params := map[string]interface{}{
		"phone": map[string]interface{}{
			"primary":   "9898989898",
			"secondary": "9898989898",
		},
	}
	body, err := utils.Client.Stakeholder.Edit(TestStakeholderAccountId, TestStakeholderId, params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestStakeholderAll(t *testing.T) {
	url := "/v2" + constants.ACCOUNT_URL + "/" + TestStakeholderAccountId + constants.STAKEHOLDER_URL
	teardown, fixture := utils.StartMockServer(url, "stakeholder_collection")
	defer teardown()
	body, err := utils.Client.Stakeholder.All(TestStakeholderAccountId, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestStakeholderFetchDoc(t *testing.T) {
	url := "/v2" + constants.ACCOUNT_URL + "/" + TestStakeholderAccountId + constants.STAKEHOLDER_URL + "/" + TestStakeholderId + "/documents"
	teardown, fixture := utils.StartMockServer(url, "fake_doc_data")
	defer teardown()
	body, err := utils.Client.Stakeholder.FetchStakeholderDoc(TestStakeholderAccountId, TestStakeholderId, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
