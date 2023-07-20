package resources_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestInvoiceID = "fake_invoice_id"

const medium = "email"

func TestInvoiceAll(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.INVOICE_URL
	teardown, fixture := utils.StartMockServer(url, "invoice_collection")
	defer teardown()
	body, err := utils.Client.Invoice.All(nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestInvoiceAllWithOptions(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.INVOICE_URL
	teardown, fixture := utils.StartMockServer(url, "invoice_collection_with_one_invoice")
	defer teardown()
	data := map[string]interface{}{
		"count": 1,
	}
	body, err := utils.Client.Invoice.All(data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestInvoiceFetch(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.INVOICE_URL + "/" + TestInvoiceID
	teardown, fixture := utils.StartMockServer(url, "fake_invoice")
	defer teardown()
	body, err := utils.Client.Invoice.Fetch(TestInvoiceID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestInvoiceCreate(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.INVOICE_URL
	teardown, fixture := utils.StartMockServer(url, "fake_invoice")
	defer teardown()
	line_item := map[string]interface{}{
		"name":   "name",
		"amount": 100,
	}
	line_items := []map[string]interface{}{line_item}
	data := map[string]interface{}{
		"type":        "link",
		"decsription": "test",
		"line_items":  line_items,
	}
	body, err := utils.Client.Invoice.Create(data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestInvoiceCancel(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.INVOICE_URL + "/" + TestInvoiceID + "/cancel"
	teardown, fixture := utils.StartMockServer(url, "fake_invoice")
	defer teardown()
	body, err := utils.Client.Invoice.Cancel(TestInvoiceID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestInvoiceDelete(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.INVOICE_URL + "/" + TestInvoiceID
	teardown, fixture := utils.StartMockServer(url, "fake_invoice")
	defer teardown()
	body, err := utils.Client.Invoice.Delete(TestInvoiceID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestInvoiceUpdate(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.INVOICE_URL + "/" + TestInvoiceID
	teardown, fixture := utils.StartMockServer(url, "fake_invoice")
	defer teardown()
	data := map[string]interface{}{
		"notes": map[string]interface{}{
			"notes_key_1": "Tea, Earl Grey, Hot",
			"notes_key_2": "Tea, Earl Grey… decaf.",
		},
	}
	body, err := utils.Client.Invoice.Update(TestInvoiceID, data, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestInvoiceIssue(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.INVOICE_URL + "/" + TestInvoiceID + "/issue"
	teardown, fixture := utils.StartMockServer(url, "fake_invoice")
	defer teardown()
	body, err := utils.Client.Invoice.Issue(TestInvoiceID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestInvoiceNotify(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.INVOICE_URL + "/" + TestInvoiceID + "/notify_by/" + medium
	teardown, fixture := utils.StartMockServer(url, "success")
	defer teardown()
	body, err := utils.Client.Invoice.Notify(TestInvoiceID, medium, nil, nil)
	jsonByteArray, _ := json.Marshal(body)

	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestCreateRegistrationLink(t *testing.T) {
	url := fmt.Sprintf("/%s/subscription_registration/auth_links", constants.VERSION_V1)
	teardown, fixture := utils.StartMockServer(url, "fake_registration_link")
	defer teardown()
	params := map[string]interface{}{
		"customer": map[string]interface{}{
			"name":    "Gaurav Kumar",
			"email":   "gaurav.kumar@example.com",
			"contact": "9123456780",
		},
		"type":        "link",
		"amount":      0,
		"currency":    "INR",
		"description": "12 p.m. Meals",
		"subscription_registration": map[string]interface{}{
			"first_payment_amount": 100,
			"method":               "emandate",
			"auth_type":            "netbanking",
			"expire_at":            1580480689,
			"max_amount":           50000,
			"bank_account": map[string]interface{}{
				"beneficiary_name": "Gaurav Kumar",
				"account_number":   "11214311215411",
				"account_type":     "savings",
				"ifsc_code":        "HDFC0001233",
			},
		},
		"receipt":      "Receipt no. 1",
		"expire_by":    1880480689,
		"sms_notify":   1,
		"email_notify": 1,
		"notes": map[string]interface{}{
			"note_key 1": "Beam me up Scotty",
			"note_key 2": "Tea. Earl Gray. Hot.",
		},
	}
	body, err := utils.Client.Invoice.CreateRegistrationLink(params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
