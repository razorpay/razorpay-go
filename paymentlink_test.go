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

func TestPaymentLinkNewWithOptionalParams(t *testing.T) {
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
	_, err := client.PaymentLinks.New(context.TODO(), razorpay.PaymentLinkNewParams{
		Amount:         1000,
		Currency:       "INR",
		AcceptPartial:  razorpay.Bool(true),
		CallbackMethod: razorpay.String("get"),
		CallbackURL:    razorpay.String("https://example-callback-url.com/"),
		Customer: razorpay.PaymentLinkNewParamsCustomer{
			Contact: razorpay.String("+919000090000"),
			Email:   razorpay.String("gaurav.kumar@example.com"),
			Name:    razorpay.String("Gaurav Kumar"),
		},
		Description:           razorpay.String("Payment for policy no"),
		ExpireBy:              razorpay.Int(1691097057),
		FirstMinPartialAmount: razorpay.Int(100),
		Notes: razorpay.NotesUnionParam{
			OfStringMap: map[string]string{
				"key1": "value3",
				"key2": "value2",
			},
		},
		Notify: razorpay.PaymentLinkNewParamsNotify{
			Email: razorpay.Bool(true),
			SMS:   razorpay.Bool(true),
		},
		ReferenceID:    razorpay.String("TS1989"),
		ReminderEnable: razorpay.Bool(true),
		UpiLink:        razorpay.Bool(true),
	})
	if err != nil {
		var apierr *razorpay.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentLinkGet(t *testing.T) {
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
	_, err := client.PaymentLinks.Get(context.TODO(), "id")
	if err != nil {
		var apierr *razorpay.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentLinkUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.PaymentLinks.Update(
		context.TODO(),
		"id",
		razorpay.PaymentLinkUpdateParams{
			AcceptPartial: razorpay.Bool(false),
			ExpireBy:      razorpay.Int(1653347540),
			Notes: razorpay.NotesUnionParam{
				OfStringMap: map[string]string{
					"key1": "value3",
					"key2": "value2",
				},
			},
			ReferenceID: razorpay.String("TS35"),
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

func TestPaymentLinkListWithOptionalParams(t *testing.T) {
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
	_, err := client.PaymentLinks.List(context.TODO(), razorpay.PaymentLinkListParams{
		PaymentID:   razorpay.String("payment_id"),
		ReferenceID: razorpay.String("reference_id"),
	})
	if err != nil {
		var apierr *razorpay.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentLinkNotify(t *testing.T) {
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
	_, err := client.PaymentLinks.Notify(
		context.TODO(),
		razorpay.PaymentLinkNotifyParamsMediumSMS,
		razorpay.PaymentLinkNotifyParams{
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
