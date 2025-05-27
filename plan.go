// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package razorpay

import (
	"context"
	"net/http"

	"github.com/razorpay/razorpay-go/v2/internal/apijson"
	"github.com/razorpay/razorpay-go/v2/internal/requestconfig"
	"github.com/razorpay/razorpay-go/v2/option"
	"github.com/razorpay/razorpay-go/v2/packages/param"
	"github.com/razorpay/razorpay-go/v2/packages/respjson"
)

// PlanService contains methods and other services that help with interacting with
// the gotue API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlanService] method instead.
type PlanService struct {
	Options []option.RequestOption
}

// NewPlanService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPlanService(opts ...option.RequestOption) (r PlanService) {
	r = PlanService{}
	r.Options = opts
	return
}

// Use this endpoint to create a plan with frequency, interval, item details, and
// notes
func (r *PlanService) New(ctx context.Context, body PlanNewParams, opts ...option.RequestOption) (res *Plan, err error) {
	opts = append(r.Options[:], opts...)
	path := "plans"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Plan struct {
	// The unique identifier linked to a plan
	ID string `json:"id"`
	// The Unix timestamp at which the plan was created
	CreatedAt int64 `json:"created_at"`
	// The entity being created. Here, it is plan
	Entity string `json:"entity"`
	// Used together with period to define how often the customer should be charged
	Interval int64    `json:"interval"`
	Item     PlanItem `json:"item"`
	// Key-value pair that can be used to store additional information about the
	// entity.
	Notes NotesUnion `json:"notes"`
	// Used together with interval to define how often the customer should be charged.
	// Possible values are daily, weekly, monthly, yearly
	//
	// Any of "daily", "weekly", "monthly", "yearly".
	Period PlanPeriod `json:"period"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Entity      respjson.Field
		Interval    respjson.Field
		Item        respjson.Field
		Notes       respjson.Field
		Period      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Plan) RawJSON() string { return r.JSON.raw }
func (r *Plan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlanItem struct {
	// Unique identifier for the item
	ID string `json:"id"`
	// Indicates if the item is active
	Active bool `json:"active"`
	// Amount for the plan item in currency subunits
	Amount int64 `json:"amount"`
	// The Unix timestamp at which the item was created
	CreatedAt int64 `json:"created_at"`
	// ISO code for the currency
	Currency string `json:"currency"`
	// Description for the plan item
	Description string `json:"description,nullable"`
	// HSN code for the item
	HsnCode string `json:"hsn_code,nullable"`
	// Name of the plan item
	Name string `json:"name"`
	// SAC code for the item
	SacCode string `json:"sac_code,nullable"`
	// Tax group ID for the item
	TaxGroupID string `json:"tax_group_id,nullable"`
	// Tax ID for the item
	TaxID string `json:"tax_id,nullable"`
	// Indicates if tax is inclusive
	TaxInclusive bool `json:"tax_inclusive"`
	// Tax rate for the item
	TaxRate float64 `json:"tax_rate,nullable"`
	// Type of the item
	Type string `json:"type"`
	// Unit for the plan item
	Unit string `json:"unit,nullable"`
	// Unit amount for the plan item in currency subunits
	UnitAmount int64 `json:"unit_amount"`
	// The Unix timestamp at which the item was last updated
	UpdatedAt int64 `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Active       respjson.Field
		Amount       respjson.Field
		CreatedAt    respjson.Field
		Currency     respjson.Field
		Description  respjson.Field
		HsnCode      respjson.Field
		Name         respjson.Field
		SacCode      respjson.Field
		TaxGroupID   respjson.Field
		TaxID        respjson.Field
		TaxInclusive respjson.Field
		TaxRate      respjson.Field
		Type         respjson.Field
		Unit         respjson.Field
		UnitAmount   respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlanItem) RawJSON() string { return r.JSON.raw }
func (r *PlanItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Used together with interval to define how often the customer should be charged.
// Possible values are daily, weekly, monthly, yearly
type PlanPeriod string

const (
	PlanPeriodDaily   PlanPeriod = "daily"
	PlanPeriodWeekly  PlanPeriod = "weekly"
	PlanPeriodMonthly PlanPeriod = "monthly"
	PlanPeriodYearly  PlanPeriod = "yearly"
)

type PlanNewParams struct {
	// This, combined with period, defines the frequency of the plan. If the billing
	// cycle is 2 months, the value should be 2. For daily plans, the minimum value
	// should be 7
	Interval int64             `json:"interval,required"`
	Item     PlanNewParamsItem `json:"item,omitzero,required"`
	// This, combined with interval, defines the frequency of the plan. Possible values
	// are daily, weekly, monthly, quarterly, yearly
	//
	// Any of "daily", "weekly", "monthly", "quarterly", "yearly".
	Period PlanNewParamsPeriod `json:"period,omitzero,required"`
	// Key-value pair that can be used to store additional information about the
	// entity.
	Notes NotesUnionParam `json:"notes,omitzero"`
	paramObj
}

func (r PlanNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PlanNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlanNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlanNewParamsItem struct {
	// Description for the plan item
	Description param.Opt[string] `json:"description,omitzero"`
	// HSN code for the item
	HsnCode param.Opt[string] `json:"hsn_code,omitzero"`
	// SAC code for the item
	SacCode param.Opt[string] `json:"sac_code,omitzero"`
	// Tax group ID for the item
	TaxGroupID param.Opt[string] `json:"tax_group_id,omitzero"`
	// Tax ID for the item
	TaxID param.Opt[string] `json:"tax_id,omitzero"`
	// Tax rate for the item
	TaxRate param.Opt[float64] `json:"tax_rate,omitzero"`
	// Unit for the plan item
	Unit param.Opt[string] `json:"unit,omitzero"`
	// Unique identifier for the item
	ID param.Opt[string] `json:"id,omitzero"`
	// Indicates if the item is active
	Active param.Opt[bool] `json:"active,omitzero"`
	// Amount for the plan item in currency subunits
	Amount param.Opt[int64] `json:"amount,omitzero"`
	// The Unix timestamp at which the item was created
	CreatedAt param.Opt[int64] `json:"created_at,omitzero"`
	// ISO code for the currency
	Currency param.Opt[string] `json:"currency,omitzero"`
	// Name of the plan item
	Name param.Opt[string] `json:"name,omitzero"`
	// Indicates if tax is inclusive
	TaxInclusive param.Opt[bool] `json:"tax_inclusive,omitzero"`
	// Type of the item
	Type param.Opt[string] `json:"type,omitzero"`
	// Unit amount for the plan item in currency subunits
	UnitAmount param.Opt[int64] `json:"unit_amount,omitzero"`
	// The Unix timestamp at which the item was last updated
	UpdatedAt param.Opt[int64] `json:"updated_at,omitzero"`
	paramObj
}

func (r PlanNewParamsItem) MarshalJSON() (data []byte, err error) {
	type shadow PlanNewParamsItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlanNewParamsItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This, combined with interval, defines the frequency of the plan. Possible values
// are daily, weekly, monthly, quarterly, yearly
type PlanNewParamsPeriod string

const (
	PlanNewParamsPeriodDaily     PlanNewParamsPeriod = "daily"
	PlanNewParamsPeriodWeekly    PlanNewParamsPeriod = "weekly"
	PlanNewParamsPeriodMonthly   PlanNewParamsPeriod = "monthly"
	PlanNewParamsPeriodQuarterly PlanNewParamsPeriod = "quarterly"
	PlanNewParamsPeriodYearly    PlanNewParamsPeriod = "yearly"
)
