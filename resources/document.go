package resources

import (
	"fmt"
	"net/url"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Document ...
type Document struct {
	Request *requests.Request
}

// Create a Document
func (doc *Document) Create(params requests.FileUploadParams, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.DOCUMENT)
	return doc.Request.File(url, params, extraHeaders)
}

// Fetch document by given documentId.
func (doc *Document) Fetch(documentId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.DOCUMENT, url.PathEscape(documentId))
	return doc.Request.Get(url, queryParams, extraHeaders)
}
