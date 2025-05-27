// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package razorpay

import (
	"context"
	"encoding/json"
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

// OrderService contains methods and other services that help with interacting with
// the gotue API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrderService] method instead.
type OrderService struct {
	Options []option.RequestOption
}

// NewOrderService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOrderService(opts ...option.RequestOption) (r OrderService) {
	r = OrderService{}
	r.Options = opts
	return
}

// Create a new order.
func (r *OrderService) New(ctx context.Context, body OrderNewParams, opts ...option.RequestOption) (res *Order, err error) {
	opts = append(r.Options[:], opts...)
	path := "orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch an order by its unique ID.
func (r *OrderService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Order, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("orders/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an order's notes or other editable fields.
func (r *OrderService) Update(ctx context.Context, id string, body OrderUpdateParams, opts ...option.RequestOption) (res *Order, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("orders/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Use this endpoint to retrieve the details of all the orders you created.
func (r *OrderService) List(ctx context.Context, query OrderListParams, opts ...option.RequestOption) (res *OrderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Use this endpoint to retrieve all payments corresponding to a specific order.
// [See docs](https://razorpay.com/docs/api/payments/fetch-payments-orders)
func (r *OrderService) ListPayments(ctx context.Context, id string, query OrderListPaymentsParams, opts ...option.RequestOption) (res *OrderListPaymentsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("orders/%s/payments", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// NotesUnion contains all possible properties and values from [map[string]string],
// [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type NotesUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u NotesUnion) AsStringMap() (v map[string]string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NotesUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u NotesUnion) RawJSON() string { return u.JSON.raw }

func (r *NotesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this NotesUnion to a NotesUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// NotesUnionParam.Overrides()
func (r NotesUnion) ToParam() NotesUnionParam {
	return param.Override[NotesUnionParam](r.RawJSON())
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type NotesUnionParam struct {
	OfStringMap   map[string]string `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u NotesUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[NotesUnionParam](u.OfStringMap, u.OfStringArray)
}
func (u *NotesUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *NotesUnionParam) asAny() any {
	if !param.IsOmitted(u.OfStringMap) {
		return &u.OfStringMap
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

type Order struct {
	ID             string `json:"id"`
	Amount         int64  `json:"amount"`
	AmountDue      int64  `json:"amount_due"`
	AmountPaid     int64  `json:"amount_paid"`
	Attempts       int64  `json:"attempts"`
	CodFee         int64  `json:"cod_fee"`
	CreatedAt      int64  `json:"created_at"`
	Currency       string `json:"currency"`
	Entity         string `json:"entity"`
	LineItemsTotal int64  `json:"line_items_total"`
	// Key-value pair that can be used to store additional information about the
	// entity.
	Notes       NotesUnion `json:"notes"`
	OfferID     string     `json:"offer_id,nullable"`
	Receipt     string     `json:"receipt"`
	ShippingFee int64      `json:"shipping_fee"`
	Status      string     `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Amount         respjson.Field
		AmountDue      respjson.Field
		AmountPaid     respjson.Field
		Attempts       respjson.Field
		CodFee         respjson.Field
		CreatedAt      respjson.Field
		Currency       respjson.Field
		Entity         respjson.Field
		LineItemsTotal respjson.Field
		Notes          respjson.Field
		OfferID        respjson.Field
		Receipt        respjson.Field
		ShippingFee    respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Order) RawJSON() string { return r.JSON.raw }
func (r *Order) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderListResponse struct {
	Count  int64   `json:"count"`
	Entity string  `json:"entity"`
	Items  []Order `json:"items"`
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
func (r OrderListResponse) RawJSON() string { return r.JSON.raw }
func (r *OrderListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderListPaymentsResponse struct {
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
func (r OrderListPaymentsResponse) RawJSON() string { return r.JSON.raw }
func (r *OrderListPaymentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderNewParams struct {
	// Amount in the smallest currency unit (e.g., paise)
	Amount int64 `json:"amount,required"`
	// ISO currency code
	Currency string `json:"currency,required"`
	// Minimum amount for the first partial payment
	FirstPaymentMinAmount param.Opt[int64] `json:"first_payment_min_amount,omitzero"`
	// Unique identifier for the offer
	OfferID param.Opt[string] `json:"offer_id,omitzero"`
	// Allow partial payments
	PartialPayment param.Opt[bool] `json:"partial_payment,omitzero"`
	// Receipt number for internal reference
	Receipt param.Opt[string] `json:"receipt,omitzero"`
	// Key-value pair that can be used to store additional information about the
	// entity.
	Notes NotesUnionParam `json:"notes,omitzero"`
	paramObj
}

func (r OrderNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OrderNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderUpdateParams struct {
	// Key-value pair that can be used to store additional information about the
	// entity.
	Notes NotesUnionParam `json:"notes,omitzero"`
	paramObj
}

func (r OrderUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OrderUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderListParams struct {
	// The number of orders to be fetched. The default value is 10. The maximum value
	// is 100.
	Count param.Opt[int64] `query:"count,omitzero" json:"-"`
	// Timestamp (in Unix format) from when the orders should be fetched.
	From param.Opt[int64] `query:"from,omitzero" json:"-"`
	// Retrieves the orders that contain the provided value for receipt.
	Receipt param.Opt[string] `query:"receipt,omitzero" json:"-"`
	// The number of orders to be skipped. The default value is 0.
	Skip param.Opt[int64] `query:"skip,omitzero" json:"-"`
	// Timestamp (in Unix format) up till when orders are to be fetched.
	To param.Opt[int64] `query:"to,omitzero" json:"-"`
	// Retrieves Orders for which payments have been authorized. Payment and order
	// states differ.
	//
	// Any of true, false.
	Authorized bool `query:"authorized,omitzero" json:"-"`
	// Used to retrieve additional information about the payment.
	//
	// Any of "payments", "payments.card", "transfers", "virtual_account".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrderListParams]'s query parameters as `url.Values`.
func (r OrderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrderListPaymentsParams struct {
	// Use to expand the card or EMI details for each payment.
	//
	// Any of "card", "emi".
	Expand OrderListPaymentsParamsExpand `query:"expand[],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrderListPaymentsParams]'s query parameters as
// `url.Values`.
func (r OrderListPaymentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Use to expand the card or EMI details for each payment.
type OrderListPaymentsParamsExpand string

const (
	OrderListPaymentsParamsExpandCard OrderListPaymentsParamsExpand = "card"
	OrderListPaymentsParamsExpandEmi  OrderListPaymentsParamsExpand = "emi"
)
