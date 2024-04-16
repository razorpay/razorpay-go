package resources_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestDisputeID = "fake_dispute_id"

func TestDisputeFetch(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.DISPUTE, TestDisputeID)
	teardown, fixture := utils.StartMockServer(url, "fake_dispute")
	defer teardown()
	body, err := utils.Client.Dispute.Fetch(TestDisputeID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestDisputeAll(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.DISPUTE)
	teardown, fixture := utils.StartMockServer(url, "dispute_collection")
	defer teardown()
	body, err := utils.Client.Dispute.All(nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestDisputeAccept(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s/accept", constants.VERSION_V1, constants.DISPUTE, TestDisputeID)
	teardown, fixture := utils.StartMockServer(url, "fake_dispute")
	defer teardown()
	body, err := utils.Client.Dispute.Accept(TestDisputeID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestDisputeContest(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s/contest", constants.VERSION_V1, constants.DISPUTE, TestDisputeID)
	teardown, fixture := utils.StartMockServer(url, "fake_dispute")
	defer teardown()

	shipping_proof := [2]string{"doc_EFtmUsbwpXwBH9", "doc_EFtmUsbwpXwBH8"}
	doc := [2]string{"doc_EFtmUsbwpXwBH1", "doc_EFtmUsbwpXwBH7"}

	other := make(map[string]interface{})
	other["0"] = map[string]interface{}{
		"type":         "receipt_signed_by_customer",
		"document_ids": doc,
	}

	params := map[string]interface{}{
		"amount":         5000,
		"summary":        "goods delivered",
		"shipping_proof": shipping_proof,
		"others":         other,
		"action":         "draft",
	}

	body, err := utils.Client.Dispute.Contest(TestDisputeID, params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
