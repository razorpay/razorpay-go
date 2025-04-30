# Razorpay Go SDK

**Official Go bindings for the Razorpay API**

This SDK allows you to seamlessly integrate Razorpay APIs into your Go applications for managing payments, orders, customers, and more — programmatically.

> 💡 **New to Razorpay?**  
> Check out the [Getting Started Guide](https://razorpay.com/docs/get-started/) to understand payment flows and set up your account.

---

## Documentation

For detailed API references and integration workflows, visit the [Razorpay API Documentation](https://razorpay.com/docs/api/).

---

## Requirements

- Go 1.18 or later

### Installation

```bash
go get github.com/razorpay/razorpay-go
```

## Usage
**Initializing the Client**

The `NewClient` method creates a new instance of the Razorpay client, which is used to make API calls.

```go
import (
razorpay "github.com/razorpay/razorpay-go"
)

client := razorpay.NewClient("<YOUR_API_KEY>", "<YOUR_API_SECRET>")
```

- <YOUR_API_KEY>: Your Razorpay API Key ID.

- <YOUR_API_SECRET>: Your Razorpay API Key Secret.

You can find your API keys in the [Razorpay Dashboard](https://dashboard.razorpay.com/#/app/keys).

Once initialized, client provides access to different Razorpay resources like Orders, Payments, Customers, etc.

For example:

### Create an Order
```go
request := &resources.OrderRequest{
	Amount:                1000,
	Currency:              "INR",
	PartialPayment:        true,
	FirstPaymentMinAmount: 100,
	Notes: resources.Notes{
		"policy_name": "Jeevan Bima",
	},
}

resp, err := client.Order.Create(request, nil)

if err != nil {
    fmt.Printf("Error creating order: %v\n", err)
    return
}
fmt.Printf("Create order Response:\n%s\n", resp.ID)
```

### Update an Order
```go
request := &resources.OrderUpdateRequest{
    Notes: resources.Notes{
        "policy_name": "Jeevan Bima",
    },
}
resp, err := client.Order.Update("order_QOSJNLwLsRK0bB", request, nil)

if err != nil {
    fmt.Printf("Error updating order: %v\n", err)
    return
}
fmt.Printf("Updating order Response:\n%s\n", resp.ID)
```

### Fetch an Order
```go
response, err := client.Order.Fetch("order_QOSJNLwLsRK0bB", nil, nil)

if err != nil {
    fmt.Printf("Error fetching order: %v\n", err)
    return
}
fmt.Printf("Fetching order Response:\n%s\n", response.ID)
```


## License

The Razorpay Go SDK is released under the MIT License. See [LICENSE](LICENSE) file for more details.