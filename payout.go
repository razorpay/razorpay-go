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

// PayoutService contains methods and other services that help with interacting
// with the gotue API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPayoutService] method instead.
type PayoutService struct {
	Options []option.RequestOption
}

// NewPayoutService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPayoutService(opts ...option.RequestOption) (r PayoutService) {
	r = PayoutService{}
	r.Options = opts
	return
}

// Use this endpoint to retrieve the details of a particular payout in the system.
// [See docs](https://razorpay.com/docs/api/x/payouts/fetch-with-id)
func (r *PayoutService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Payout, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("payouts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Use this endpoint to retrieve the details of all the available payouts in the
// system. [See docs](https://razorpay.com/docs/api/x/payouts/fetch-all/)
func (r *PayoutService) List(ctx context.Context, query PayoutListParams, opts ...option.RequestOption) (res *PayoutListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "payouts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Payout struct {
	// The unique identifier of the payout.
	ID string `json:"id"`
	// The payout amount, in paise. The value does not include fees and tax. Fees and
	// tax, if any, are deducted from your account balance.
	Amount int64 `json:"amount"`
	// This value is returned if the contact was created as part of a bulk upload.
	BatchID string `json:"batch_id,nullable"`
	// Indicates the Unix timestamp when this payout was created.
	CreatedAt int64 `json:"created_at"`
	// The payout's currency. Here, it is INR.
	Currency string `json:"currency"`
	// The entity being created. Here, it will be payout.
	Entity string `json:"entity"`
	// Indicates the fee type charged for the payout. Possible value is free_payout.
	FeeType string `json:"fee_type"`
	// The fees for the payout. This value is returned only when the payout moves to
	// the processing state.
	Fees int64 `json:"fees"`
	// The unique identifier linked to the fund account.
	FundAccountID string `json:"fund_account_id"`
	// The mode used to make the payout.
	//
	// Any of "NEFT", "RTGS", "IMPS", "card", "UPI", "amazonpay".
	Mode PayoutMode `json:"mode"`
	// Custom note that also appears on the bank statement. Maximum length 30
	// characters.
	Narration string `json:"narration"`
	// Key-value pair that can be used to store additional information about the
	// entity.
	Notes NotesUnion `json:"notes"`
	// The purpose of the payout that is being created.
	Purpose string `json:"purpose"`
	// A user-generated reference given to the payout. Maximum length is 40 characters.
	ReferenceID string `json:"reference_id"`
	// The payout status.
	//
	// Any of "queued", "pending", "rejected", "processing", "processed", "cancelled",
	// "reversed", "failed".
	Status PayoutStatus `json:"status"`
	// This parameter returns the current status of the payout.
	StatusDetails PayoutStatusDetails `json:"status_details"`
	// The tax that is applicable for the fee being charged. This value is returned
	// only when the payout moves to the processing state.
	Tax int64 `json:"tax"`
	// The unique transaction number linked to a payout.
	Utr string `json:"utr,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Amount        respjson.Field
		BatchID       respjson.Field
		CreatedAt     respjson.Field
		Currency      respjson.Field
		Entity        respjson.Field
		FeeType       respjson.Field
		Fees          respjson.Field
		FundAccountID respjson.Field
		Mode          respjson.Field
		Narration     respjson.Field
		Notes         respjson.Field
		Purpose       respjson.Field
		ReferenceID   respjson.Field
		Status        respjson.Field
		StatusDetails respjson.Field
		Tax           respjson.Field
		Utr           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Payout) RawJSON() string { return r.JSON.raw }
func (r *Payout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The mode used to make the payout.
type PayoutMode string

const (
	PayoutModeNeft      PayoutMode = "NEFT"
	PayoutModeRtgs      PayoutMode = "RTGS"
	PayoutModeImps      PayoutMode = "IMPS"
	PayoutModeCard      PayoutMode = "card"
	PayoutModeUpi       PayoutMode = "UPI"
	PayoutModeAmazonpay PayoutMode = "amazonpay"
)

// The payout status.
type PayoutStatus string

const (
	PayoutStatusQueued     PayoutStatus = "queued"
	PayoutStatusPending    PayoutStatus = "pending"
	PayoutStatusRejected   PayoutStatus = "rejected"
	PayoutStatusProcessing PayoutStatus = "processing"
	PayoutStatusProcessed  PayoutStatus = "processed"
	PayoutStatusCancelled  PayoutStatus = "cancelled"
	PayoutStatusReversed   PayoutStatus = "reversed"
	PayoutStatusFailed     PayoutStatus = "failed"
)

// This parameter returns the current status of the payout.
type PayoutStatusDetails struct {
	// Status description.
	Description string `json:"description"`
	// Status reason.
	Reason string `json:"reason"`
	// Status source.
	Source string `json:"source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Reason      respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayoutStatusDetails) RawJSON() string { return r.JSON.raw }
func (r *PayoutStatusDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PayoutListResponse struct {
	// Number of payouts returned.
	Count int64 `json:"count"`
	// Name of the entity. Always collection.
	Entity string `json:"entity"`
	// List of payout objects.
	Items []Payout `json:"items"`
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
func (r PayoutListResponse) RawJSON() string { return r.JSON.raw }
func (r *PayoutListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PayoutListParams struct {
	// The account from which the payouts were done. For example, 7878780080316316.
	AccountNumber string `query:"account_number,required" json:"-"`
	// The unique identifier of the contact for which you want to fetch payouts.
	ContactID param.Opt[string] `query:"contact_id,omitzero" json:"-"`
	// Number of payouts to be fetched. Default value is 10. Maximum value is 100.
	Count param.Opt[int64] `query:"count,omitzero" json:"-"`
	// Timestamp, in Unix, from when you want to fetch payouts.
	From param.Opt[int64] `query:"from,omitzero" json:"-"`
	// The unique identifier of the fund account for which you want to fetch payouts.
	FundAccountID param.Opt[string] `query:"fund_account_id,omitzero" json:"-"`
	// The user-generated reference for which payouts are to be fetched. Maximum length
	// is 40 characters.
	ReferenceID param.Opt[string] `query:"reference_id,omitzero" json:"-"`
	// Numbers of payouts to be skipped. Default value is 0.
	Skip param.Opt[int64] `query:"skip,omitzero" json:"-"`
	// Timestamp, in Unix, till when you want to fetch payouts.
	To param.Opt[int64] `query:"to,omitzero" json:"-"`
	// The mode for which payouts are to be fetched. You can use one of the following
	// payout modes NEFT, RTGS, IMPS, UPI, card, amazonpay.
	//
	// Any of "NEFT", "RTGS", "IMPS", "UPI", "card", "amazonpay".
	Mode PayoutListParamsMode `query:"mode,omitzero" json:"-"`
	// The payout status. Possible payout states queued, pending, rejected, processing,
	// processed, cancelled, reversed, failed.
	//
	// Any of "queued", "pending", "rejected", "processing", "processed", "cancelled",
	// "reversed", "failed".
	Status PayoutListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PayoutListParams]'s query parameters as `url.Values`.
func (r PayoutListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The mode for which payouts are to be fetched. You can use one of the following
// payout modes NEFT, RTGS, IMPS, UPI, card, amazonpay.
type PayoutListParamsMode string

const (
	PayoutListParamsModeNeft      PayoutListParamsMode = "NEFT"
	PayoutListParamsModeRtgs      PayoutListParamsMode = "RTGS"
	PayoutListParamsModeImps      PayoutListParamsMode = "IMPS"
	PayoutListParamsModeUpi       PayoutListParamsMode = "UPI"
	PayoutListParamsModeCard      PayoutListParamsMode = "card"
	PayoutListParamsModeAmazonpay PayoutListParamsMode = "amazonpay"
)

// The payout status. Possible payout states queued, pending, rejected, processing,
// processed, cancelled, reversed, failed.
type PayoutListParamsStatus string

const (
	PayoutListParamsStatusQueued     PayoutListParamsStatus = "queued"
	PayoutListParamsStatusPending    PayoutListParamsStatus = "pending"
	PayoutListParamsStatusRejected   PayoutListParamsStatus = "rejected"
	PayoutListParamsStatusProcessing PayoutListParamsStatus = "processing"
	PayoutListParamsStatusProcessed  PayoutListParamsStatus = "processed"
	PayoutListParamsStatusCancelled  PayoutListParamsStatus = "cancelled"
	PayoutListParamsStatusReversed   PayoutListParamsStatus = "reversed"
	PayoutListParamsStatusFailed     PayoutListParamsStatus = "failed"
)
