## Refunds

### Create a normal refund

```go

data := map[string]interface{}{
  "speed": "normal",
  "notes": map[string]interface{}{
    "key_1": "value1",
    "key_2": "value2",
  },
}
body, err := client.Payment.Refund("<paymentId>",1200, data, nil)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
|  paymentId*   | string      | The id of the payment                       |
|  amount       | integer      | The amount to be captured (should be equal to the authorized amount, in paise) |                       |
|  speed        | string      | Here, it must be normal                |
|  notes        | object       | A key-value pair                |
|  receipt      | string      | A unique identifier provided by you for your internal reference. |

**Response:**
```json
{
  "id": "rfnd_FP8QHiV938haTz",
  "entity": "refund",
  "amount": 500100,
  "receipt": "Receipt No. 31",
  "currency": "INR",
  "payment_id": "pay_29QQoUBi66xm2f",
  "notes": [],
  "receipt": null,
  "acquirer_data": {
    "arn": null
  },
  "created_at": 1597078866,
  "batch_id": null,
  "status": "processed",
  "speed_processed": "normal",
  "speed_requested": "normal"
}
```
-------------------------------------------------------------------------------------------------------

### Create an instant refund

```go
data := map[string]interface{}{
    "speed": "optimum",
    "receipt": "Receipt No. 31",
}
body, err := client.Payment.Refund("<paymentId>",100, data, nil)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
|  paymentId*  | string      | The id of the payment                       |
|  amount       | integer      | The amount to be captured (should be equal to the authorized amount, in paise) |
|  speed*        | string      | Here, it must be optimum                    |
|  receipt      | string      | A unique identifier provided by you for your internal reference. |

**Response:**
```json
{
  "id": "rfnd_FP8R8EGjGbPkVb",
  "entity": "refund",
  "amount": 500100,
  "currency": "INR",
  "payment_id": "pay_FC8MmhMBZPKDHF",
  "notes": {
    "notes_key_1": "Tea, Earl Grey, Hot",
    "notes_key_2": "Tea, Earl Grey… decaf."
  },
  "receipt": "Receipt No. 31",
  "acquirer_data": {
    "arn": null
  },
  "created_at": 1597078914,
  "batch_id": null,
  "status": "processed",
  "speed_requested": "optimum"
}
```
-------------------------------------------------------------------------------------------------------

### Fetch multiple refunds for a payment

```go
option := map[string]interface{}{
    "count" : 1,
}
body, err := client.Payment.FetchMultipleRefund("<paymentId>",option, nil)
```

**Parameters:**

| Name  | Type      | Description                                      |
|-------|-----------|--------------------------------------------------|
| paymentId*  | string      | The id of the payment                       |
| from  | timestamp | timestamp after which the payments were created  |
| to    | timestamp | timestamp before which the payments were created |
| count | integer   | number of payments to fetch (default: 10)        |
| skip  | integer   | number of payments to be skipped (default: 0)    |

**Refund:**
```json
{
  "entity": "collection",
  "count": 1,
  "items": [
    {
      "id": "rfnd_FP8DDKxqJif6ca",
      "entity": "refund",
      "amount": 300100,
      "currency": "INR",
      "payment_id": "pay_FIKOnlyii5QGNx",
      "notes": {
        "comment": "Comment for refund"
      },
      "receipt": null,
      "acquirer_data": {
        "arn": "10000000000000"
      },
      "created_at": 1597078124,
      "batch_id": null,
      "status": "processed",
      "speed_processed": "normal",
      "speed_requested": "optimum"
    }
  ]
}
 ```
-------------------------------------------------------------------------------------------------------

### Fetch a specific refund for a payment
```go
body, err := client.Payment.FetchRefund("<paymentId>","<refundId>", nil, nil)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
|  paymentId*   | string      | The id of the payment to be fetched        |
|  refundId*   | string      | The id of the refund to be fetched           |

**Response:**
```json
{
  "id": "rfnd_FP8DDKxqJif6ca",
  "entity": "refund",
  "amount": 300100,
  "currency": "INR",
  "payment_id": "pay_FIKOnlyii5QGNx",
  "notes": {
    "comment": "Comment for refund"
  },
  "receipt": null,
  "acquirer_data": {
    "arn": "10000000000000"
  },
  "created_at": 1597078124,
  "batch_id": null,
  "status": "processed",
  "speed_processed": "normal",
  "speed_requested": "optimum"
}
```
-------------------------------------------------------------------------------------------------------

### Fetch all refunds
```go
option := map[string]interface{}{
    "count" : 1,
}
body, err := client.Payment.All(option, nil)
```

**Parameters:**

| Name  | Type      | Description                                      |
|-------|-----------|--------------------------------------------------|
| from  | timestamp | timestamp after which the payments were created  |
| to    | timestamp | timestamp before which the payments were created |
| count | integer   | number of payments to fetch (default: 10)        |
| skip  | integer   | number of payments to be skipped (default: 0)    |

**Response:**
```json
{
    "entity": "collection",
    "count": 1,
    "items": [
        {
            "id": "rfnd_Jnh48KzUkrUZ2I",
            "entity": "refund",
            "amount": 100,
            "currency": "INR",
            "payment_id": "pay_JRSS9bMrRMds3w",
            "notes": [],
            "receipt": "#999",
            "acquirer_data": {
                "arn": "Jnh4MAOSUc7w63"
            },
            "created_at": 1656589374,
            "batch_id": null,
            "status": "processed",
            "speed_processed": "instant",
            "speed_requested": "optimum"
        }
    ]
}
```
-------------------------------------------------------------------------------------------------------

### Fetch particular refund
```go
body, err := client.Refund.Fetch("<refundId>", nil, nil)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
|  refundId*   | string      | The id of the refund to be fetched           |

**Response:**
```json
{
  "id": "rfnd_EqWThTE7dd7utf",
  "entity": "refund",
  "amount": 6000,
  "currency": "INR",
  "payment_id": "pay_EpkFDYRirena0f",
  "notes": {
    "comment": "Issuing an instant refund"
  },
  "receipt": null,
  "acquirer_data": {
    "arn": "10000000000000"
  },
  "created_at": 1589521675,
  "batch_id": null,
  "status": "processed",
  "speed_processed": "optimum",
  "speed_requested": "optimum"
}
```
-------------------------------------------------------------------------------------------------------

### Update the refund
```go
data:= map[string]interface{}{
  "notes": map[string]interface{}{
    "notes_key_1":"Beam me up Scotty.",
    "notes_key_2":"Engage",
  },
}

body, err := client.Refund.Update("<refundId>", data, nil)
```

**Parameters:**

| Name  | Type      | Description                                      |
|-------|-----------|--------------------------------------------------|
| refundId*   | string    | The id of the refund to be fetched     |
| notes* | object  | A key-value pair                                 |

**Response:**
```json
{
  "id": "rfnd_FP8DDKxqJif6ca",
  "entity": "refund",
  "amount": 300100,
  "currency": "INR",
  "payment_id": "pay_FIKOnlyii5QGNx",
  "notes": {
    "notes_key_1": "Beam me up Scotty.",
    "notes_key_2": "Engage"
  },
  "receipt": null,
  "acquirer_data": {
    "arn": "10000000000000"
  },
  "created_at": 1597078124,
  "batch_id": null,
  "status": "processed",
  "speed_processed": "normal",
  "speed_requested": "optimum"
}
```
-------------------------------------------------------------------------------------------------------

**PN: * indicates mandatory fields**
<br>
<br>
**For reference click [here](https://razorpay.com/docs/api/refunds/)**
