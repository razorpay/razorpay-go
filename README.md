# Razorpay Go Client

Golang bindings for interacting with the Razorpay API

This is primarily meant for merchants who wish to perform interactions with the Razorpay API programatically

## Usage
You need to setup your key and secret using the following:
You can find your API keys at <https://dashboard.razorpay.com/#/app/keys>.

```go
import (
razorpay "github.com/razorpay/razorpay-go"
)

client := razorpay.NewClient("<YOUR_API_KEY>", "<YOUR_API_SECRET>")

```

Note: All methods below return a `map[string]interface{}` and `error`. All methods accept an `extraHeaders` param of type `map[string]string`, allowing you to optinally set extra HTTP headers on the request.

### Payments

- Fetch all payments
    ```go
    body, err := client.Payment.All(nil, nil)
    ```
- Fetch a particular payment
    ```go
    body, err := client.Payment.Fetch("<payment_id>" , nil, nil)
    ```
- Capture a payment
    ```go
    body, err := client.Payment.Capture("<payment_id>", <amount>, nil, nil)
    ```
    Note: amount is in paisa
- Refund a payment
    ```go
    body, err := client.Payment.Refund("<payment_id>", <amount_to_be_refunded>, nil, nil)
    ```

### Refunds
- Fetch all refunds
    ```go
    body, err := client.Refund.All(nil, nil)
    ```
- Fetch a particular refund
    ```go
    body, err := client.Refund.Fetch("<refund_id>", nil, nil)
    ```

### Orders
- Create a new order

    ```go
    data := map[string]interface{}{
        "amount":          1234,
        "currency":        "INR",
        "receipt":         "some_receipt_id",
        "payment_capture": 1,
    }
    body, err := client.Order.Create(data, nil)
    ```
    Note: data is a map and should contain these keys. See the [Orders entity reference](https://razorpay.com/docs/api/orders/#order-entity) for details on these keys.

        amount           : amount of order (in paisa)
        currency         : currency of order
        receipt          : receipt id of order
        payment_capture  : 1 if capture should be done automatically or else 0
        notes(optional)  : optional notes for order

- Fetch a particular order
    ```go
    body, err := client.Order.Fetch("<order_id>", nil, nil)
    ```
- Fetch all orders
    ```go
    body, err := client.Order.All(nil, nil)
    ```
- Fetch all payments for order
    ```go
    body, err := client.Order.Payments("<order_id>", nil, nil)
    ```
