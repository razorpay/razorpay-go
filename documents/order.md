## Orders

### Create order

```go
data := map[string]interface{}{
  "amount": 50000,
  "currency": "INR",
  "receipt": "some_receipt_id",
  "partial_payment": false,
  "notes": map[string]interface{}{
      "key1": "value1",
      "key2": "value2",
    },
}
body, err := client.Order.Create(data, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
| amount*          | integer | Amount of the order to be paid                                               |
| currency*        | string  | Currency of the order. Currently only `INR` is supported.                      |
| receipt         | string  | Your system order reference id.                                              |
| partial_payment  | boolean  | Indicates whether the customer can make a partial payment. Possible values: `true` or `false`    |
| notes           | object  | A key-value pair                                                             |

**Response:**

```json
{
  "id": "order_EKwxwAgItmmXdp",
  "entity": "order",
  "amount": 50000,
  "amount_paid": 0,
  "amount_due": 50000,
  "currency": "INR",
  "receipt": "receipt#1",
  "offer_id": null,
  "status": "created",
  "attempts": 0,
  "notes": [],
  "created_at": 1582628071
}
```

-------------------------------------------------------------------------------------------------------

### Create order (Third party validation)

```go
data := map[string]interface{}{
  "amount": 500,
  "method": "netbanking",
  "receipt": "BILL13375649",
  "currency": "INR",
  "bank_account": map[string]interface{}{
    "account_number": "765432123456789",
    "name": "Gaurav Kumar",
    "ifsc": "HDFC0000053",
  },
}

body, err := client.Order.Create(data, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
| amount*          | integer | Amount of the order to be paid                                               |
| method        | string  | The payment method used to make the payment. If this parameter is not passed, customers will be able to make payments using both netbanking and UPI payment methods. Possible values is `netbanking` or `upi`|
| currency*        | string  | Currency of the order. Currently only `INR` is supported.       |
| receipt         | string  | Your system order reference id.                                              |
| notes         | object      | A key-value pair  |
|bank_account | object  | All keys listed [here](https://razorpay.com/docs/payments/third-party-validation/#step-2-create-an-order) are supported |

**Response:**

```json
{
  "id": "order_GAWN9beXgaqRyO",
  "entity": "order",
  "amount": 500,
  "amount_paid": 0,
  "amount_due": 500,
  "currency": "INR",
  "receipt": "BILL13375649",
  "offer_id": null,
  "status": "created",
  "attempts": 0,
  "notes": [],
  "created_at": 1573044247
}
```

-------------------------------------------------------------------------------------------------------

### Fetch all orders

```go
options := map[string]interface{}{
  "count": 1,
}
body, err := client.Order.All(options, nil)
```

**Parameters**

| Name       | Type      | Description                                                  |
|------------|-----------|--------------------------------------------------------------|
| from       | timestamp | timestamp after which the orders were created              |
| to         | timestamp | timestamp before which the orders were created             |
| count      | integer   | number of orders to fetch (default: 10)                    |
| skip       | integer   | number of orders to be skipped (default: 0)                |
| authorized | boolean   | Orders for which orders are currently in authorized state. |
| receipt    | string    | Orders with the provided value for receipt.                  |

**Response:**

```json
{
  "entity": "collection",
  "count": 1,
  "items": [
    {
      "id": "order_EKzX2WiEWbMxmx",
      "entity": "order",
      "amount": 1234,
      "amount_paid": 0,
      "amount_due": 1234,
      "currency": "INR",
      "receipt": "Receipt No. 1",
      "offer_id": null,
      "status": "created",
      "attempts": 0,
      "notes": [],
      "created_at": 1582637108
    }
  ]
}
```
-------------------------------------------------------------------------------------------------------

### Fetch particular order

```go
body, err := client.Order.Fetch("<orderId>", nil, nil)
```
**Parameters**

| Name     | Type   | Description                         |
|----------|--------|-------------------------------------|
| orderId* | string | The id of the order to be fetched |

**Response:**

```json
{
  "amount": 50000,
  "amount_due": 50000,
  "amount_paid": 0,
  "attempts": 0,
  "created_at": 1656571893,
  "currency": "INR",
  "entity": "order",
  "id": "order_Jnc6NatYAYnayr",
  "notes": {
    "key1": "value1",
    "key2": "value2"
  },
  "offer_id": "offer_JTUADI4ZWBGWur",
  "offers": [
    "offer_JTUADI4ZWBGWur"
  ],
  "receipt": "some_receipt_id",
  "status": "created"
}
```
-------------------------------------------------------------------------------------------------------

### Fetch payments for an order

```go
body, err := client.Order.Payments("<orderId>", nil, nil)
```
**Parameters**

| Name     | Type   | Description                         |
|----------|--------|-------------------------------------|
| orderId* | string | The id of the order for which the payments should be retrieved |

**Response:**
```json
{
    "entity": "collection",
    "count": 1,
    "items": [
        {
            "id": "pay_JmnuyQBkD6eOo7",
            "entity": "payment",
            "amount": 1000,
            "currency": "INR",
            "status": "captured",
            "order_id": "order_J7lfDm8cppLIVV",
            "invoice_id": "inv_J7lfDkGnM3zHYV",
            "international": false,
            "method": "upi",
            "amount_refunded": 0,
            "refund_status": null,
            "captured": true,
            "description": "Invoice #inv_J7lfDkGnM3zHYV",
            "card_id": null,
            "bank": null,
            "wallet": null,
            "vpa": "success@razorpay",
            "email": "sofiya@gmail.com",
            "contact": "+919702219003",
            "notes": [],
            "fee": 24,
            "tax": 4,
            "error_code": null,
            "error_description": null,
            "error_source": null,
            "error_step": null,
            "error_reason": null,
            "acquirer_data": {
                "rrn": "368096450516",
                "upi_transaction_id": "97505784AB5852CEEBFE6BE418BB2373"
            },
            "created_at": 1656395165,
            "account_id": "acc_Hn1ukn2d32Fqww"
        }
    ]
}
```
-------------------------------------------------------------------------------------------------------

### Update order

```go
data := map[string]interface{}{
  "notes": map[string]interface{}{
        "notes_key_1": "value1",
        "notes_key_2": "value2",
      }, 
}
body, err := client.Order.Update("<orderId>", data, nil)
```
**Parameters**

| Name     | Type   | Description                         |
|----------|--------|-------------------------------------|
| orderId* | string | The id of the order in which the Notes field must be updated. |
| notes*   | object  | A key-value pair                    |

**Response:**
```json
{
  "id":"order_DaaS6LOUAASb7Y",
  "entity":"order",
  "amount":2200,
  "amount_paid":0,
  "amount_due":2200,
  "currency":"INR",
  "receipt":"Receipt #211",
  "offer_id":null,
  "status":"attempted",
  "attempts":1,
  "notes":{
    "notes_key_1":"Tea, Earl Grey, Hot",
    "notes_key_2":"Tea, Earl Grey… decaf."
  },
  "created_at":1572505143
}
```
-------------------------------------------------------------------------------------------------------

### View RTO/Risk Reasons

```go
orderId = "order_DaaS6LOUAASb7Y"

body, err := client.Order.ViewRtoReview(TestOrderID, nil, nil)
```
**Parameters**

| Name     | Type   | Description                         |
|----------|--------|-------------------------------------|
| orderId* | string | The unique identifier of an order to access the fulfillment information. |

**Response:**
```json
{
  "risk_tier": "high",
  "rto_reasons": [
    {
      "reason": "short_shipping_address",
      "description": "Short shipping address",
      "bucket": "address"
    },
    {
      "reason": "address_pincode_state_mismatch",
      "description": "Incorrect pincode state entered",
      "bucket": "address"
    }
  ]
}
```
-------------------------------------------------------------------------------------------------------

### Update the Fulfillment Details

```go
params := map[string]interface{}{
  "payment_method": "upi",
  "shipping": map[string]interface{}{
    "waybill":  "123456789",
    "status":   "rto",
    "provider": "Bluedart",
  },
}

body, err := client.Order.EditFulfillment(orderId, params, nil)
```
**Parameters**

| Name     | Type   | Description                         |
|----------|--------|-------------------------------------|
| orderId* | string | The unique identifier of an order to access the fulfillment information. |
| payment_method | string | Payment Method opted by the customer to complete the payment. Possible values is `upi`, `card`, `wallet`, `netbanking`, `cod`, `emi`, `cardless_emi`, `paylater`, `recurring`, `other`. |
| shipping | object  | Contains the shipping data. [here](https://razorpay.com/docs/payments/magic-checkout/rto-intelligence/#step-3-update-the-fulfillment-details) are supported |

**Response:**
```json
{
  "entity": "order.fulfillment",
  "order_id": "EKwxwAgItXXXX",
  "payment_method": "upi",
  "shipping": {
    "waybill": "123456789",
    "status": "rto",
    "provider": "Bluedart"
  }
}
```
-------------------------------------------------------------------------------------------------------

**PN: * indicates mandatory fields**
<br>
<br>
**For reference click [here](https://razorpay.com/docs/api/orders/)**