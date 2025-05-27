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

// CustomerService contains methods and other services that help with interacting
// with the gotue API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerService] method instead.
type CustomerService struct {
	Options []option.RequestOption
}

// NewCustomerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerService(opts ...option.RequestOption) (r CustomerService) {
	r = CustomerService{}
	r.Options = opts
	return
}

// Use this endpoint to create or add a customer with basic details such as name
// and contact details.
func (r *CustomerService) New(ctx context.Context, body CustomerNewParams, opts ...option.RequestOption) (res *Customer, err error) {
	opts = append(r.Options[:], opts...)
	path := "customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Use this endpoint to retrieve details of a customer with id.
// [See docs](https://razorpay.com/docs/api/customers/fetch-with-id/)
func (r *CustomerService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Customer, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Use this endpoint to edit the customer details such as name, email, and contact
// details. The combination of email and contact must be unique for every customer.
func (r *CustomerService) Update(ctx context.Context, id string, body CustomerUpdateParams, opts ...option.RequestOption) (res *Customer, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Use this endpoint to retrieve the details of all the customers.
// [See docs](https://razorpay.com/docs/api/customers/fetch-all/)
func (r *CustomerService) List(ctx context.Context, query CustomerListParams, opts ...option.RequestOption) (res *CustomerListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Customer struct {
	// Unique identifier of the customer.
	ID string `json:"id"`
	// The customer's phone number. A maximum length of 15 characters including country
	// code.
	Contact string `json:"contact"`
	// UNIX timestamp, when the customer was created.
	CreatedAt int64 `json:"created_at"`
	// The customer's email address. A maximum length of 64 characters.
	Email string `json:"email"`
	// Indicates the type of entity.
	Entity string `json:"entity"`
	// GST number linked to the customer.
	Gstin string `json:"gstin,nullable"`
	// Customer's name. Alphanumeric, with period (.), apostrophe ('), forward slash
	// (/), at (@) and parentheses allowed. The name must be between 3-50 characters in
	// length.
	Name string `json:"name"`
	// Key-value pair that can be used to store additional information about the
	// entity.
	Notes NotesUnion `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Contact     respjson.Field
		CreatedAt   respjson.Field
		Email       respjson.Field
		Entity      respjson.Field
		Gstin       respjson.Field
		Name        respjson.Field
		Notes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Customer) RawJSON() string { return r.JSON.raw }
func (r *Customer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerListResponse struct {
	// The number of customer records returned in this response.
	Count int64 `json:"count"`
	// Indicates the type of entity. Always 'collection' for this response.
	Entity string `json:"entity"`
	// List of customer objects.
	Items []Customer `json:"items"`
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
func (r CustomerListResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomerListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerNewParams struct {
	// The customer's phone number. A maximum length of 15 characters including country
	// code.
	Contact string `json:"contact,required"`
	// The customer's email address. A maximum length of 64 characters.
	Email string `json:"email,required"`
	// Customer's name. Alphanumeric value with period (.), apostrophe ('), forward
	// slash (/), at (@) and parentheses are allowed. The name must be between 3-50
	// characters in length.
	Name string `json:"name,required"`
	// Possible values: false (fetches details of existing customer), true (default,
	// throws error if customer exists).
	FailExisting param.Opt[bool] `json:"fail_existing,omitzero"`
	// Customer's GST number, if available.
	Gstin param.Opt[string] `json:"gstin,omitzero"`
	// Key-value pair that can be used to store additional information about the
	// entity.
	Notes NotesUnionParam `json:"notes,omitzero"`
	paramObj
}

func (r CustomerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerUpdateParams struct {
	// The customer's phone number. A maximum length of 15 characters including country
	// code.
	Contact param.Opt[string] `json:"contact,omitzero"`
	// The customer's email address. A maximum length of 64 characters.
	Email param.Opt[string] `json:"email,omitzero"`
	// Customer's name. Alphanumeric, with period (.), apostrophe ('), forward slash
	// (/), at (@) and parentheses allowed. The name must be between 3-50 characters in
	// length.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r CustomerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerListParams struct {
	// The number of customer records to be retrieved from the system. The default
	// value is 10. The maximum value is 100. This can be used for pagination in
	// combination with skip.
	Count param.Opt[int64] `query:"count,omitzero" json:"-"`
	// The number of customer records that must be skipped. The default value is 0.
	// This can be used for pagination in combination with count.
	Skip param.Opt[int64] `query:"skip,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CustomerListParams]'s query parameters as `url.Values`.
func (r CustomerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
