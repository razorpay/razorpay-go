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
	"github.com/razorpay/razorpay-go/v2/packages/param"
)

func TestPlanNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Plans.New(context.TODO(), razorpay.PlanNewParams{
		Interval: 1,
		Item: razorpay.PlanNewParamsItem{
			ID:           razorpay.String("item_00000000000001"),
			Active:       razorpay.Bool(true),
			Amount:       razorpay.Int(69900),
			CreatedAt:    razorpay.Int(1580219935),
			Currency:     razorpay.String("INR"),
			Description:  razorpay.String("Description for the test plan - Weekly"),
			HsnCode:      param.Null[string](),
			Name:         razorpay.String("Test plan - Weekly"),
			SacCode:      param.Null[string](),
			TaxGroupID:   param.Null[string](),
			TaxID:        param.Null[string](),
			TaxInclusive: razorpay.Bool(false),
			TaxRate:      param.Null[float64](),
			Type:         razorpay.String("plan"),
			Unit:         param.Null[string](),
			UnitAmount:   razorpay.Int(69900),
			UpdatedAt:    razorpay.Int(1580219935),
		},
		Period: razorpay.PlanNewParamsPeriodWeekly,
		Notes: razorpay.NotesUnionParam{
			OfStringMap: map[string]string{
				"key1": "value3",
				"key2": "value2",
			},
		},
	})
	if err != nil {
		var apierr *razorpay.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
