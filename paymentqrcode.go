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

// PaymentQrCodeService contains methods and other services that help with
// interacting with the gotue API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentQrCodeService] method instead.
type PaymentQrCodeService struct {
	Options []option.RequestOption
}

// NewPaymentQrCodeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaymentQrCodeService(opts ...option.RequestOption) (r PaymentQrCodeService) {
	r = PaymentQrCodeService{}
	r.Options = opts
	return
}

// Use this endpoint to create a QR Code. You can share the short URL with
// customers to accept payments, print, and download it.
// [See docs](https://razorpay.com/docs/api/qr-codes/create/)
func (r *PaymentQrCodeService) New(ctx context.Context, body PaymentQrCodeNewParams, opts ...option.RequestOption) (res *QrCode, err error) {
	opts = append(r.Options[:], opts...)
	path := "payments/qr_codes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Use this endpoint to retrieve the details of a QR Code by using a Payment Id.
// [See docs](https://razorpay.com/docs/api/qr-codes/fetch-payment-id/)
func (r *PaymentQrCodeService) Get(ctx context.Context, id string, query PaymentQrCodeGetParams, opts ...option.RequestOption) (res *QrCode, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("payments/qr_codes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Use this endpoint to retrieve details of all QR Codes. By default, only the last
// 10 QR Codes are returned. You can use count and skip query parameters to change
// that behaviour. You can also filter QR Codes by customer_id.
// [See docs](https://razorpay.com/docs/api/qr-codes/fetch-customer-id)
func (r *PaymentQrCodeService) List(ctx context.Context, query PaymentQrCodeListParams, opts ...option.RequestOption) (res *PaymentQrCodeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "payments/qr_codes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Use this endpoint to close a QR Code.
// [See docs](https://razorpay.com/docs/api/qr-codes/close/)
func (r *PaymentQrCodeService) Close(ctx context.Context, id string, opts ...option.RequestOption) (res *QrCode, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("payments/qr_codes/%s/close", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type QrCode struct {
	// The unique identifier of the QR Code.
	ID string `json:"id"`
	// Unix timestamp at which the QR Code is scheduled to be automatically closed.
	CloseBy int64 `json:"close_by"`
	// The reason for the closure of the QR Code.
	//
	// Any of "on_demand", "paid".
	CloseReason QrCodeCloseReason `json:"close_reason,nullable"`
	// Unix timestamp at which the QR Code is automatically closed.
	ClosedAt int64 `json:"closed_at"`
	// Unix timestamp at which the QR Code is created.
	CreatedAt int64 `json:"created_at"`
	// The unique identifier of the customer the QR Code is linked with.
	CustomerID string `json:"customer_id"`
	// A brief description about the QR Code.
	Description string `json:"description"`
	// Indicates the type of entity. Here, it is qr_code.
	Entity string `json:"entity"`
	// Indicates if the QR Code should accept payments of specific amounts or any
	// amount.
	FixedAmount bool `json:"fixed_amount"`
	// The URL of the QR Code image.
	ImageURL string `json:"image_url"`
	// Label entered to identify the QR Code.
	Name string `json:"name"`
	// Key-value pair that can be used to store additional information about the
	// entity.
	Notes NotesUnion `json:"notes"`
	// The amount allowed for a transaction (in the smallest currency unit).
	PaymentAmount int64 `json:"payment_amount"`
	// The total amount received on the QR Code.
	PaymentsAmountReceived int64 `json:"payments_amount_received"`
	// The total number of payments received on the QR Code.
	PaymentsCountReceived int64 `json:"payments_count_received"`
	// Indicates the status of the QR Code.
	//
	// Any of "active", "closed".
	Status     QrCodeStatus       `json:"status"`
	TaxInvoice []QrCodeTaxInvoice `json:"tax_invoice"`
	// The type of the QR Code. Only 'upi_qr' is supported.
	//
	// Any of "upi_qr".
	Type QrCodeType `json:"type"`
	// Indicates if the QR Code should accept single or multiple payments.
	//
	// Any of "single_use", "multiple_use".
	Usage QrCodeUsage `json:"usage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		CloseBy                respjson.Field
		CloseReason            respjson.Field
		ClosedAt               respjson.Field
		CreatedAt              respjson.Field
		CustomerID             respjson.Field
		Description            respjson.Field
		Entity                 respjson.Field
		FixedAmount            respjson.Field
		ImageURL               respjson.Field
		Name                   respjson.Field
		Notes                  respjson.Field
		PaymentAmount          respjson.Field
		PaymentsAmountReceived respjson.Field
		PaymentsCountReceived  respjson.Field
		Status                 respjson.Field
		TaxInvoice             respjson.Field
		Type                   respjson.Field
		Usage                  respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QrCode) RawJSON() string { return r.JSON.raw }
func (r *QrCode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The reason for the closure of the QR Code.
type QrCodeCloseReason string

const (
	QrCodeCloseReasonOnDemand QrCodeCloseReason = "on_demand"
	QrCodeCloseReasonPaid     QrCodeCloseReason = "paid"
)

// Indicates the status of the QR Code.
type QrCodeStatus string

const (
	QrCodeStatusActive QrCodeStatus = "active"
	QrCodeStatusClosed QrCodeStatus = "closed"
)

// Tax invoice details for GST compliant transactions
type QrCodeTaxInvoice struct {
	// The GSTIN mentioned on the invoice. If not passed, it is picked up from the
	// database.
	BusinessGstin string `json:"business_gstin,required"`
	// CESS Amount on the invoice in paise. If not provided, the transaction will
	// default to the non-GST compliant UPI flow.
	CessAmount int64 `json:"cess_amount,required"`
	// Customer name on the invoice. If not provided, the transaction will default to
	// non-GST compliant UPI flow.
	CustomerName string `json:"customer_name,required"`
	// Unix Timestamp that indicates the issue date of the invoice. If not provided, it
	// will default to the current date.
	Date int64 `json:"date,required"`
	// GST amount on the invoice in paise. If not provided, the transaction will
	// default to the non-GST compliant UPI flow.
	GstAmount int64 `json:"gst_amount,required"`
	// This is the invoice number against which the payment is collected. If not
	// provided, the transaction will default to non-GST compliant UPI flow.
	Number string `json:"number,required"`
	// Indicates whether the transaction is interstate or intrastate
	//
	// Any of "interstate", "intrastate".
	SupplyType string `json:"supply_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BusinessGstin respjson.Field
		CessAmount    respjson.Field
		CustomerName  respjson.Field
		Date          respjson.Field
		GstAmount     respjson.Field
		Number        respjson.Field
		SupplyType    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QrCodeTaxInvoice) RawJSON() string { return r.JSON.raw }
func (r *QrCodeTaxInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the QR Code. Only 'upi_qr' is supported.
type QrCodeType string

const (
	QrCodeTypeUpiQr QrCodeType = "upi_qr"
)

// Indicates if the QR Code should accept single or multiple payments.
type QrCodeUsage string

const (
	QrCodeUsageSingleUse   QrCodeUsage = "single_use"
	QrCodeUsageMultipleUse QrCodeUsage = "multiple_use"
)

type PaymentQrCodeListResponse struct {
	// Number of QR Codes returned.
	Count int64 `json:"count"`
	// Name of the entity. Always 'collection'.
	Entity string `json:"entity"`
	// List of QR Code objects.
	Items []QrCode `json:"items"`
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
func (r PaymentQrCodeListResponse) RawJSON() string { return r.JSON.raw }
func (r *PaymentQrCodeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentQrCodeNewParams struct {
	// Indicates if the QR should accept payments of specific amounts or any amount.
	FixedAmount bool `json:"fixed_amount,required"`
	// The type of the QR Code. Only 'upi_qr' is supported.
	//
	// Any of "upi_qr".
	Type PaymentQrCodeNewParamsType `json:"type,omitzero,required"`
	// Indicates if the QR Code should accept single or multiple payments.
	//
	// Any of "single_use", "multiple_use".
	Usage PaymentQrCodeNewParamsUsage `json:"usage,omitzero,required"`
	// Unix timestamp at which the QR Code is scheduled to be automatically closed.
	CloseBy param.Opt[int64] `json:"close_by,omitzero"`
	// The unique identifier of the customer the QR Code is linked with.
	CustomerID param.Opt[string] `json:"customer_id,omitzero"`
	// A brief description about the QR Code.
	Description param.Opt[string] `json:"description,omitzero"`
	// Label entered to identify the QR Code.
	Name param.Opt[string] `json:"name,omitzero"`
	// The amount allowed for a transaction (in the smallest currency unit).
	PaymentAmount param.Opt[int64] `json:"payment_amount,omitzero"`
	// Key-value pair that can be used to store additional information about the
	// entity.
	Notes NotesUnionParam `json:"notes,omitzero"`
	paramObj
}

func (r PaymentQrCodeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PaymentQrCodeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentQrCodeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the QR Code. Only 'upi_qr' is supported.
type PaymentQrCodeNewParamsType string

const (
	PaymentQrCodeNewParamsTypeUpiQr PaymentQrCodeNewParamsType = "upi_qr"
)

// Indicates if the QR Code should accept single or multiple payments.
type PaymentQrCodeNewParamsUsage string

const (
	PaymentQrCodeNewParamsUsageSingleUse   PaymentQrCodeNewParamsUsage = "single_use"
	PaymentQrCodeNewParamsUsageMultipleUse PaymentQrCodeNewParamsUsage = "multiple_use"
)

type PaymentQrCodeGetParams struct {
	// The unique identifier of the payment whose QR Codes are to be fetched.
	PaymentID string `query:"payment_id,required" json:"-"`
	paramObj
}

// URLQuery serializes [PaymentQrCodeGetParams]'s query parameters as `url.Values`.
func (r PaymentQrCodeGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PaymentQrCodeListParams struct {
	// Number of QR Codes to be fetched. Default is 10. Maximum is 100.
	Count param.Opt[int64] `query:"count,omitzero" json:"-"`
	// The unique identifier of the customer whose QR Codes are to be fetched.
	CustomerID param.Opt[string] `query:"customer_id,omitzero" json:"-"`
	// UNIX timestamp, in seconds, from when QR Codes are to be fetched.
	From param.Opt[int64] `query:"from,omitzero" json:"-"`
	// The unique identifier of the payment whose QR Codes are to be fetched.
	PaymentID param.Opt[string] `query:"payment_id,omitzero" json:"-"`
	// Number of records to be skipped while fetching the QR Codes.
	Skip param.Opt[int64] `query:"skip,omitzero" json:"-"`
	// UNIX timestamp, in seconds, till when QR Codes are to be fetched.
	To param.Opt[int64] `query:"to,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaymentQrCodeListParams]'s query parameters as
// `url.Values`.
func (r PaymentQrCodeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
