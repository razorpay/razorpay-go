## Generic Entity


```go
client := razorpay.NewClient("key", "secret")
```

### Method Signature
```go
body, err := client.Entity.Do(url, method, data, headers)
```    

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| url*          | string      | The endpoint to which the request will be made. This should include the version information (e.g., "v1/contacts" or "v2/accounts")  |
| method*       | string      | The HTTP method for the request (e.g., 'GET', 'POST', 'PUT', 'PATCH', 'DELETE'). |
| data          | object    | The data to be sent with the request.|
| headers         | object      | The headers to be included in the request. |

-------------------------------------------------------------------------------------------------------

### Create a contacts using POST

```go

data := map[string]interface{}{
  "name": "Gaurav Kumar",
  "email": "gaurav.kumar@example.com",
  "contact": "9123456789",
  "type": "employee",
  "reference_id":"Acme Contact ID 12345",
  "notes": map[string]interface{}{
    "notes_key_1":"Tea, Earl Grey, Hot",
    "notes_key_2":"Tea, Earl Grey… decaf.",
  },
}

body, err := client.Entity.Do("v1/contacts", http.MethodPost, data, nil)
```

**Response:**

```json
{
  "id": "cont_00000000000001",
  "entity": "contact",
  "name": "Gaurav Kumar",
  "contact": "9123456789",
  "email": "gaurav.kumar@example.com",
  "type": "employee",
  "reference_id": "Acme Contact ID 12345",
  "batch_id": null,
  "active": true,
  "notes": {
    "notes_key_1": "Tea, Earl Grey, Hot",
    "notes_key_2": "Tea, Earl Grey… decaf."
  },
  "created_at": 1545320320
}
```

-------------------------------------------------------------------------------------------------------

### Fetch an order using GET

```go

body, err := client.Entity.Do("v1/orders/order_00000000000001", http.MethodGet, nil, nil)
```

**Response:**

```json
{
  "amount": 307,
  "amount_due": 0,
  "amount_paid": 307,
  "attempts": 1,
  "created_at": 1695625101,
  "currency": "INR",
  "entity": "order",
  "id": "order_00000000000001",
  "notes": [],
  "offer_id": null,
  "receipt": "851617",
  "status": "paid"
}
```

-------------------------------------------------------------------------------------------------------

### Fetch payments of a linked account using headers

```go
headers := map[string]string{
  "X-Razorpay-Account": "acc_00000000000001",
}

body, _ := client.Entity.Do("v1/payments/pay_00000000000001", http.MethodGet, nil, headers)
```

**Response:**

```json
{
  "entity": "collection",
  "count": 2,
  "items": [
    {
      "id": "pay_00000000000001",
      "entity": "payment",
      "amount": 10000,
      "currency": "INR",
      "status": "captured",
      "order_id": "order_00000000000001",
      "invoice_id": null,
      "international": false,
      "method": "netbanking",
      "amount_refunded": 0,
      "refund_status": null,
      "captured": true,
      "description": "#JJCqaOhFihfkVE",
      "card_id": null,
      "bank": "YESB",
      "wallet": null,
      "vpa": null,
      "email": "john.example@example.com",
      "contact": "9999999999",
      "notes": [],
      "fee": 236,
      "tax": 36,
      "error_code": null,
      "error_description": null,
      "error_source": null,
      "error_step": null,
      "error_reason": null,
      "acquirer_data": {
        "bank_transaction_id": "2118867"
      },
      "created_at": 1649932775
    }
  ]
}
```

-------------------------------------------------------------------------------------------------------

**PN: * indicates mandatory fields**
<br>
<br>
