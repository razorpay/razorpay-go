package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Card ...
type Card struct {
	Request *requests.Request
}

//Fetch ...
func (card *Card) Fetch(cardID string, data map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.CARD_URL, cardID)
	return card.Request.Get(url, data, options)
}
