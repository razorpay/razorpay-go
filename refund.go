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
)

// RefundService contains methods and other services that help with interacting
// with the gotue API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRefundService] method instead.
type RefundService struct {
	Options []option.RequestOption
}

// NewRefundService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRefundService(opts ...option.RequestOption) (r RefundService) {
	r = RefundService{}
	r.Options = opts
	return
}

// Use this endpoint to retrieve the refund using the id.
// [See docs](https://razorpay.com/docs/api/refunds/fetch-with-id/)
func (r *RefundService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Refund, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("refunds/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Use this endpoint to retrieve details of all refunds. By default, only the last
// 10 refunds are returned. You can use count and skip query parameters to change
// that behaviour. [See docs](https://razorpay.com/docs/api/refunds/fetch-all/)
func (r *RefundService) List(ctx context.Context, query RefundListParams, opts ...option.RequestOption) (res *RefundList, err error) {
	opts = append(r.Options[:], opts...)
	path := "refunds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Use this endpoint to update the notes field for a particular refund. Only the
// notes field can be updated.
// [See docs](https://razorpay.com/docs/api/refunds/update/)
func (r *RefundService) UpdateNotes(ctx context.Context, id string, body RefundUpdateNotesParams, opts ...option.RequestOption) (res *Refund, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("refunds/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type RefundListParams struct {
	// The number of refunds to fetch. Maximum is 100.
	Count param.Opt[int64] `query:"count,omitzero" json:"-"`
	// UNIX timestamp at which the refunds were created.
	From param.Opt[int64] `query:"from,omitzero" json:"-"`
	// The number of refunds to be skipped.
	Skip param.Opt[int64] `query:"skip,omitzero" json:"-"`
	// UNIX timestamp till which the refunds were created.
	To param.Opt[int64] `query:"to,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RefundListParams]'s query parameters as `url.Values`.
func (r RefundListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RefundUpdateNotesParams struct {
	// Key-value pair that can be used to store additional information about the
	// entity.
	Notes NotesUnionParam `json:"notes,omitzero,required"`
	paramObj
}

func (r RefundUpdateNotesParams) MarshalJSON() (data []byte, err error) {
	type shadow RefundUpdateNotesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RefundUpdateNotesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
