package resources_test

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestCustomerID = "fake_customer_id"
const TestEligibility = "elig_XXXXXXXXXXXXX1"

func TestCustomerFetch(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.CUSTOMER_URL + "/" + TestCustomerID
	teardown, fixture := utils.StartMockServer(url, "fake_customer")
	defer teardown()
	body, err := utils.Client.Customer.Fetch(TestCustomerID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestCustomerCreate(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.CUSTOMER_URL
	teardown, fixture := utils.StartMockServer(url, "fake_customer")
	defer teardown()
	params := map[string]interface{}{
		"name":  "test",
		"email": "test@test.com",
	}
	body, err := utils.Client.Customer.Create(params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestCustomerEdit(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.CUSTOMER_URL + "/" + TestCustomerID
	teardown, fixture := utils.StartMockServer(url, "fake_customer")
	defer teardown()
	params := map[string]interface{}{
		"email": "test@test.com",
	}
	body, err := utils.Client.Customer.Edit(TestCustomerID, params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestCustomerAll(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.CUSTOMER_URL
	teardown, fixture := utils.StartMockServer(url, "customers_collection")
	defer teardown()
	data := map[string]interface{}{
		"count": 1,
	}
	body, err := utils.Client.Customer.All(data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestAddBankAccount(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s/bank_account", constants.VERSION_V1, constants.CUSTOMER_URL, url.PathEscape(TestCustomerID))
	teardown, fixture := utils.StartMockServer(url, "bank_account")
	defer teardown()

	data := map[string]interface{}{
		"ifsc_code":            "UTIB0000194",
		"account_number":       "916010082985661",
		"beneficiary_name":     "Pratheek",
		"beneficiary_address1": "address 1",
		"beneficiary_address2": "address 2",
		"beneficiary_address3": "address 3",
		"beneficiary_address4": "address 4",
		"beneficiary_email":    "random@email.com",
		"beneficiary_mobile":   "8762489310",
		"beneficiary_city":     "Bangalore",
		"beneficiary_state":    "KA",
		"beneficiary_country":  "IN",
	}
	body, err := utils.Client.Customer.AddBankAccount(TestCustomerID, data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestDeleteBankAccount(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s/bank_account/%s", constants.VERSION_V1, constants.CUSTOMER_URL, url.PathEscape(TestCustomerID), url.PathEscape("ba_LSZht1Cm7xFTwF"))
	teardown, fixture := utils.StartMockServer(url, "bank_account")
	defer teardown()

	body, err := utils.Client.Customer.DeleteBankAccount(TestCustomerID, "ba_LSZht1Cm7xFTwF", nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestFetchEligibility(t *testing.T) {
	url := fmt.Sprintf("/%s%s/eligibility/%s", constants.VERSION_V1, constants.CUSTOMER_URL, url.PathEscape(TestEligibility))
	teardown, fixture := utils.StartMockServer(url, "fake_eligibility")
	defer teardown()

	body, err := utils.Client.Customer.FetchEligibility(TestEligibility, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestRequestEligibilityCheck(t *testing.T) {
	url := fmt.Sprintf("/%s%s/eligibility", constants.VERSION_V1, constants.CUSTOMER_URL)
	teardown, fixture := utils.StartMockServer(url, "eligibility_check")
	defer teardown()

	data := map[string]interface{}{
		"inquiry":  "affordability",
		"amount":   500000,
		"currency": "INR",
		"customer": map[string]interface{}{
			"id":         "cust_KhP5dO1dKmc0Rm",
			"contact":    "+918220276214",
			"ip":         "105.106.107.108",
			"referrer":   "https://merchansite.com/example/paybill",
			"user_agent": "Mozilla/5.0",
		},
	}

	body, err := utils.Client.Customer.RequestEligibilityCheck(data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
