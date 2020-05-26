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

// Fetch fetches card having the given cardID.
func (card *Card) Fetch(cardID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.CARD_URL, cardID)
	return card.Request.Get(url, queryParams, extraHeaders)
}
