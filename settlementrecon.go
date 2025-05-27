// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package razorpay

import (
	"context"
	"net/http"
	"net/url"

	"github.com/razorpay/razorpay-go/v2/internal/apijson"
	"github.com/razorpay/razorpay-go/v2/internal/apiquery"
	"github.com/razorpay/razorpay-go/v2/internal/requestconfig"
	"github.com/razorpay/razorpay-go/v2/option"
	"github.com/razorpay/razorpay-go/v2/packages/param"
	"github.com/razorpay/razorpay-go/v2/packages/respjson"
)

// SettlementReconService contains methods and other services that help with
// interacting with the gotue API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettlementReconService] method instead.
type SettlementReconService struct {
	Options []option.RequestOption
}

// NewSettlementReconService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettlementReconService(opts ...option.RequestOption) (r SettlementReconService) {
	r = SettlementReconService{}
	r.Options = opts
	return
}

// Use this endpoint to return a list of all transactions such as payments,
// refunds, transfers and adjustments settled to your account on a particular day
// or month. [See docs](https://razorpay.com/docs/api/settlements/fetch-recon)
func (r *SettlementReconService) Get(ctx context.Context, query SettlementReconGetParams, opts ...option.RequestOption) (res *SettlementReconGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "settlements/recon/combined"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SettlementReconGetResponse struct {
	// Number of settlement recon records returned.
	Count int64 `json:"count"`
	// Name of the entity. Always 'collection'.
	Entity string `json:"entity"`
	// List of settlement recon objects.
	Items []SettlementReconGetResponseItem `json:"items"`
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
func (r SettlementReconGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SettlementReconGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SettlementReconGetResponseItem struct {
	// The total amount of the transaction (in the smallest unit of currency).
	Amount int64 `json:"amount"`
	// The 4-character code denoting the issuing bank.
	CardIssuer string `json:"card_issuer,nullable"`
	// The card network used to process the payment.
	CardNetwork string `json:"card_network,nullable"`
	// The card type used to process the payment.
	CardType string `json:"card_type,nullable"`
	// Unix timestamp at which the transaction was created.
	CreatedAt int64 `json:"created_at"`
	// The amount, in currency subunits, that has been credited to your account for
	// this transaction.
	Credit int64 `json:"credit"`
	// The type of credit for the transaction.
	CreditType string `json:"credit_type"`
	// The currency of the transaction.
	Currency string `json:"currency"`
	// The amount, in currency subunits, that has been debited from your account for
	// this transaction.
	Debit int64 `json:"debit"`
	// Brief description about the transaction.
	Description string `json:"description,nullable"`
	// The unique identifier of any dispute, if any, for this transaction.
	DisputeID string `json:"dispute_id,nullable"`
	// The unique identifier of the transaction that has been settled.
	EntityID string `json:"entity_id"`
	// The fee charged for processing the transaction.
	Fee int64 `json:"fee"`
	// The payment method used to complete the payment.
	Method string `json:"method,nullable"`
	// Key-value pair that can be used to store additional information about the
	// entity.
	Notes NotesUnion `json:"notes"`
	// Indicates whether the account settlement for transfer is on hold.
	OnHold bool `json:"on_hold"`
	// Order id linked to the payment made by the customer that has been settled.
	OrderID string `json:"order_id,nullable"`
	// Receipt number entered while creating the Order.
	OrderReceipt string `json:"order_receipt,nullable"`
	// The unique identifier of the payment linked to refund or transfer that has been
	// settled.
	PaymentID string `json:"payment_id,nullable"`
	// Unix timestamp when the transaction was posted.
	PostedAt int64 `json:"posted_at,nullable"`
	// Indicates whether the transaction has been settled or not.
	Settled bool `json:"settled"`
	// Unix timestamp when the transaction was settled.
	SettledAt int64 `json:"settled_at"`
	// The unique identifier of the settlement transaction.
	SettlementID string `json:"settlement_id"`
	// The unique reference number linked to the settlement.
	SettlementUtr string `json:"settlement_utr,nullable"`
	// The tax charged for processing the transaction.
	Tax int64 `json:"tax"`
	// Indicates the type of transaction.
	//
	// Any of "payment", "refund", "transfer", "adjustment".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount        respjson.Field
		CardIssuer    respjson.Field
		CardNetwork   respjson.Field
		CardType      respjson.Field
		CreatedAt     respjson.Field
		Credit        respjson.Field
		CreditType    respjson.Field
		Currency      respjson.Field
		Debit         respjson.Field
		Description   respjson.Field
		DisputeID     respjson.Field
		EntityID      respjson.Field
		Fee           respjson.Field
		Method        respjson.Field
		Notes         respjson.Field
		OnHold        respjson.Field
		OrderID       respjson.Field
		OrderReceipt  respjson.Field
		PaymentID     respjson.Field
		PostedAt      respjson.Field
		Settled       respjson.Field
		SettledAt     respjson.Field
		SettlementID  respjson.Field
		SettlementUtr respjson.Field
		Tax           respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SettlementReconGetResponseItem) RawJSON() string { return r.JSON.raw }
func (r *SettlementReconGetResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SettlementReconGetParams struct {
	// The month the settlement was received in the MM format. For example, 06.
	Month int64 `query:"month,required" json:"-"`
	// The year the settlement was received in the YYYY format. For example, 2022.
	Year int64 `query:"year,required" json:"-"`
	// Specifies the number of available settlements to be fetched. Possible values- 1
	// to 1000.
	Count param.Opt[int64] `query:"count,omitzero" json:"-"`
	// The date on which the settlement was received in the DD format. For example, 11.
	Day param.Opt[int64] `query:"day,omitzero" json:"-"`
	// Specifies the number of available settlements to be skipped when fetching a
	// count.
	Skip param.Opt[int64] `query:"skip,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SettlementReconGetParams]'s query parameters as
// `url.Values`.
func (r SettlementReconGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
