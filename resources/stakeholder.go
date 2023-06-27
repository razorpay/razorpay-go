package resources

import (
	"fmt"
	"net/url"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Stakeholder ...
type Stakeholder struct {
	Request *requests.Request
}

func (stake *Stakeholder) Create(accountId string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), constants.STAKEHOLDER_URL)
	return stake.Request.Post(url, data, extraHeaders)
}

func (stake *Stakeholder) Fetch(accountId string, stakeholderId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s%s/%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), constants.STAKEHOLDER_URL, url.PathEscape(stakeholderId))
	return stake.Request.Get(url, queryParams, extraHeaders)
}

func (stake *Stakeholder) All(accountId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), constants.STAKEHOLDER_URL)
	return stake.Request.Get(url, queryParams, extraHeaders)
}

func (stake *Stakeholder) Edit(accountId string, stakeholderId string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s%s/%s", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), constants.STAKEHOLDER_URL, url.PathEscape(stakeholderId))
	return stake.Request.Patch(url, data, extraHeaders)
}

func (stake *Stakeholder) UploadStakeholderDoc(accountId string, stakeholderId string, params requests.FileUploadParams) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/stakeholders/%s/documents", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), url.PathEscape(stakeholderId))
	return stake.Request.File(url, params)
}

func (stake *Stakeholder) FetchStakeholderDoc(accountId string, stakeholderId string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/stakeholders/%s/documents", constants.VERSION_V2, constants.ACCOUNT_URL, url.PathEscape(accountId), url.PathEscape(stakeholderId))
	return stake.Request.Get(url, queryParams, extraHeaders)
}
