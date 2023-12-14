package resources_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestEntityCustomerID = "fake_customer_id"
const TestEntityAccountID = "acc_M28vQMUgbIBo1N"

func TestEntityCustomerFetch(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.CUSTOMER_URL + "/" + TestCustomerID
	entityUrl := constants.VERSION_V1 + constants.CUSTOMER_URL + "/" + TestCustomerID
	teardown, fixture := utils.StartMockServer(url, "fake_customer")
	defer teardown()
	body, err := utils.Client.Entity.Do(entityUrl, http.MethodGet, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestEntityFetchAccount(t *testing.T) {
	url := "/v2" + constants.ACCOUNT_URL + "/" + TestAccountID
	entityUrl := constants.VERSION_V2 + constants.ACCOUNT_URL + "/" + TestAccountID
	teardown, fixture := utils.StartMockServer(url, "fake_account")
	defer teardown()
	body, err := utils.Client.Entity.Do(entityUrl, http.MethodPatch, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestEntityCustomerCreate(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.CUSTOMER_URL
	entityUrl := constants.VERSION_V1 + constants.CUSTOMER_URL
	teardown, fixture := utils.StartMockServer(url, "fake_customer")
	defer teardown()
	params := map[string]interface{}{
		"name":  "test",
		"email": "test@test.com",
	}
	body, err := utils.Client.Entity.Do(entityUrl, http.MethodPost, params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestEntityAccountCreate(t *testing.T) {
	url := "/v2" + constants.ACCOUNT_URL
	entityUrl := constants.VERSION_V2 + constants.ACCOUNT_URL
	teardown, fixture := utils.StartMockServer(url, "fake_account")
	defer teardown()
	b, err := ioutil.ReadFile("../testdata/fake_account_request.json")
	if err != nil {
		panic(err)
	}
	fixture = strings.Replace(strings.Replace(string(b), "\n", "", -1), " ", "", -1)

	var mapData map[string]interface{}
	if err := json.Unmarshal([]byte(fixture), &mapData); err != nil {
		fmt.Println(err)
	}

	body, err := utils.Client.Entity.Do(entityUrl, http.MethodPost, mapData, nil)
	assert.Equal(t, err, nil)
	assert.Equal(t, body["email"], "gauriagain.kumar@example.org")
}

func TestEntityCustomerEdit(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.CUSTOMER_URL + "/" + TestCustomerID
	entityUrl := constants.VERSION_V1 + constants.CUSTOMER_URL + "/" + TestCustomerID
	teardown, fixture := utils.StartMockServer(url, "fake_customer")
	defer teardown()
	params := map[string]interface{}{
		"email": "test@test.com",
	}
	body, err := utils.Client.Entity.Do(entityUrl, http.MethodPut, params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
