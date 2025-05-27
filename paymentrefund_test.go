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

func TestPaymentRefundNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.Refunds.New(
		context.TODO(),
		"id",
		razorpay.PaymentRefundNewParams{
			Amount: razorpay.Int(500100),
			Notes: razorpay.NotesUnionParam{
				OfStringMap: map[string]string{
					"key1": "value3",
					"key2": "value2",
				},
			},
			Receipt: razorpay.String("Receipt No. 31"),
			Speed:   razorpay.PaymentRefundNewParamsSpeedOptimum,
		},
	)
	if err != nil {
		var apierr *razorpay.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentRefundGet(t *testing.T) {
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
	_, err := client.Payments.Refunds.Get(
		context.TODO(),
		"refund_id",
		razorpay.PaymentRefundGetParams{
			ID: "id",
		},
	)
	if err != nil {
		var apierr *razorpay.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentRefundListWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.Refunds.List(
		context.TODO(),
		"id",
		razorpay.PaymentRefundListParams{
			Count: razorpay.Int(100),
			From:  razorpay.Int(0),
			Skip:  razorpay.Int(0),
			To:    razorpay.Int(0),
		},
	)
	if err != nil {
		var apierr *razorpay.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
