package resources

import "time"

type OrderCreateResponse struct {
	ID         string      `json:"id"`
	Entity     string      `json:"entity"`
	Amount     uint        `json:"amount"`
	AmountPaid uint        `json:"amount_paid,omitempty"`
	AmountDue  uint        `json:"amount_due,omitempty"`
	Currency   string      `json:"currency"`
	Receipt    string      `json:"receipt"`
	OfferID    interface{} `json:"offer_id,omitempty"`
	Status     string      `json:"status"`
	Attempts   int         `json:"attempts"`
	Notes      interface{} `json:"notes"`
	CreatedAt  int         `json:"created_at"`
}

func (a *OrderCreateResponse) GetNotes() map[string]string {
	switch noteValues := a.Notes.(type) {
	case map[string]interface{}:
		notes := make(map[string]string)
		for key, value := range noteValues {
			notes[key] = value.(string)
		}
		return notes
	case []interface{}:
		return nil
	}
	return nil
}

func (a *OrderCreateResponse) CreatedTime() time.Time {
	return time.Unix(int64(a.CreatedAt), 0)
}

type OrderCreateRequest struct {
	Amount                uint              `json:"amount"`
	Currency              string            `json:"currency"`
	Receipt               string            `json:"receipt"`
	PartialPayment        bool              `json:"partial_payment"`
	Notes                 map[string]string `json:"notes,omitempty"`
	FirstPaymentMinAmount uint              `json:"first_payment_min_amount,omitempty"`
}
