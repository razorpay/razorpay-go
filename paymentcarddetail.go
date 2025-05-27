// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package razorpay

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/razorpay/razorpay-go/v2/internal/requestconfig"
	"github.com/razorpay/razorpay-go/v2/option"
)

// PaymentCardDetailService contains methods and other services that help with
// interacting with the gotue API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentCardDetailService] method instead.
type PaymentCardDetailService struct {
	Options []option.RequestOption
}

// NewPaymentCardDetailService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPaymentCardDetailService(opts ...option.RequestOption) (r PaymentCardDetailService) {
	r = PaymentCardDetailService{}
	r.Options = opts
	return
}

// Use this endpoint to retrieve the details of the card used to make a payment.
// [See docs](https://razorpay.com/docs/api/payments/fetch-card-details-payment)
func (r *PaymentCardDetailService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Card, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("payments/%s/card", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
