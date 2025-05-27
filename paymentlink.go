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

// PaymentLinkService contains methods and other services that help with
// interacting with the gotue API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentLinkService] method instead.
type PaymentLinkService struct {
	Options []option.RequestOption
}

// NewPaymentLinkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaymentLinkService(opts ...option.RequestOption) (r PaymentLinkService) {
	r = PaymentLinkService{}
	r.Options = opts
	return
}

// Use this endpoint to create a Payment Link using basic details such as amount,
// expiry date, reference id, description, customer details and so on.
// [See docs](https://razorpay.com/docs/api/payments/payment-links/create-standard/)
func (r *PaymentLinkService) New(ctx context.Context, body PaymentLinkNewParams, opts ...option.RequestOption) (res *PaymentLink, err error) {
	opts = append(r.Options[:], opts...)
	path := "payment_links"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Use this endpoint to retrieve the details of a Payment Link using its id.
// [See docs](https://razorpay.com/docs/api/payments/payment-links/fetch-id-standard/)
func (r *PaymentLinkService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PaymentLink, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("payment_links/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Use this endpoint to edit the Standard Payment Link details such as the
// reference id, expiry date, enabling reminders and so on.
// [See docs](https://razorpay.com/docs/api/payments/payment-links/update-standard)
func (r *PaymentLinkService) Update(ctx context.Context, id string, body PaymentLinkUpdateParams, opts ...option.RequestOption) (res *PaymentLink, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("payment_links/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Use this endpoint to retrieve the details of all Payment Links. To get only UPI
// Payment Links, filter the response objects where `upi_link` is `true`.
// [See docs](https://razorpay.com/docs/api/payments/payment-links/fetch-all-upi)
func (r *PaymentLinkService) List(ctx context.Context, query PaymentLinkListParams, opts ...option.RequestOption) (res *PaymentLinkListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "payment_links"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Use this endpoint to send or resend notifications to your customers via email or
// SMS. [See docs](https://razorpay.com/docs/api/payments/payment-links/resend/)
func (r *PaymentLinkService) Notify(ctx context.Context, medium PaymentLinkNotifyParamsMedium, body PaymentLinkNotifyParams, opts ...option.RequestOption) (res *PaymentLinkNotifyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("payment_links/%s/notify_by/%v", body.ID, medium)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type PaymentLink struct {
	// Unique identifier of the Payment Link.
	ID string `json:"id"`
	// Indicates whether customers can make partial payments.
	AcceptPartial bool `json:"accept_partial"`
	// Amount to be paid using the Payment Link.
	Amount int64 `json:"amount"`
	// Amount paid by the customer.
	AmountPaid int64 `json:"amount_paid"`
	// Callback method.
	CallbackMethod string `json:"callback_method"`
	// Redirect URL after payment.
	CallbackURL string `json:"callback_url"`
	// Timestamp at which the Payment Link was cancelled.
	CancelledAt int64 `json:"cancelled_at"`
	// Timestamp when the Payment Link was created.
	CreatedAt int64 `json:"created_at"`
	// Currency code.
	Currency string `json:"currency"`
	// Customer details.
	Customer PaymentLinkCustomer `json:"customer"`
	// Description of the Payment Link.
	Description string `json:"description"`
	// Name of the entity. Always 'payment_link'.
	Entity string `json:"entity"`
	// Expiry timestamp.
	ExpireBy int64 `json:"expire_by"`
	// Timestamp at which the Payment Link expired.
	ExpiredAt int64 `json:"expired_at"`
	// Minimum amount for the first partial payment.
	FirstMinPartialAmount int64 `json:"first_min_partial_amount"`
	// Key-value pair that can be used to store additional information about the
	// entity.
	Notes NotesUnion `json:"notes"`
	// Notification settings.
	Notify PaymentLinkNotify `json:"notify"`
	// Payment details such as amount, payment id, Payment Link id and more. This array
	// is populated only after a payment is made by the customer or if the payment
	// fails. Until then, the value is null.
	Payments []PaymentLinkPayment `json:"payments,nullable"`
	// Reference number tagged to a Payment Link.
	ReferenceID string `json:"reference_id"`
	// Used to send reminders for the Payment Link.
	ReminderEnable bool                      `json:"reminder_enable"`
	Reminders      PaymentLinkRemindersUnion `json:"reminders"`
	// The unique short URL generated for the Payment Link.
	ShortURL string `json:"short_url"`
	// Source of the Payment Link creation (if applicable).
	Source string `json:"source"`
	// Identifier for the source (if applicable).
	SourceID string `json:"source_id"`
	// Current state of the Payment Link.
	//
	// Any of "created", "partially_paid", "expired", "cancelled", "paid".
	Status PaymentLinkStatus `json:"status"`
	// Timestamp when the Payment Link was updated.
	UpdatedAt int64 `json:"updated_at"`
	// Unique identifier for the user role through which the Payment Link was created.
	UserID string `json:"user_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		AcceptPartial         respjson.Field
		Amount                respjson.Field
		AmountPaid            respjson.Field
		CallbackMethod        respjson.Field
		CallbackURL           respjson.Field
		CancelledAt           respjson.Field
		CreatedAt             respjson.Field
		Currency              respjson.Field
		Customer              respjson.Field
		Description           respjson.Field
		Entity                respjson.Field
		ExpireBy              respjson.Field
		ExpiredAt             respjson.Field
		FirstMinPartialAmount respjson.Field
		Notes                 respjson.Field
		Notify                respjson.Field
		Payments              respjson.Field
		ReferenceID           respjson.Field
		ReminderEnable        respjson.Field
		Reminders             respjson.Field
		ShortURL              respjson.Field
		Source                respjson.Field
		SourceID              respjson.Field
		Status                respjson.Field
		UpdatedAt             respjson.Field
		UserID                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentLink) RawJSON() string { return r.JSON.raw }
func (r *PaymentLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Customer details.
type PaymentLinkCustomer struct {
	// Contact number of the customer.
	Contact string `json:"contact"`
	// Email address of the customer.
	Email string `json:"email"`
	// Name of the customer.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contact     respjson.Field
		Email       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentLinkCustomer) RawJSON() string { return r.JSON.raw }
func (r *PaymentLinkCustomer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification settings.
type PaymentLinkNotify struct {
	// Email notification enabled.
	Email bool `json:"email"`
	// SMS notification enabled.
	SMS bool `json:"sms"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		SMS         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentLinkNotify) RawJSON() string { return r.JSON.raw }
func (r *PaymentLinkNotify) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentLinkPayment struct {
	// The amount paid by the customer using the Payment Link.
	Amount int64 `json:"amount"`
	// Timestamp, in Unix, indicating when the payment was made.
	CreatedAt int64 `json:"created_at"`
	// The payment method used to make the payment.
	//
	// Any of "netbanking", "card", "wallet", "upi", "emi", "bank_transfer".
	Method string `json:"method"`
	// Unique identifier of the payment made against the Payment Link.
	PaymentID string `json:"payment_id"`
	// Unique identifier of the Payment Link.
	PlinkID string `json:"plink_id"`
	// The payment state.
	Status string `json:"status"`
	// Timestamp, in Unix, indicating when the payment was updated.
	UpdatedAt int64 `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		CreatedAt   respjson.Field
		Method      respjson.Field
		PaymentID   respjson.Field
		PlinkID     respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentLinkPayment) RawJSON() string { return r.JSON.raw }
func (r *PaymentLinkPayment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PaymentLinkRemindersUnion contains all possible properties and values from
// [map[string]string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type PaymentLinkRemindersUnion struct {
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

func (u PaymentLinkRemindersUnion) AsStringMap() (v map[string]string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PaymentLinkRemindersUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PaymentLinkRemindersUnion) RawJSON() string { return u.JSON.raw }

func (r *PaymentLinkRemindersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current state of the Payment Link.
type PaymentLinkStatus string

const (
	PaymentLinkStatusCreated       PaymentLinkStatus = "created"
	PaymentLinkStatusPartiallyPaid PaymentLinkStatus = "partially_paid"
	PaymentLinkStatusExpired       PaymentLinkStatus = "expired"
	PaymentLinkStatusCancelled     PaymentLinkStatus = "cancelled"
	PaymentLinkStatusPaid          PaymentLinkStatus = "paid"
)

type PaymentLinkListResponse struct {
	// List of Payment Links.
	PaymentLinks []PaymentLink `json:"payment_links"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaymentLinks respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentLinkListResponse) RawJSON() string { return r.JSON.raw }
func (r *PaymentLinkListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentLinkNotifyResponse struct {
	// Indicates whether the notification was sent successfully.
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentLinkNotifyResponse) RawJSON() string { return r.JSON.raw }
func (r *PaymentLinkNotifyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentLinkNewParams struct {
	// Amount to be paid using the Payment Link. Must be in the smallest unit of the
	// currency.
	Amount int64 `json:"amount,required"`
	// A three-letter ISO code for the currency in which you want to accept the
	// payment.
	Currency string `json:"currency,required"`
	// Indicates whether customers can make partial payments using the Payment Link.Not
	// allowed for UPI Payment Links.
	AcceptPartial param.Opt[bool] `json:"accept_partial,omitzero"`
	// If callback_url parameter is passed, callback_method must be passed with the
	// value get.
	CallbackMethod param.Opt[string] `json:"callback_method,omitzero"`
	// If specified, adds a redirect URL to the Payment Link.
	CallbackURL param.Opt[string] `json:"callback_url,omitzero"`
	// A brief description of the Payment Link. Max 2048 characters.
	Description param.Opt[string] `json:"description,omitzero"`
	// Timestamp, in Unix, at which the Payment Link will expire.
	ExpireBy param.Opt[int64] `json:"expire_by,omitzero"`
	// Minimum amount, in currency subunits, that must be paid by the customer as the
	// first partial payment. Must be passed along with accept_partial.
	FirstMinPartialAmount param.Opt[int64] `json:"first_min_partial_amount,omitzero"`
	// Reference number tagged to a Payment Link. Must be unique for each Payment Link.
	// Max 40 characters.
	ReferenceID param.Opt[string] `json:"reference_id,omitzero"`
	// Used to send reminders for the Payment Link.
	ReminderEnable param.Opt[bool] `json:"reminder_enable,omitzero"`
	// Must be set to true for creating a UPI Payment Link. If not passed or false, a
	// Standard Payment Link will be created.
	UpiLink param.Opt[bool] `json:"upi_link,omitzero"`
	// Customer details.
	Customer PaymentLinkNewParamsCustomer `json:"customer,omitzero"`
	// Key-value pair that can be used to store additional information about the
	// entity.
	Notes NotesUnionParam `json:"notes,omitzero"`
	// Defines who handles Payment Link notification.
	Notify PaymentLinkNewParamsNotify `json:"notify,omitzero"`
	paramObj
}

func (r PaymentLinkNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PaymentLinkNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentLinkNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Customer details.
type PaymentLinkNewParamsCustomer struct {
	// Contact number of the customer.
	Contact param.Opt[string] `json:"contact,omitzero"`
	// Email address of the customer.
	Email param.Opt[string] `json:"email,omitzero"`
	// Name of the customer.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r PaymentLinkNewParamsCustomer) MarshalJSON() (data []byte, err error) {
	type shadow PaymentLinkNewParamsCustomer
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentLinkNewParamsCustomer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines who handles Payment Link notification.
type PaymentLinkNewParamsNotify struct {
	// Send email notification.
	Email param.Opt[bool] `json:"email,omitzero"`
	// Send SMS notification.
	SMS param.Opt[bool] `json:"sms,omitzero"`
	paramObj
}

func (r PaymentLinkNewParamsNotify) MarshalJSON() (data []byte, err error) {
	type shadow PaymentLinkNewParamsNotify
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentLinkNewParamsNotify) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentLinkUpdateParams struct {
	// Indicates whether customers can make partial payments. Not allowed for UPI
	// Payment Links.
	AcceptPartial param.Opt[bool] `json:"accept_partial,omitzero"`
	// Unix timestamp when the payment link should expire.
	ExpireBy param.Opt[int64] `json:"expire_by,omitzero"`
	// Adds a unique reference number to an existing link.
	ReferenceID param.Opt[string] `json:"reference_id,omitzero"`
	// Key-value pair that can be used to store additional information about the
	// entity.
	Notes NotesUnionParam `json:"notes,omitzero"`
	paramObj
}

func (r PaymentLinkUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PaymentLinkUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentLinkUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentLinkListParams struct {
	// Unique identifier of the payment associated with the Payment Link.
	PaymentID param.Opt[string] `query:"payment_id,omitzero" json:"-"`
	// The unique reference number entered while creating the Payment Link.
	ReferenceID param.Opt[string] `query:"reference_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaymentLinkListParams]'s query parameters as `url.Values`.
func (r PaymentLinkListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PaymentLinkNotifyParams struct {
	ID string `path:"id,required" json:"-"`
	paramObj
}

type PaymentLinkNotifyParamsMedium string

const (
	PaymentLinkNotifyParamsMediumSMS   PaymentLinkNotifyParamsMedium = "sms"
	PaymentLinkNotifyParamsMediumEmail PaymentLinkNotifyParamsMedium = "email"
)
