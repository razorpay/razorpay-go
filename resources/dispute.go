package resources

import (
	"fmt"
	"net/url"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Dispute ...
type Dispute struct {
	Request *requests.Request
}

func (acc *Dispute) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.DISPUTE)
	return acc.Request.Get(url, queryParams, extraHeaders)
}

func (acc *Dispute) Fetch(disputeId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.DISPUTE, url.PathEscape(disputeId))
	return acc.Request.Get(url, queryParams, extraHeaders)
}

func (acc *Dispute) Accept(disputeId string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/accept", constants.VERSION_V1, constants.DISPUTE, url.PathEscape(disputeId))
	return acc.Request.Post(url, data, extraHeaders)
}

func (acc *Dispute) Contest(disputeId string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/contest", constants.VERSION_V1, constants.DISPUTE, url.PathEscape(disputeId))
	return acc.Request.Patch(url, data, extraHeaders)
}
