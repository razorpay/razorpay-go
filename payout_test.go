// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package razorpay_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/razorpay/razorpay-go/v2"
	"github.com/razorpay/razorpay-go/v2/internal/testutil"
	"github.com/razorpay/razorpay-go/v2/option"
)

func TestPayoutGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := razorpay.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	_, err := client.Payouts.Get(context.TODO(), "id")
	if err != nil {
		var apierr *razorpay.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPayoutListWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := razorpay.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	_, err := client.Payouts.List(context.TODO(), razorpay.PayoutListParams{
		AccountNumber: "account_number",
		ContactID:     razorpay.String("contact_id"),
		Count:         razorpay.Int(100),
		From:          razorpay.Int(0),
		FundAccountID: razorpay.String("fund_account_id"),
		Mode:          razorpay.PayoutListParamsModeNeft,
		ReferenceID:   razorpay.String("reference_id"),
		Skip:          razorpay.Int(0),
		Status:        razorpay.PayoutListParamsStatusQueued,
		To:            razorpay.Int(0),
	})
	if err != nil {
		var apierr *razorpay.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
