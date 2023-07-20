package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Iin ...
type Iin struct {
	Request *requests.Request
}

func (i *Iin) Fetch(tokenIin string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.IIN, tokenIin)
	return i.Request.Get(url, queryParams, extraHeaders)
}
