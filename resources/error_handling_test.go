package resources_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	razorpay "github.com/razorpay/razorpay-go"
	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/errors"
	"github.com/stretchr/testify/assert"
)

func startErrorMockServer(url string, statusCode int, errorCode string, description string) (func(), *razorpay.Client) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	client := razorpay.NewClient("test_key", "test_secret")
	client.Request.BaseURL = server.URL

	errorJSON := fmt.Sprintf(`{"error":{"internal_error_code":"%s","description":"%s","code":"%s"}}`,
		errorCode, description, errorCode)

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write([]byte(errorJSON))
	})

	return func() { server.Close() }, client
}

func startSuccessMockServer(url string, responseJSON string) (func(), *razorpay.Client) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	client := razorpay.NewClient("test_key", "test_secret")
	client.Request.BaseURL = server.URL

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(responseJSON))
	})

	return func() { server.Close() }, client
}

func startRawErrorMockServer(url string, statusCode int, rawBody string) (func(), *razorpay.Client) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	client := razorpay.NewClient("test_key", "test_secret")
	client.Request.BaseURL = server.URL

	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write([]byte(rawBody))
	})

	return func() { server.Close() }, client
}

func TestBadRequestError_400(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.ORDER_URL)
	teardown, client := startErrorMockServer(url, 400, "BAD_REQUEST_ERROR", "The amount field is required")
	defer teardown()

	_, err := client.Order.Create(map[string]interface{}{}, nil)

	assert.NotNil(t, err, "Should return an error for 400 response")
	_, ok := err.(*errors.BadRequestError)
	assert.True(t, ok, "Should return BadRequestError for 400")
}

func TestBadRequestError_InvalidParams(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.PAYMENT_URL, "pay_invalid")
	teardown, client := startErrorMockServer(url, 400, "BAD_REQUEST_ERROR", "Invalid payment id")
	defer teardown()

	_, err := client.Payment.Fetch("pay_invalid", nil, nil)

	assert.NotNil(t, err, "Should return an error for 400 response")
	_, ok := err.(*errors.BadRequestError)
	assert.True(t, ok, "Should return BadRequestError")
}

func TestServerError_500(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.ORDER_URL)
	teardown, client := startErrorMockServer(url, 500, "SERVER_ERROR", "Internal server error occurred")
	defer teardown()

	_, err := client.Order.All(nil, nil)

	assert.NotNil(t, err, "Should return an error for 500 response")
	_, ok := err.(*errors.ServerError)
	assert.True(t, ok, "Should return ServerError for 500")
}

func TestGatewayError_502(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.PAYMENT_URL)
	teardown, client := startErrorMockServer(url, 502, "GATEWAY_ERROR", "Payment gateway is unavailable")
	defer teardown()

	_, err := client.Payment.All(nil, nil)

	assert.NotNil(t, err, "Should return an error for 502 response")
	_, ok := err.(*errors.GatewayError)
	assert.True(t, ok, "Should return GatewayError for 502")
}

func TestUnknownError_DefaultsToBadRequest(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.ORDER_URL, "order_123")
	teardown, client := startErrorMockServer(url, 422, "UNKNOWN_ERROR_CODE", "Some unknown error")
	defer teardown()

	_, err := client.Order.Fetch("order_123", nil, nil)

	assert.NotNil(t, err, "Should return an error for unknown error code")
	// Unknown error codes should default to BadRequestError
	_, ok := err.(*errors.BadRequestError)
	assert.True(t, ok, "Unknown error codes should default to BadRequestError")
}

func TestSuccessResponse_NoError(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.ORDER_URL, "order_123")
	teardown, client := startSuccessMockServer(url, `{"id": "order_123", "amount": 50000}`)
	defer teardown()

	result, err := client.Order.Fetch("order_123", nil, nil)

	assert.Nil(t, err, "Should not return error for 200 response")
	assert.NotNil(t, result, "Should return result for 200 response")
	assert.Equal(t, "order_123", result["id"], "Should parse response correctly")
}

func TestEmptyResponseBody(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.ORDER_URL, "order_123")
	teardown, client := startSuccessMockServer(url, "")
	defer teardown()

	result, err := client.Order.Fetch("order_123", nil, nil)

	assert.Nil(t, err, "Should not return error for empty 200 response")
	assert.NotNil(t, result, "Should return empty map, not nil")
	assert.Empty(t, result, "Result should be empty")
}

func TestEmptyJsonObjectResponse(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.ORDER_URL, "order_123")
	teardown, client := startSuccessMockServer(url, `{}`)
	defer teardown()

	result, err := client.Order.Fetch("order_123", nil, nil)

	assert.Nil(t, err, "Should not return error for empty JSON object")
	assert.NotNil(t, result, "Should return empty map")
}

func TestMalformedJsonResponse(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.ORDER_URL, "order_123")
	teardown, client := startSuccessMockServer(url, `{invalid json here`)
	defer teardown()

	_, err := client.Order.Fetch("order_123", nil, nil)

	assert.NotNil(t, err, "Should return error for malformed JSON")
}

// Regression tests for #116: a non-2xx response must never produce an error
// whose Error() is the empty string.

func TestNonSuccessMalformedJSONHasNonEmptyError(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.ORDER_URL, "order_123")
	teardown, client := startRawErrorMockServer(url, 400, `{invalid json`)
	defer teardown()

	_, err := client.Order.Fetch("order_123", nil, nil)

	assert.NotNil(t, err, "Should return an error for malformed non-2xx body")
	assert.NotEmpty(t, err.Error(), "Error() must not be empty")
}

func TestNonSuccessWithoutErrorObjectHasNonEmptyError(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.ORDER_URL, "order_123")
	teardown, client := startRawErrorMockServer(url, 400, `{"message":"forbidden"}`)
	defer teardown()

	_, err := client.Order.Fetch("order_123", nil, nil)

	assert.NotNil(t, err, "Should return an error when body has no error object")
	assert.NotEmpty(t, err.Error(), "Error() must not be empty")
}

func TestNonSuccessErrorWithoutDescriptionHasNonEmptyError(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.ORDER_URL, "order_123")
	teardown, client := startRawErrorMockServer(url, 400, `{"error":{"code":"BAD_REQUEST_ERROR"}}`)
	defer teardown()

	_, err := client.Order.Fetch("order_123", nil, nil)

	assert.NotNil(t, err, "Should return an error when description is absent")
	assert.NotEmpty(t, err.Error(), "Error() must not be empty")
}

func TestNonSuccessEmptyBodyHasNonEmptyError(t *testing.T) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.ORDER_URL, "order_123")
	teardown, client := startRawErrorMockServer(url, 502, "")
	defer teardown()

	_, err := client.Order.Fetch("order_123", nil, nil)

	assert.NotNil(t, err, "Should return an error for an empty non-2xx body")
	assert.NotEmpty(t, err.Error(), "Error() must not be empty")
}

func TestWellFormedErrorMessageIsUnchanged(t *testing.T) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.ORDER_URL)
	teardown, client := startErrorMockServer(url, 400, "BAD_REQUEST_ERROR", "amount required")
	defer teardown()

	_, err := client.Order.Create(map[string]interface{}{}, nil)

	assert.NotNil(t, err, "Should return an error for 400 response")
	assert.Equal(t, "amount required", err.Error(), "Well-formed error message must be unchanged")
}
