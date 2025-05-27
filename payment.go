// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package razorpay

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/razorpay/razorpay-go/v2/internal/apijson"
	"github.com/razorpay/razorpay-go/v2/internal/apiquery"
	"github.com/razorpay/razorpay-go/v2/internal/requestconfig"
	"github.com/razorpay/razorpay-go/v2/option"
	"github.com/razorpay/razorpay-go/v2/packages/param"
	"github.com/razorpay/razorpay-go/v2/packages/respjson"
)

// PaymentService contains methods and other services that help with interacting
// with the gotue API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentService] method instead.
type PaymentService struct {
	Options     []option.RequestOption
	CardDetails PaymentCardDetailService
	QrCodes     PaymentQrCodeService
	Refunds     PaymentRefundService
}

// NewPaymentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPaymentService(opts ...option.RequestOption) (r PaymentService) {
	r = PaymentService{}
	r.Options = opts
	r.CardDetails = NewPaymentCardDetailService(opts...)
	r.QrCodes = NewPaymentQrCodeService(opts...)
	r.Refunds = NewPaymentRefundService(opts...)
	return
}

// Use this endpoint to retrieve the details of a specific payment using its id.
// You can expand the card or EMI details by passing the query parameter
// `expand[]=card` or `expand[]=emi`.
// [See docs for card](https://razorpay.com/docs/api/payments/fetch-payment-expanded-card)
// [See docs for EMI](https://razorpay.com/docs/api/payments/fetch-payment-expanded-emi)
func (r *PaymentService) Get(ctx context.Context, id string, query PaymentGetParams, opts ...option.RequestOption) (res *Payment, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("payments/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Use this endpoint to update the `notes` field for a particular payment. Only the
// `notes` field can be updated. You can add up to 15 key-value pairs, each value
// not exceeding 256 characters.
// [See docs](https://razorpay.com/docs/api/payments/update)
func (r *PaymentService) Update(ctx context.Context, id string, body PaymentUpdateParams, opts ...option.RequestOption) (res *Payment, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("payments/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Use this endpoint to retrieve details of all the payments. By default, only the
// last 10 records are displayed. You can use the `count` and `skip` parameters to
// retrieve the specific number of records that you need. You can also expand card
// or EMI details using `expand[]=card` or `expand[]=emi`.
// [See docs](https://razorpay.com/docs/api/payments/fetch-all-payments)
func (r *PaymentService) List(ctx context.Context, query PaymentListParams, opts ...option.RequestOption) (res *PaymentListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "payments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Use this endpoint to change the payment status from authorized to captured.
// Attempting to capture a payment whose status is not authorized will produce an
// error. [See docs](https://razorpay.com/docs/api/payments/capture/)
func (r *PaymentService) Capture(ctx context.Context, id string, body PaymentCaptureParams, opts ...option.RequestOption) (res *Payment, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("payments/%s/capture", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A dynamic array consisting of unique reference numbers.
type AcquirerData struct {
	Arn string `json:"arn"`
	// Bank transaction ID
	BankTransactionID string `json:"bank_transaction_id"`
	Rrn               string `json:"rrn"`
	// transaction ID
	TransactionID string `json:"transaction_id"`
	// Upi transaction ID
	UpiTransactionID string `json:"upi_transaction_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arn               respjson.Field
		BankTransactionID respjson.Field
		Rrn               respjson.Field
		TransactionID     respjson.Field
		UpiTransactionID  respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AcquirerData) RawJSON() string { return r.JSON.raw }
func (r *AcquirerData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Card struct {
	// Unique identifier of the card.
	ID string `json:"id"`
	// Indicates if the card supports EMI.
	Emi bool `json:"emi"`
	// Name of the entity. Here, it is card.
	Entity string `json:"entity"`
	// Indicates if the card is international.
	International bool `json:"international"`
	// Card issuer code.
	Issuer string `json:"issuer"`
	// Last four digits of the card.
	Last4 string `json:"last4"`
	// Name on the card.
	Name string `json:"name"`
	// Card network.
	Network string `json:"network"`
	// Card sub type.
	SubType string `json:"sub_type"`
	// IIN of the tokenized card.
	TokenIin string `json:"token_iin,nullable"`
	// Card type.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Emi           respjson.Field
		Entity        respjson.Field
		International respjson.Field
		Issuer        respjson.Field
		Last4         respjson.Field
		Name          respjson.Field
		Network       respjson.Field
		SubType       respjson.Field
		TokenIin      respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Card) RawJSON() string { return r.JSON.raw }
func (r *Card) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Payment struct {
	// Unique identifier of the payment.
	ID string `json:"id"`
	// A dynamic array consisting of unique reference numbers.
	AcquirerData AcquirerData `json:"acquirer_data"`
	// The payment amount in currency subunits. For example, for an amount of ₹1.00
	// enter 100.
	Amount int64 `json:"amount"`
	// The amount refunded in currency subunits.
	AmountRefunded int64 `json:"amount_refunded"`
	// The 4-character bank code which the customer's account is associated with.
	Bank string `json:"bank,nullable"`
	// Indicates if the payment is captured.
	Captured bool `json:"captured"`
	// Details of the card used to make the payment. Present only if method is 'card'
	// and expand[]=card is used.
	Card Card `json:"card,nullable"`
	// The unique identifier of the card used by the customer to make the payment.
	CardID string `json:"card_id,nullable"`
	// Customer contact number used for the payment.
	Contact string `json:"contact"`
	// Timestamp, in UNIX format, on which the payment was created.
	CreatedAt int64 `json:"created_at"`
	// The currency in which the payment is made.
	Currency string `json:"currency"`
	// Description of the payment, if any.
	Description string `json:"description"`
	// Customer email address used for the payment.
	Email string `json:"email"`
	// Details of the EMI plan used to make the payment. Present only if method is
	// 'emi' and expand[]=emi is used.
	Emi PaymentEmi `json:"emi,nullable"`
	// Indicates the type of entity.
	Entity string `json:"entity"`
	// Error that occurred during payment.
	ErrorCode string `json:"error_code,nullable"`
	// Description of the error that occurred during payment.
	ErrorDescription string `json:"error_description,nullable"`
	// The exact error reason.
	ErrorReason string `json:"error_reason,nullable"`
	// The point of failure.
	ErrorSource string `json:"error_source,nullable"`
	// The stage where the transaction failure occurred.
	ErrorStep string `json:"error_step,nullable"`
	// Fee (including GST) charged by Razorpay.
	Fee int64 `json:"fee"`
	// Indicates whether the payment is done via an international card or a domestic
	// one.
	International bool `json:"international"`
	// Invoice id, if any.
	InvoiceID string `json:"invoice_id,nullable"`
	// The payment method used for making the payment.
	//
	// Any of "card", "netbanking", "wallet", "emi", "upi".
	Method PaymentMethod `json:"method"`
	// Key-value pair that can be used to store additional information about the
	// entity.
	Notes NotesUnion `json:"notes"`
	// Order id, if provided.
	OrderID string `json:"order_id"`
	// The refund status of the payment.
	RefundStatus  string `json:"refund_status,nullable"`
	Reward        string `json:"reward,nullable"`
	SourceChannel string `json:"source_channel,nullable"`
	// The status of the payment.
	//
	// Any of "created", "authorized", "captured", "refunded", "failed".
	Status PaymentStatus `json:"status"`
	// GST charged for the payment.
	Tax int64 `json:"tax"`
	// The customer's VPA (Virtual Payment Address) or UPI id used to make the payment.
	Vpa string `json:"vpa,nullable"`
	// The name of the wallet used by the customer to make the payment.
	Wallet string `json:"wallet,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AcquirerData     respjson.Field
		Amount           respjson.Field
		AmountRefunded   respjson.Field
		Bank             respjson.Field
		Captured         respjson.Field
		Card             respjson.Field
		CardID           respjson.Field
		Contact          respjson.Field
		CreatedAt        respjson.Field
		Currency         respjson.Field
		Description      respjson.Field
		Email            respjson.Field
		Emi              respjson.Field
		Entity           respjson.Field
		ErrorCode        respjson.Field
		ErrorDescription respjson.Field
		ErrorReason      respjson.Field
		ErrorSource      respjson.Field
		ErrorStep        respjson.Field
		Fee              respjson.Field
		International    respjson.Field
		InvoiceID        respjson.Field
		Method           respjson.Field
		Notes            respjson.Field
		OrderID          respjson.Field
		RefundStatus     respjson.Field
		Reward           respjson.Field
		SourceChannel    respjson.Field
		Status           respjson.Field
		Tax              respjson.Field
		Vpa              respjson.Field
		Wallet           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Payment) RawJSON() string { return r.JSON.raw }
func (r *Payment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details of the EMI plan used to make the payment. Present only if method is
// 'emi' and expand[]=emi is used.
type PaymentEmi struct {
	// The duration of the EMI plan in months.
	Duration int64 `json:"duration"`
	// The bank code of the EMI issuer.
	Issuer string `json:"issuer"`
	// The interest rate (in basis points) for the EMI plan.
	Rate int64  `json:"rate"`
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration    respjson.Field
		Issuer      respjson.Field
		Rate        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentEmi) RawJSON() string { return r.JSON.raw }
func (r *PaymentEmi) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payment method used for making the payment.
type PaymentMethod string

const (
	PaymentMethodCard       PaymentMethod = "card"
	PaymentMethodNetbanking PaymentMethod = "netbanking"
	PaymentMethodWallet     PaymentMethod = "wallet"
	PaymentMethodEmi        PaymentMethod = "emi"
	PaymentMethodUpi        PaymentMethod = "upi"
)

// The status of the payment.
type PaymentStatus string

const (
	PaymentStatusCreated    PaymentStatus = "created"
	PaymentStatusAuthorized PaymentStatus = "authorized"
	PaymentStatusCaptured   PaymentStatus = "captured"
	PaymentStatusRefunded   PaymentStatus = "refunded"
	PaymentStatusFailed     PaymentStatus = "failed"
)

type Refund struct {
	// The unique identifier of the refund.
	ID string `json:"id"`
	// A dynamic array consisting of unique reference numbers.
	AcquirerData AcquirerData `json:"acquirer_data"`
	// The amount to be refunded (in the smallest unit of currency).
	Amount int64 `json:"amount"`
	// This parameter is populated if the refund was created as part of a batch upload.
	BatchID string `json:"batch_id,nullable"`
	// Unix timestamp at which the refund was created.
	CreatedAt int64 `json:"created_at"`
	// The currency of payment amount for which the refund is initiated.
	Currency string `json:"currency"`
	// Indicates the type of entity. Here, it is refund.
	Entity string `json:"entity"`
	// Key-value pair that can be used to store additional information about the
	// entity.
	Notes NotesUnion `json:"notes"`
	// The unique identifier of the payment for which a refund is initiated.
	PaymentID string `json:"payment_id"`
	// A unique identifier provided by you for your internal reference.
	Receipt string `json:"receipt"`
	// The mode used to process a refund.
	//
	// Any of "instant", "normal".
	SpeedProcessed RefundSpeedProcessed `json:"speed_processed"`
	// The processing mode of the refund seen in the refund response.
	//
	// Any of "normal", "optimum".
	SpeedRequested RefundSpeedRequested `json:"speed_requested"`
	// Indicates the state of the refund.
	//
	// Any of "pending", "processed", "failed".
	Status RefundStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		AcquirerData   respjson.Field
		Amount         respjson.Field
		BatchID        respjson.Field
		CreatedAt      respjson.Field
		Currency       respjson.Field
		Entity         respjson.Field
		Notes          respjson.Field
		PaymentID      respjson.Field
		Receipt        respjson.Field
		SpeedProcessed respjson.Field
		SpeedRequested respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Refund) RawJSON() string { return r.JSON.raw }
func (r *Refund) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The mode used to process a refund.
type RefundSpeedProcessed string

const (
	RefundSpeedProcessedInstant RefundSpeedProcessed = "instant"
	RefundSpeedProcessedNormal  RefundSpeedProcessed = "normal"
)

// The processing mode of the refund seen in the refund response.
type RefundSpeedRequested string

const (
	RefundSpeedRequestedNormal  RefundSpeedRequested = "normal"
	RefundSpeedRequestedOptimum RefundSpeedRequested = "optimum"
)

// Indicates the state of the refund.
type RefundStatus string

const (
	RefundStatusPending   RefundStatus = "pending"
	RefundStatusProcessed RefundStatus = "processed"
	RefundStatusFailed    RefundStatus = "failed"
)

type PaymentListResponse struct {
	// Number of payments returned.
	Count int64 `json:"count"`
	// Name of the entity. Always 'collection'.
	Entity string `json:"entity"`
	// List of payment objects.
	Items []Payment `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Entity      respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentListResponse) RawJSON() string { return r.JSON.raw }
func (r *PaymentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentGetParams struct {
	// Use to expand the card or EMI details when the payment method is 'card' or
	// 'emi'.
	//
	// Any of "card", "emi".
	Expand PaymentGetParamsExpand `query:"expand[],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaymentGetParams]'s query parameters as `url.Values`.
func (r PaymentGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Use to expand the card or EMI details when the payment method is 'card' or
// 'emi'.
type PaymentGetParamsExpand string

const (
	PaymentGetParamsExpandCard PaymentGetParamsExpand = "card"
	PaymentGetParamsExpandEmi  PaymentGetParamsExpand = "emi"
)

type PaymentUpdateParams struct {
	// Contains user-defined fields, stored for reference purposes. Maximum 15
	// key-value pairs, 256 characters each.
	Notes PaymentUpdateParamsNotesUnion `json:"notes,omitzero,required"`
	paramObj
}

func (r PaymentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PaymentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PaymentUpdateParamsNotesUnion struct {
	OfStringMap   map[string]string `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u PaymentUpdateParamsNotesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[PaymentUpdateParamsNotesUnion](u.OfStringMap, u.OfStringArray)
}
func (u *PaymentUpdateParamsNotesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PaymentUpdateParamsNotesUnion) asAny() any {
	if !param.IsOmitted(u.OfStringMap) {
		return &u.OfStringMap
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

type PaymentListParams struct {
	// Number of payments to be fetched. Default is 10. Maximum is 100.
	Count param.Opt[int64] `query:"count,omitzero" json:"-"`
	// UNIX timestamp, in seconds, from when payments are to be fetched.
	From param.Opt[int64] `query:"from,omitzero" json:"-"`
	// Number of records to be skipped while fetching the payments.
	Skip param.Opt[int64] `query:"skip,omitzero" json:"-"`
	// UNIX timestamp, in seconds, till when payments are to be fetched.
	To param.Opt[int64] `query:"to,omitzero" json:"-"`
	// Use to expand the card or EMI details for each payment.
	//
	// Any of "card", "emi".
	Expand PaymentListParamsExpand `query:"expand[],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaymentListParams]'s query parameters as `url.Values`.
func (r PaymentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Use to expand the card or EMI details for each payment.
type PaymentListParamsExpand string

const (
	PaymentListParamsExpandCard PaymentListParamsExpand = "card"
	PaymentListParamsExpandEmi  PaymentListParamsExpand = "emi"
)

type PaymentCaptureParams struct {
	// The amount to be captured.
	Amount param.Opt[int64] `json:"amount,omitzero"`
	// The currency in which the payment is made.
	Currency param.Opt[string] `json:"currency,omitzero"`
	paramObj
}

func (r PaymentCaptureParams) MarshalJSON() (data []byte, err error) {
	type shadow PaymentCaptureParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentCaptureParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
