package resources_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
	"github.com/razorpay/razorpay-go/utils"
	"github.com/stretchr/testify/assert"
)

const TestDocumentId = "fake_document_id"

func TestDocumentFetch(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.DOCUMENT + "/" + TestDocumentId
	teardown, fixture := utils.StartMockServer(url, "fake_document")
	defer teardown()
	body, err := utils.Client.Document.Fetch(TestDocumentId, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}

func TestDocumentCreate(t *testing.T) {
	url := "/" + constants.VERSION_V1 + constants.DOCUMENT + "/"
	teardown, fixture := utils.StartMockServer(url, "fake_document")
	defer teardown()

	file, err := os.Open("../testdata/fake_document.json")
	fields := map[string]string{
		"purpose": "dispute_evidence",
	}

	params := requests.FileUploadParams{
		File:   file,
		Fields: fields,
	}

	body, err := utils.Client.Document.Create(params, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
