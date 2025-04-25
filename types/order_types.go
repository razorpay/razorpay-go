package types

type OrderRequest struct {
	Amount   int               `json:"amount"`
	Currency string            `json:"currency"`
	Receipt  string            `json:"receipt"`
	Notes    map[string]string `json:"notes,omitempty"`
}

type OrderResponse struct {
	ID         string            `json:"id"`
	Entity     string            `json:"entity"`
	Amount     int               `json:"amount"`
	AmountPaid int               `json:"amount_paid"`
	AmountDue  int               `json:"amount_due"`
	Currency   string            `json:"currency"`
	Receipt    string            `json:"receipt"`
	OfferID    *string           `json:"offer_id"`
	Status     string            `json:"status"`
	Attempts   int               `json:"attempts"`
	Notes      map[string]string `json:"notes,omitempty"`
	CreatedAt  int               `json:"created_at"`
}
