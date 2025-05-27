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

// PaymentRefundService contains methods and other services that help with
// interacting with the gotue API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentRefundService] method instead.
type PaymentRefundService struct {
	Options []option.RequestOption
}

// NewPaymentRefundService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaymentRefundService(opts ...option.RequestOption) (r PaymentRefundService) {
	r = PaymentRefundService{}
	r.Options = opts
	return
}

// Use this endpoint to process refunds instantaneously to your customers. The
// instant refund is enabled by default for your account. You should set the refund
// speed to `optimum` when creating a refund request to ensure refunds are
// processed instantly.
// [See docs](https://razorpay.com/docs/api/refunds/create-instant/)
func (r *PaymentRefundService) New(ctx context.Context, id string, body PaymentRefundNewParams, opts ...option.RequestOption) (res *Refund, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("payments/%s/refund", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Use this endpoint to retrieve details of a specific refund made for a payment.
// [See docs](https://razorpay.com/docs/api/refunds/fetch-specific-refund-payment)
func (r *PaymentRefundService) Get(ctx context.Context, refundID string, query PaymentRefundGetParams, opts ...option.RequestOption) (res *Refund, err error) {
	opts = append(r.Options[:], opts...)
	if query.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if refundID == "" {
		err = errors.New("missing required refund_id parameter")
		return
	}
	path := fmt.Sprintf("payments/%s/refunds/%s", query.ID, refundID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Use this endpoint to retrieve multiple refunds for a payment. By default, only
// the last 10 refunds are returned. You can use count and skip query parameters to
// change that behaviour.
// [See docs](https://razorpay.com/docs/api/refunds/fetch-multiple-refund-payment)
func (r *PaymentRefundService) List(ctx context.Context, id string, query PaymentRefundListParams, opts ...option.RequestOption) (res *RefundList, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("payments/%s/refunds", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RefundList struct {
	// Number of refunds returned.
	Count int64 `json:"count"`
	// Name of the entity. Always 'collection'.
	Entity string `json:"entity"`
	// List of refund objects.
	Items []Refund `json:"items"`
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
func (r RefundList) RawJSON() string { return r.JSON.raw }
func (r *RefundList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentRefundNewParams struct {
	// The amount to be refunded. Amount should be in the smallest unit of the currency
	// in which the payment was made. Required in case of partial refund.
	Amount param.Opt[int64] `json:"amount,omitzero"`
	// A unique identifier provided by you for your internal reference.
	Receipt param.Opt[string] `json:"receipt,omitzero"`
	// Key-value pair that can be used to store additional information about the
	// entity.
	Notes NotesUnionParam `json:"notes,omitzero"`
	// Indicates that the refund will be processed at an optimal speed based on
	// Razorpay's internal fund transfer logic. Must be 'optimum'.
	//
	// Any of "optimum".
	Speed PaymentRefundNewParamsSpeed `json:"speed,omitzero"`
	paramObj
}

func (r PaymentRefundNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PaymentRefundNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentRefundNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates that the refund will be processed at an optimal speed based on
// Razorpay's internal fund transfer logic. Must be 'optimum'.
type PaymentRefundNewParamsSpeed string

const (
	PaymentRefundNewParamsSpeedOptimum PaymentRefundNewParamsSpeed = "optimum"
)

type PaymentRefundGetParams struct {
	ID string `path:"id,required" json:"-"`
	paramObj
}

type PaymentRefundListParams struct {
	// The number of refunds to fetch for the payment. Maximum is 100.
	Count param.Opt[int64] `query:"count,omitzero" json:"-"`
	// UNIX timestamp at which the refunds were created.
	From param.Opt[int64] `query:"from,omitzero" json:"-"`
	// The number of refunds to be skipped for the payment.
	Skip param.Opt[int64] `query:"skip,omitzero" json:"-"`
	// UNIX timestamp till which the refunds were created.
	To param.Opt[int64] `query:"to,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaymentRefundListParams]'s query parameters as
// `url.Values`.
func (r PaymentRefundListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
