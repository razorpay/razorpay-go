package resources_test

import (
	"encoding/json"
	"testing"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestInvoiceID = "fake_invoice_id"

func TestInvoiceAll(t *testing.T) {
	url := constants.InvoiceURL
	teardown, fixture := utils.StartMockServer(url, "invoice_collection")
	defer teardown()
	body, err := utils.Client.Invoice.All(nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestInvoiceAllWithOptions(t *testing.T) {
	url := constants.InvoiceURL
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
	url := constants.InvoiceURL + "/" + TestInvoiceID
	teardown, fixture := utils.StartMockServer(url, "fake_invoice")
	defer teardown()
	body, err := utils.Client.Invoice.Fetch(TestInvoiceID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestInvoiceCreate(t *testing.T) {
	url := constants.InvoiceURL
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
