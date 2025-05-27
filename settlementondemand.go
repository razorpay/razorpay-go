// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package razorpay

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/razorpay/razorpay-go/v2/internal/apijson"
	"github.com/razorpay/razorpay-go/v2/internal/requestconfig"
	"github.com/razorpay/razorpay-go/v2/option"
	"github.com/razorpay/razorpay-go/v2/packages/param"
	"github.com/razorpay/razorpay-go/v2/packages/respjson"
)

// SettlementOndemandService contains methods and other services that help with
// interacting with the gotue API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettlementOndemandService] method instead.
type SettlementOndemandService struct {
	Options []option.RequestOption
}

// NewSettlementOndemandService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettlementOndemandService(opts ...option.RequestOption) (r SettlementOndemandService) {
	r = SettlementOndemandService{}
	r.Options = opts
	return
}

// Use this endpoint to create an Instant Settlement.
// [See docs](https://razorpay.com/docs/api/settlements/instant/create/)
func (r *SettlementOndemandService) New(ctx context.Context, body SettlementOndemandNewParams, opts ...option.RequestOption) (res *SettlementOndemand, err error) {
	opts = append(r.Options[:], opts...)
	path := "settlements/ondemand"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Use this endpoint to retrieve the details of a particular Instant Settlement.
// [See docs](https://razorpay.com/docs/api/settlements/instant/fetch-with-id/)
func (r *SettlementOndemandService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *SettlementOndemand, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("settlements/ondemand/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SettlementOndemand struct {
	// The unique identifier of the instant settlement transaction.
	ID string `json:"id"`
	// Portion of the requested amount, in paise, yet to be settled to you.
	AmountPending int64 `json:"amount_pending"`
	// The settlement amount, in paise, requested by you.
	AmountRequested int64 `json:"amount_requested"`
	// Portion of the requested amount, in paise, that was not settled to you. This
	// amount is reversed to your PG current balance.
	AmountReversed int64 `json:"amount_reversed"`
	// Total amount (minus fees and tax), in paise, settled to the bank account.
	AmountSettled int64 `json:"amount_settled"`
	// Unix timestamp at which the instant settlement was created.
	CreatedAt int64 `json:"created_at"`
	// The 3-letter ISO currency code for the settlement.
	Currency string `json:"currency"`
	// Custom note for the instant settlement.
	Description string `json:"description"`
	// Indicates the type of entity. Here it is settlement.ondemand.
	Entity string `json:"entity"`
	// Total amount (fees+tax), in paise, deducted for the instant settlement.
	Fees int64 `json:"fees"`
	// Key-value pair that can be used to store additional information about the
	// entity.
	Notes NotesUnion `json:"notes"`
	// List of payouts created for the instant settlement.
	OndemandPayouts SettlementOndemandOndemandPayouts `json:"ondemand_payouts"`
	// Indicates whether full balance is settled.
	SettleFullBalance bool `json:"settle_full_balance"`
	// Indicates the state of the instant settlement.
	//
	// Any of "created", "initiated", "partially_processed", "processed", "reversed".
	Status SettlementOndemandStatus `json:"status"`
	// Total tax, in paise, charged for the fee component.
	Tax int64 `json:"tax"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AmountPending     respjson.Field
		AmountRequested   respjson.Field
		AmountReversed    respjson.Field
		AmountSettled     respjson.Field
		CreatedAt         respjson.Field
		Currency          respjson.Field
		Description       respjson.Field
		Entity            respjson.Field
		Fees              respjson.Field
		Notes             respjson.Field
		OndemandPayouts   respjson.Field
		SettleFullBalance respjson.Field
		Status            respjson.Field
		Tax               respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SettlementOndemand) RawJSON() string { return r.JSON.raw }
func (r *SettlementOndemand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List of payouts created for the instant settlement.
type SettlementOndemandOndemandPayouts struct {
	Count  int64                                   `json:"count"`
	Entity string                                  `json:"entity"`
	Items  []SettlementOndemandOndemandPayoutsItem `json:"items"`
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
func (r SettlementOndemandOndemandPayouts) RawJSON() string { return r.JSON.raw }
func (r *SettlementOndemandOndemandPayouts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SettlementOndemandOndemandPayoutsItem struct {
	ID            string `json:"id"`
	Amount        int64  `json:"amount"`
	AmountSettled int64  `json:"amount_settled,nullable"`
	CreatedAt     int64  `json:"created_at"`
	Entity        string `json:"entity"`
	Fees          int64  `json:"fees"`
	InitiatedAt   int64  `json:"initiated_at,nullable"`
	ProcessedAt   int64  `json:"processed_at,nullable"`
	ReversedAt    int64  `json:"reversed_at,nullable"`
	Status        string `json:"status"`
	Tax           int64  `json:"tax"`
	Utr           string `json:"utr,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Amount        respjson.Field
		AmountSettled respjson.Field
		CreatedAt     respjson.Field
		Entity        respjson.Field
		Fees          respjson.Field
		InitiatedAt   respjson.Field
		ProcessedAt   respjson.Field
		ReversedAt    respjson.Field
		Status        respjson.Field
		Tax           respjson.Field
		Utr           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SettlementOndemandOndemandPayoutsItem) RawJSON() string { return r.JSON.raw }
func (r *SettlementOndemandOndemandPayoutsItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the state of the instant settlement.
type SettlementOndemandStatus string

const (
	SettlementOndemandStatusCreated            SettlementOndemandStatus = "created"
	SettlementOndemandStatusInitiated          SettlementOndemandStatus = "initiated"
	SettlementOndemandStatusPartiallyProcessed SettlementOndemandStatus = "partially_processed"
	SettlementOndemandStatusProcessed          SettlementOndemandStatus = "processed"
	SettlementOndemandStatusReversed           SettlementOndemandStatus = "reversed"
)

type SettlementOndemandNewParams struct {
	// The amount, in paise, you want to get settled instantly.
	Amount int64 `json:"amount,required"`
	// Custom note for the instant settlement. Max 30 characters. Allowed a-z, A-Z,
	// 0-9, space.
	Description param.Opt[string] `json:"description,omitzero"`
	// Indicates whether full balance is settled. If true, Razorpay will settle the
	// maximum amount possible and ignore the amount parameter.
	SettleFullBalance param.Opt[bool] `json:"settle_full_balance,omitzero"`
	// Key-value pair that can be used to store additional information about the
	// entity.
	Notes NotesUnionParam `json:"notes,omitzero"`
	paramObj
}

func (r SettlementOndemandNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SettlementOndemandNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SettlementOndemandNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
