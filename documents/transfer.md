## Transfers

### Create transfers from payment

```go
data:= map[string]interface{}{
  "transfers": []interface{}{
    map[string]interface{}{
      "account": "acc_HjVXbtpSCIxENR",
      "amount": 100,
      "currency": "INR",
      "notes": map[string]interface{}{
        "name": "Gaurav Kumar",
        "roll_no": "IEC2011025",
      },
      "linked_account_notes": []string{
        "roll_no",
      },
      "on_hold": true,
      "on_hold_until": 1671222870,
    },
  },
}
body, err := client.Payment.Transfer("<paymentId>", data, nil)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| paymentId*   | string      | The id of the payment to be fetched  |
| transfers   | object     | All parameters listed [here](https://razorpay.com/docs/api/route/#create-transfers-from-payments) are supported |

**Response:**
```json
{
  "entity": "collection",
  "count": 1,
  "items": [
    {
      "id": "trf_E9uhYLFLLZ2pks",
      "entity": "transfer",
      "source": "pay_E8JR8E0XyjUSZd",
      "recipient": "acc_CPRsN1LkFccllA",
      "amount": 100,
      "currency": "INR",
      "amount_reversed": 0,
      "notes": {
        "name": "Gaurav Kumar",
        "roll_no": "IEC2011025"
      },
      "on_hold": true,
      "on_hold_until": 1671222870,
      "recipient_settlement_id": null,
      "created_at": 1580218356,
      "linked_account_notes": [
        "roll_no"
      ],
      "processed_at": 1580218357
    }
  ]
}
```
-------------------------------------------------------------------------------------------------------

### Create transfers from order

```go
data:= map[string]interface{}{
  "amount": 2000,
  "currency": "INR",
  "transfers": []interface{}{
    map[string]interface{}{
      "account": "acc_HjVXbtpSCIxENR",
      "amount": 100,
      "currency": "INR",
      "notes": map[string]interface{}{
        "branch": "Acme Corp Bangalore North",
        "name": "Gaurav Kumar",
      },
      "linked_account_notes": []string{
        "branch",
      },
      "on_hold": true,
      "on_hold_until": 1671222870,
    },
  },
}
body, err := client.Order.Create(data, nil)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| amount*   | integer      | The transaction amount, in paise |
| currency*   | string  | The currency of the payment (defaults to INR)  |
|  receipt      | string      | A unique identifier provided by you for your internal reference. |
| transfers   | object     | All parameters listed [here](https://razorpay.com/docs/api/route/#create-transfers-from-orders) are supported |

**Response:**
```json
{
  "id": "order_E9uTczH8uWPCyQ",
  "entity": "order",
  "amount": 2000,
  "amount_paid": 0,
  "amount_due": 2000,
  "currency": "INR",
  "receipt": null,
  "offer_id": null,
  "status": "created",
  "attempts": 0,
  "notes": [],
  "created_at": 1580217565,
  "transfers": [
    {
      "recipient": "acc_CPRsN1LkFccllA",
      "amount": 1000,
      "currency": "INR",
      "notes": {
        "branch": "Acme Corp Bangalore North",
        "name": "Gaurav Kumar"
      },
      "linked_account_notes": [
        "branch"
      ],
      "on_hold": true,
      "on_hold_until": 1671222870
    },
    {
      "recipient": "acc_CNo3jSI8OkFJJJ",
      "amount": 1000,
      "currency": "INR",
      "notes": {
        "branch": "Acme Corp Bangalore South",
        "name": "Saurav Kumar"
      },
      "linked_account_notes": [
        "branch"
      ],
      "on_hold": false,
      "on_hold_until": null
    }
  ]
}
```
-------------------------------------------------------------------------------------------------------

### Direct transfers

```go
data:= map[string]interface{}{
  "account": "<accountId>",
  "amount": 100,
  "currency": "INR",
}
body, err := client.Transfer.Create(data, nil)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| accountId*   | string      | The id of the account to be fetched  |
| amount*   | integer      | The amount to be captured (should be equal to the authorized amount, in paise) |
| currency*   | string  | The currency of the payment (defaults to INR)  |

**Response:**
```json
{
  "id": "trf_E9utgtfGTcpcmm",
  "entity": "transfer",
  "source": "acc_CJoeHMNpi0nC7k",
  "recipient": "acc_CPRsN1LkFccllA",
  "amount": 100,
  "currency": "INR",
  "amount_reversed": 0,
  "notes": [],
  "fees": 1,
  "tax": 0,
  "on_hold": false,
  "on_hold_until": null,
  "recipient_settlement_id": null,
  "created_at": 1580219046,
  "linked_account_notes": [],
  "processed_at": 1580219046
}
```
-------------------------------------------------------------------------------------------------------

### Fetch transfer for a payment

```go
body, err := client.Payment.Transfers("<paymentId>", nil, nil)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| paymentId*   | string      | The id of the payment to be fetched  |

**Response:**
```json
{
    "entity": "collection",
    "count": 1,
    "items": [
        {
            "id": "trf_JGQjgcy8zHFq7e",
            "entity": "transfer",
            "status": "partially_reversed",
            "source": "order_JGQjgaUikLJo8n",
            "recipient": "acc_HalyQGZh9ZyiGg",
            "amount": 500,
            "currency": "INR",
            "amount_reversed": 100,
            "fees": 1,
            "tax": 0,
            "notes": {
                "branch": "Acme Corp Bangalore South",
                "name": "Saurav Kumar"
            },
            "linked_account_notes": [
                "branch"
            ],
            "on_hold": true,
            "on_hold_until": 1679691505,
            "settlement_status": "on_hold",
            "recipient_settlement_id": null,
            "created_at": 1649326643,
            "processed_at": 1649326701,
            "error": {
                "code": null,
                "description": null,
                "reason": null,
                "field": null,
                "step": null,
                "id": "trf_JGQjgcy8zHFq7e",
                "source": null,
                "metadata": null
            }
        }
    ]
}
```
-------------------------------------------------------------------------------------------------------

### Fetch transfer for an order

```go
data := map[string]interface{}{
  "expand[]": "transfers",
}
body, err := client.Order.Fetch("<orderId>", data, nil)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| orderId*   | string      | The id of the order to be fetched  |
| expand*   | string    | Supported value is `status`  |

**Response:**
```json
{
  "id": "order_DSkl2lBNvueOly",
  "entity": "order",
  "amount": 1000,
  "amount_paid": 1000,
  "amount_due": 0,
  "currency": "INR",
  "receipt": null,
  "offer_id": null,
  "status": "paid",
  "attempts": 1,
  "notes": [],
  "created_at": 1570794714,
  "transfers": {
    "entity": "collection",
    "count": 1,
    "items": [
      {
        "id": "trf_DSkl2lXWbiADZG",
        "entity": "transfer",
        "source": "order_DSkl2lBNvueOly",
        "recipient": "acc_CNo3jSI8OkFJJJ",
        "amount": 500,
        "currency": "INR",
        "amount_reversed": 0,
        "notes": {
          "branch": "Acme Corp Bangalore North",
          "name": "Gaurav Kumar"
        },
        "fees": 2,
        "tax": 0,
        "on_hold": true,
        "on_hold_until": 1670776632,
        "recipient_settlement_id": null,
        "created_at": 1570794714,
        "linked_account_notes": [
          "Acme Corp Bangalore North"
        ],
        "processed_at": 1570794772
      }
    ]
  }
}

```
-------------------------------------------------------------------------------------------------------

### Fetch transfer

```go
body, err := client.Transfer.Fetch("<transferId>", nil, nil)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| transferId*   | string      | The id of the transfer to be fetched  |

**Response:**
```json
{
    "id": "trf_IJOI2DHWQYwqU3",
    "entity": "transfer",
    "status": "created",
    "source": "order_IJOI2CD6CNIywP",
    "recipient": "acc_HjVXbtpSCIxENR",
    "amount": 100,
    "currency": "INR",
    "amount_reversed": 0,
    "fees": 0,
    "tax": null,
    "notes": {
        "branch": "Acme Corp Bangalore North",
        "name": "Gaurav Kumar"
    },
    "linked_account_notes": [
        "branch"
    ],
    "on_hold": true,
    "on_hold_until": 1671222870,
    "settlement_status": null,
    "recipient_settlement_id": null,
    "created_at": 1636435963,
    "processed_at": null,
    "error": {
        "code": null,
        "description": null,
        "reason": null,
        "field": null,
        "step": null,
        "id": "trf_IJOI2DHWQYwqU3",
        "source": null,
        "metadata": null
    }
}
```
-------------------------------------------------------------------------------------------------------

### Fetch transfers for a settlement

```go
data:= map[string]interface{}{
  "recipient_settlement_id": "<recipientSettlementId>",
}
body, err := client.Transfer.All(data, nil)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| recipientSettlementId*   | string    | The recipient settlement id obtained from the settlement.processed webhook payload.  |

**Response:**
```json
{
  "entity": "collection",
  "count": 1,
  "items": [
    {
      "id": "trf_DGSTeXzBkEVh48",
      "entity": "transfer",
      "source": "pay_DGSRhvMbOqeCe7",
      "recipient": "acc_CMaomTz4o0FOFz",
      "amount": 500,
      "currency": "INR",
      "amount_reversed": 0,
      "notes": [],
      "fees": 2,
      "tax": 0,
      "on_hold": false,
      "on_hold_until": null,
      "recipient_settlement_id": "setl_DHYJ3dRPqQkAgV",
      "created_at": 1568110256,
      "linked_account_notes": [],
      "processed_at": null
    }
  ]
}
```
-------------------------------------------------------------------------------------------------------

### Fetch settlement details

```go
data:= map[string]interface{}{
  "expand[]": "recipient_settlement",
}
body, err := client.Transfer.All(data, nil)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| expand*   | string    | Supported value is `recipient_settlement`  |

**Response:**
```json
{
    "entity": "collection",
    "count": 1,
    "items": [
        {
            "id": "trf_IGLlpyrGLqWb7u",
            "entity": "transfer",
            "status": "processed",
            "source": "pay_IGJYtFZPiKj5FE",
            "recipient": "acc_HSUD5wqmJ0MTDI",
            "amount": 100,
            "currency": "INR",
            "amount_reversed": 0,
            "fees": 0,
            "tax": 0,
            "notes": [],
            "linked_account_notes": [],
            "on_hold": true,
            "on_hold_until": 1636654800,
            "settlement_status": "on_hold",
            "recipient_settlement_id": null,
            "recipient_settlement": null,
            "created_at": 1635772070,
            "processed_at": 1635772075,
            "error": {
                "code": null,
                "description": null,
                "reason": null,
                "field": null,
                "step": null,
                "id": "trf_IGLlpyrGLqWb7u",
                "source": null,
                "metadata": null
            }
        }
    ]
}
```
-------------------------------------------------------------------------------------------------------

### Refund payments and reverse transfer from a linked account

```go
data:= map[string]interface{}{
  "reverse_all": 1,
}
body, err := client.Payment.Refund("<paymentId>", <amount>, data, nil)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| paymentId*   | string      | The id of the payment to be fetched  |
| amount*   | integer      | The amount to be captured (should be equal to the authorized amount, in paise) |
| reverse_all   | boolean    | Reverses transfer made to a linked account. Possible values:<br> * `1` - Reverses transfer made to a linked account.<br>* `0` - Does not reverse transfer made to a linked account.|

**Response:**
```json
{
  "id": "rfnd_JJFNlNXPHY640A",
  "entity": "refund",
  "amount": 100,
  "currency": "INR",
  "payment_id": "pay_JJCqynf4fQS0N1",
  "notes": [],
  "receipt": null,
  "acquirer_data": {
    "arn": null
  },
  "created_at": 1649941680,
  "batch_id": null,
  "status": "processed",
  "speed_processed": "normal",
  "speed_requested": "normal"
}
```
-------------------------------------------------------------------------------------------------------

### Fetch payments of a linked account

```go
extraHeaders:= map[string]string{
  "X-Razorpay-Account": "<linkedAccountId>",
}
body, err := client.Payment.All(nil, extraHeaders)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| X-Razorpay-Account   | string      | The linked account id to fetch the payments received by linked account |

**Response:**
```json
{
  "entity": "collection",
  "count": 2,
  "items": [
    {
      "id": "pay_JJCqynf4fQS0N1",
      "entity": "payment",
      "amount": 10000,
      "currency": "INR",
      "status": "captured",
      "order_id": "order_JJCqnZG8f3754z",
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
      "contact": "+919820958250",
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

### Reverse transfers from all linked accounts
```go
data:= map[string]interface{}{
  "amount": 100,
}
body, err := client.Transfer.Reverse("<transferId>", data, nil)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| transferId*   | string      | The id of the transfer to be fetched  |
| amount   | integer      | The amount that is to be reversed from the linked account to your account. If this parameter is not passed, the entire transfer amount will be reversed. |

**Response:**
```json
{
  "id": "rvrsl_EB0BWgGDAu7tOz",
  "entity": "reversal",
  "transfer_id": "trf_EAznuJ9cDLnF7Y",
  "amount": 100,
  "fee": 0,
  "tax": 0,
  "currency": "INR",
  "notes": [],
  "initiator_id": "CJoeHMNpi0nC7k",
  "customer_refund_id": null,
  "created_at": 1580456007
}
```
-------------------------------------------------------------------------------------------------------

### Hold settlements for transfers
```go
data:= map[string]interface{}{
  "transfers": []interface{}{
    map[string]interface{}{
      "account": "acc_HjVXbtpSCIxENR",
      "amount": 100,
      "currency": "INR",
      "on_hold": true,
    },
  },
}
body, err := client.Payment.Transfer("<paymentId>", data, nil)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| paymentId*   | string      | The id of the payment to be fetched  |
| transfers   | object     | All parameters listed [here](https://razorpay.com/docs/api/route/#hold-settlements-for-transfers) are supported |

**Response:**
```json
{
  "entity": "collection",
  "count": 1,
  "items": [
    {
      "id": "trf_Jfm1KCF6w1oWgy",
      "entity": "transfer",
      "status": "pending",
      "source": "pay_JXPULbHbkkkS8D",
      "recipient": "acc_I0QRP7PpvaHhpB",
      "amount": 100,
      "currency": "INR",
      "amount_reversed": 0,
      "notes": [],
      "linked_account_notes": [],
      "on_hold": true,
      "on_hold_until": null,
      "recipient_settlement_id": null,
      "created_at": 1654860101,
      "processed_at": null,
      "error": {
        "code": null,
        "description": null,
        "reason": null,
        "field": null,
        "step": null,
        "id": "trf_Jfm1KCF6w1oWgy",
        "source": null,
        "metadata": null
      }
    }
  ]
}
```
-------------------------------------------------------------------------------------------------------

### Modify settlement hold for transfers
```go
data:= map[string]interface{}{
  "on_hold": 1,
  "on_hold_until": 1679691505,
}
body, err := client.Transfer.Edit("<transferId>", data, nil)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| transferId*   | string      | The id of the transfer for which settlement configurations should be modified.  |
| on_hold*   | boolean     | Indicates whether the account settlement for transfer is on hold. Possible values: `1` or `0` |
| on_hold_until   | integer     | Timestamp, in Unix, that indicates until when the settlement of the transfer must be put on hold. If no value is passed, the settlement is put on hold indefinitely. |

**Response:**
```json
{
    "entity": "collection",
    "count": 1,
    "items": [
        {
            "id": "trf_JhemwjNekar9Za",
            "entity": "transfer",
            "status": "pending",
            "source": "pay_I7watngocuEY4P",
            "recipient": "acc_HjVXbtpSCIxENR",
            "amount": 100,
            "currency": "INR",
            "amount_reversed": 0,
            "notes": [],
            "linked_account_notes": [],
            "on_hold": true,
            "on_hold_until": null,
            "recipient_settlement_id": null,
            "created_at": 1655271313,
            "processed_at": null,
            "error": {
                "code": null,
                "description": null,
                "reason": null,
                "field": null,
                "step": null,
                "id": "trf_JhemwjNekar9Za",
                "source": null,
                "metadata": null
            }
        }
    ]
}
```

-------------------------------------------------------------------------------------------------------

**PN: * indicates mandatory fields**
<br>
<br>
**For reference click [here](https://razorpay.com/docs/api/route/#transfers/)**
