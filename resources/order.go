package resources

import (
	"fmt"
	"net/http"

	"github.com/razorpay/razorpay-go/v2/constants"
	"github.com/razorpay/razorpay-go/v2/requests"
)

//Order ...
type Order struct {
	Request *requests.Request
}

// All fetches multiple orders for the given query params.
func (order *Order) All(queryParams *OrderList, extraHeaders map[string]string) (*OrderListResponse, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.ORDER_URL)
	var result *OrderListResponse
	err := order.Request.Call(http.MethodGet, url, queryParams, &result, extraHeaders)
	return result, err
}

// Fetch fetches an order having the given orderID.
func (order *Order) Fetch(orderID string, queryParams map[string]interface{}, extraHeaders map[string]string) (*OrderResponse, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.ORDER_URL, orderID)
	var result *OrderResponse
	err := order.Request.Call(http.MethodGet, url, queryParams, &result, extraHeaders)
	return result, err
}

// Create creates a new order for the given data
func (order *Order) Create(data *OrderRequest, extraHeaders map[string]string) (*OrderResponse, error) {
	url := fmt.Sprintf("/%s%s", constants.VERSION_V1, constants.ORDER_URL)
	var result *OrderResponse
	err := order.Request.Call(http.MethodPost, url, data, &result, extraHeaders)
	return result, err
}

// Update updates an order having the given orderID.
func (order *Order) Update(orderID string, data *OrderUpdateRequest, extraHeaders map[string]string) (*OrderResponse, error) {
	url := fmt.Sprintf("/%s%s/%s", constants.VERSION_V1, constants.ORDER_URL, orderID)
	var result *OrderResponse
	err := order.Request.Call(http.MethodPatch, url, data, &result, extraHeaders)
	return result, err
}

// Payments fetches the payments for the given orderID.
func (order *Order) Payments(orderID string, queryParams map[string]interface{}, extraHeaders map[string]string) (*OrderPaymentsResponse, error) {
	url := fmt.Sprintf("/%s%s/%s/payments", constants.VERSION_V1, constants.ORDER_URL, orderID)
	var result *OrderPaymentsResponse
	err := order.Request.Call(http.MethodGet, url, queryParams, &result, extraHeaders)
	return result, err
}

func (order *Order) ViewRtoReview(orderID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/rto_review", constants.VERSION_V1, constants.ORDER_URL, orderID)
	return order.Request.Post(url, data, extraHeaders)
}

func (order *Order) EditFulfillment(orderID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/%s%s/%s/fulfillment", constants.VERSION_V1, constants.ORDER_URL, orderID)
	return order.Request.Post(url, data, extraHeaders)
}
