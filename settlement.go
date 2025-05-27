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

// SettlementService contains methods and other services that help with interacting
// with the gotue API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettlementService] method instead.
type SettlementService struct {
	Options  []option.RequestOption
	Recon    SettlementReconService
	Ondemand SettlementOndemandService
}

// NewSettlementService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettlementService(opts ...option.RequestOption) (r SettlementService) {
	r = SettlementService{}
	r.Options = opts
	r.Recon = NewSettlementReconService(opts...)
	r.Ondemand = NewSettlementOndemandService(opts...)
	return
}

// Use this endpoint to retrieve details of a settlement with its id.
// [See docs](https://razorpay.com/docs/api/settlements/fetch-with-id)
func (r *SettlementService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Settlement, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("settlements/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Use this endpoint to retrieve details of all settlements.
// [See docs](https://razorpay.com/docs/api/settlements/fetch-all/)
func (r *SettlementService) List(ctx context.Context, query SettlementListParams, opts ...option.RequestOption) (res *SettlementListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "settlements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Settlement struct {
	// The unique identifier of the settlement transaction.
	ID string `json:"id"`
	// The amount to be settled (in the smallest unit of currency).
	Amount int64 `json:"amount"`
	// Unix timestamp at which the settlement transaction was created.
	CreatedAt int64 `json:"created_at"`
	// Indicates the type of entity. Here, it is settlement.
	Entity string `json:"entity"`
	// The total fee charged for processing all payments received from customers
	// settled to you in this settlement transaction.
	Fees int64 `json:"fees"`
	// Indicates the settlement state.
	//
	// Any of "created", "processed", "failed".
	Status SettlementStatus `json:"status"`
	// The total tax, in currency subunits, charged on the fees collected to process
	// all payments received from customers settled to you in this settlement
	// transaction.
	Tax int64 `json:"tax"`
	// The Unique Transaction Reference (UTR) number available across banks.
	Utr string `json:"utr"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		CreatedAt   respjson.Field
		Entity      respjson.Field
		Fees        respjson.Field
		Status      respjson.Field
		Tax         respjson.Field
		Utr         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Settlement) RawJSON() string { return r.JSON.raw }
func (r *Settlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the settlement state.
type SettlementStatus string

const (
	SettlementStatusCreated   SettlementStatus = "created"
	SettlementStatusProcessed SettlementStatus = "processed"
	SettlementStatusFailed    SettlementStatus = "failed"
)

type SettlementListResponse struct {
	// Number of settlements returned.
	Count int64 `json:"count"`
	// Name of the entity. Always 'collection'.
	Entity string `json:"entity"`
	// List of settlement objects.
	Items []Settlement `json:"items"`
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
func (r SettlementListResponse) RawJSON() string { return r.JSON.raw }
func (r *SettlementListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SettlementListParams struct {
	// Number of settlement records to be fetched. Default is 10. Possible value- 1
	// to 100.
	Count param.Opt[int64] `query:"count,omitzero" json:"-"`
	// UNIX timestamp (in seconds) from when settlements are to be fetched.
	From param.Opt[int64] `query:"from,omitzero" json:"-"`
	// Number of settlement records to be skipped. Default is 0.
	Skip param.Opt[int64] `query:"skip,omitzero" json:"-"`
	// UNIX timestamp (in seconds) till when settlements are to be fetched.
	To param.Opt[int64] `query:"to,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SettlementListParams]'s query parameters as `url.Values`.
func (r SettlementListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
