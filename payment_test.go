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

func TestPaymentGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.Get(
		context.TODO(),
		"id",
		razorpay.PaymentGetParams{
			Expand: razorpay.PaymentGetParamsExpandCard,
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

func TestPaymentUpdate(t *testing.T) {
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
	_, err := client.Payments.Update(
		context.TODO(),
		"id",
		razorpay.PaymentUpdateParams{
			Notes: razorpay.PaymentUpdateParamsNotesUnion{
				OfStringMap: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			},
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

func TestPaymentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.List(context.TODO(), razorpay.PaymentListParams{
		Count:  razorpay.Int(100),
		Expand: razorpay.PaymentListParamsExpandCard,
		From:   razorpay.Int(1593320020),
		Skip:   razorpay.Int(0),
		To:     razorpay.Int(1624856020),
	})
	if err != nil {
		var apierr *razorpay.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentCaptureWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.Capture(
		context.TODO(),
		"id",
		razorpay.PaymentCaptureParams{
			Amount:   razorpay.Int(1000),
			Currency: razorpay.String("INR"),
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
