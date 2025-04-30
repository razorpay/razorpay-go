package resources

// OrderRequest represents the request structure for create order
type OrderRequest struct {
	// Amount: The amount for which the order was created, in currency subunits.
	Amount int `json:"amount"`
	// Currency: The three-letter ISO currency code.
	Currency string `json:"currency"`
	// Receipt: A unique identifier provided by you for your internal reference.
	Receipt string `json:"receipt,omitempty"`
	// Notes: Key-value pair that can be used to store additional information about the order.
	Notes Notes `json:"notes,omitempty"`
	// PartialPayment: Indicates whether customers can make partial payments for the order.
	PartialPayment bool `json:"partial_payment,omitempty"`
	// FirstPaymentMinAmount: The minimum amount that must be paid as the first payment for a partial payment order.
	FirstPaymentMinAmount int `json:"first_payment_min_amount,omitempty"`
}

// OrderList represents the request structure for fetch all orders
type OrderList struct {
	// From: Starting timestamp of the time range for which orders are to be retrieved.
	From int `json:"from,omitempty"`
	// To: Ending timestamp of the time range for which orders are to be retrieved.
	To int `json:"to,omitempty"`
	// Count: Number of orders to be fetched.
	Count int `json:"count,omitempty"`
	// Skip: Number of records to be skipped for pagination.
	Skip int `json:"skip,omitempty"`
}

// OrderResponse represents the response structure for order response
type OrderResponse struct {
	// ID: The unique identifier of the order.
	ID string `json:"id"`
	// Entity: The type of resource being returned.
	Entity string `json:"entity"`
	// Amount: The total amount of the order, in currency subunits.
	Amount int `json:"amount"`
	// AmountPaid: The amount that has been paid for the order, in currency subunits.
	AmountPaid *int `json:"amount_paid"`
	// AmountDue: The amount that is remaining to be paid for the order, in currency subunits.
	AmountDue *int `json:"amount_due"`
	// Currency: The three-letter ISO currency code.
	Currency string `json:"currency"`
	// Receipt: The receipt number that you provided for this order.
	Receipt string `json:"receipt"`
	// OfferID: The unique identifier of the offer applied to the order.
	OfferID *string `json:"offer_id"`
	// Status: The current status of the order.
	Status string `json:"status"`
	// Attempts: The number of payment attempts made for the order.
	Attempts int `json:"attempts"`
	// Notes: Key-value pair containing additional information about the order.
	Notes interface{} `json:"notes"`
	// CreatedAt: Timestamp at which the order was created.
	CreatedAt int `json:"created_at"`
}

// OrderListResponse represents the response structure for fetch all orders
type OrderListResponse struct {
	// Entity: The type of resource being returned.
	Entity string `json:"entity"`
	// Count: The total number of orders returned.
	Count int `json:"count"`
	// Items: Array of order objects.
	Items []OrderResponse `json:"items"`
}

type OrderUpdateRequest struct {
	// Notes: Key-value pair that can be used to store additional information about the order.
	Notes Notes `json:"notes"`
}

// Notes represents a map of key-value pairs for storing additional information.
type Notes map[string]interface{}

// OrderPaymentsResponse represents the response structure for order payments
type OrderPaymentsResponse struct {
	Entity string              `json:"entity"`
	Count  int                 `json:"count"`
	Items  []OrderPaymentsItem `json:"items"`
}

// OrderPaymentsItem represents the item structure for order payments
type OrderPaymentsItem struct {
	// ID: The unique identifier of the payment.
	ID string `json:"id"`
	// Entity: The type of resource being returned.
	Entity string `json:"entity"`
	// Amount: The amount of the payment, in currency subunits.
	Amount int `json:"amount"`
	// Currency: The three-letter ISO currency code.
	Currency string `json:"currency"`
	// Status: The current status of the payment.
	Status string `json:"status"`
	// OrderID: The unique identifier of the order.
	OrderID string `json:"order_id"`
	// InvoiceID: The unique identifier of the invoice.
	InvoiceID *string `json:"invoice_id"`
	// International: Indicates whether the payment is international.
	International bool `json:"international"`
	// Method: The method of the payment.
	Method string `json:"method"`
	// AmountRefunded: The amount refunded for the payment, in currency subunits.
	AmountRefunded int `json:"amount_refunded"`
	// RefundStatus: The current status of the refund.
	RefundStatus *string `json:"refund_status"`
	// Captured: Indicates whether the payment is captured.
	Captured bool `json:"captured"`
	// Description: The description of the payment.
	Description *string `json:"description"`
	// CardID: The unique identifier of the card.
	CardID *string `json:"card_id"`
	// Bank: The bank of the payment.
	Bank *string `json:"bank"`
	// Wallet: The wallet of the payment.
	Wallet *string `json:"wallet"`
	// Vpa: The VPA of the payment.
	Vpa *string `json:"vpa"`
	// Email: The email of the payment.
	Email string `json:"email"`
	// Contact: The contact of the payment.
	Contact string `json:"contact"`
	// Notes: Key-value pair that can be used to store additional information about the order.
	Notes Notes `json:"notes"`
	// Fee: The fee of the payment.
	Fee *int `json:"fee"`
	// Tax: The tax of the payment.
	Tax *int `json:"tax"`
	// ErrorCode: The error code of the payment.
	ErrorCode *string `json:"error_code"`
	// ErrorDescription: The error description of the payment.
	ErrorDescription *string `json:"error_description"`
	// ErrorSource: The error source of the payment.
	ErrorSource *string `json:"error_source"`
	// ErrorStep: The error step of the payment.
	ErrorStep *string `json:"error_step"`
	// ErrorReason: The error reason of the payment.
	ErrorReason *string `json:"error_reason"`
	// AcquirerData: The acquirer data of the payment.
	AcquirerData interface{} `json:"acquirer_data"`
	// CreatedAt: The timestamp at which the payment was created.
	CreatedAt int `json:"created_at"`
	// Upi: The UPI of the payment.
	Upi struct {
		Vpa string `json:"vpa"`
	} `json:"upi"`
}
