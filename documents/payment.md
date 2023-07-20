## Payments

### Capture payment

```go
paymentId := "pay_JIskP2THPhAUXg"
amount := 40

data := map[string]interface{}{
  "currency": "INR",
}

body, err := client.Payment.Capture(paymentId, amount, data, nil)
```

**Parameters:**

| Name      | Type    | Description                                                                    |
|-----------|---------|--------------------------------------------------------------------------------|
| paymentId* | string  | Id of the payment to capture                                                   |
| amount*    | integer | The amount to be captured (should be equal to the authorized amount, in paise) |
| currency*   | string  | The currency of the payment (defaults to INR)                                  |

**Response:**
```json
{
  "id": "pay_G8VQzjPLoAvm6D",
  "entity": "payment",
  "amount": 1000,
  "currency": "INR",
  "status": "captured",
  "order_id": "order_G8VPOayFxWEU28",
  "invoice_id": null,
  "international": false,
  "method": "upi",
  "amount_refunded": 0,
  "refund_status": null,
  "captured": true,
  "description": "Purchase Shoes",
  "card_id": null,
  "bank": null,
  "wallet": null,
  "vpa": "gaurav.kumar@exampleupi",
  "email": "gaurav.kumar@example.com",
  "contact": "+919999999999",
  "customer_id": "cust_DitrYCFtCIokBO",
  "notes": [],
  "fee": 24,
  "tax": 4,
  "error_code": null,
  "error_description": null,
  "error_source": null,
  "error_step": null,
  "error_reason": null,
  "acquirer_data": {
    "rrn": "033814379298"
  },
  "created_at": 1606985209
}
```
-------------------------------------------------------------------------------------------------------

### Fetch all payments

```go
data := map[string]interface{}{
  "count": 1,
}

body, err := client.Payment.All(data, nil)
```

**Parameters:**

| Name  | Type      | Description                                      |
|-------|-----------|--------------------------------------------------|
| from  | timestamp | timestamp after which the payments were created  |
| to    | timestamp | timestamp before which the payments were created |
| count | integer   | number of payments to fetch (default: 10)        |
| skip  | integer   | number of payments to be skipped (default: 0)    |
| expand[]   | string    |  Used to retrieve additional information about the payment. Possible value is `cards` or `emi`|

**Response:**
```json
{
  "entity": "collection",
  "count": 2,
  "items": [
    {
      "id": "pay_G8VaL2Z68LRtDs",
      "entity": "payment",
      "amount": 900,
      "currency": "INR",
      "status": "captured",
      "order_id": "order_G8VXfKDWDEOHHd",
      "invoice_id": null,
      "international": false,
      "method": "netbanking",
      "amount_refunded": 0,
      "refund_status": null,
      "captured": true,
      "description": "Purchase Shoes",
      "card_id": null,
      "bank": "KKBK",
      "wallet": null,
      "vpa": null,
      "email": "gaurav.kumar@example.com",
      "contact": "+919999999999",
      "customer_id": "cust_DitrYCFtCIokBO",
      "notes": [],
      "fee": 22,
      "tax": 4,
      "error_code": null,
      "error_description": null,
      "error_source": null,
      "error_step": null,
      "error_reason": null,
      "acquirer_data": {
        "bank_transaction_id": "0125836177"
      },
      "created_at": 1606985740
    }
  ]
}
```
-------------------------------------------------------------------------------------------------------

### Fetch a payment

```go
paymentId := "pay_G8VQzjPLoAvm6D"

body, err := client.Payment.Fetch(paymentId, nil, nil)
```

**Parameters:**

| Name       | Type   | Description                       |
|------------|--------|-----------------------------------|
| paymentId* | string | Id of the payment to be retrieved |
| expand[]   | string    |  Used to retrieve additional information about the payment. Possible value is `cards` or `emi`|

**Response:**
```json
{
  "id": "pay_G8VQzjPLoAvm6D",
  "entity": "payment",
  "amount": 1000,
  "currency": "INR",
  "status": "captured",
  "order_id": "order_G8VPOayFxWEU28",
  "invoice_id": null,
  "international": false,
  "method": "upi",
  "amount_refunded": 0,
  "refund_status": null,
  "captured": true,
  "description": "Purchase Shoes",
  "card_id": null,
  "bank": null,
  "wallet": null,
  "vpa": "gaurav.kumar@exampleupi",
  "email": "gaurav.kumar@example.com",
  "contact": "+919999999999",
  "customer_id": "cust_DitrYCFtCIokBO",
  "notes": [],
  "fee": 24,
  "tax": 4,
  "error_code": null,
  "error_description": null,
  "error_source": null,
  "error_step": null,
  "error_reason": null,
  "acquirer_data": {
    "rrn": "033814379298"
  },
  "created_at": 1606985209
}
```
-------------------------------------------------------------------------------------------------------

### Fetch payments for an order

```go
orderId := "order_DovFx48wjYEr2I"

body, err := client.Order.Payments(orderId, nil, nil)
```
**Parameters**

| Name     | Type   | Description                         |
|----------|--------|-------------------------------------|
| orderId* | string | The id of the order to be retrieve payment info |

**Response:**
```json
{
  "count": 1,
  "entity": "collection",
  "items": [
    {
      "id": "pay_DovGQXOkPBJjjU",
      "entity": "payment",
      "amount": 600,
      "currency": "INR",
      "status": "captured",
      "order_id": "order_DovFx48wjYEr2I",
      "method": "netbanking",
      "amount_refunded": 0,
      "refund_status": null,
      "captured": true,
      "description": "A Wild Sheep Chase is a novel by Japanese author Haruki Murakami",
      "card_id": null,
      "bank": "SBIN",
      "wallet": null,
      "vpa": null,
      "email": "gaurav.kumar@example.com",
      "contact": "9364591752",
      "fee": 70,
      "tax": 10,
      "error_code": null,
      "error_description": null,
      "error_source": null,
      "error_step": null,
      "error_reason": null,
      "notes": [],
      "acquirer_data": {
        "bank_transaction_id": "0125836177"
      },
      "created_at": 1400826750
    }
  ]
}
```
-------------------------------------------------------------------------------------------------------

### Update a payment

```go
paymentId := "pay_CBYy6tLmJTzn3Q"

data := map[string]interface{}{
        "notes": map[string]interface{}{
            "key1": "value1",
            "key2": "value2",
        },
    }

body, err := client.Payment.Edit(paymentId, data, nil)
```

**Parameters:**

| Name        | Type    | Description                          |
|-------------|---------|--------------------------------------|
| paymentId* | string  | Id of the payment to update          |
| notes*       | object  | A key-value pair                     |

**Response:**
```json
{
    "id": "pay_JpFbdMHunrN6LJ",
    "entity": "payment",
    "amount": 80000,
    "currency": "INR",
    "status": "captured",
    "order_id": "order_JcbBWAhl9z0qJx",
    "invoice_id": "inv_JcbBW8MrnNhuiC",
    "international": false,
    "method": "card",
    "amount_refunded": 0,
    "refund_status": null,
    "captured": true,
    "description": "Start Subscription",
    "card_id": "card_JpFbdquXOLR4rg",
    "bank": null,
    "wallet": null,
    "vpa": null,
    "email": "you@example.com",
    "contact": "+917000569565",
    "customer_id": "cust_I3FToKbnExwDLu",
    "token_id": "token_JPz75tHtcA0Yu4",
    "notes": {
        "key1": "value1",
        "key2": "value2"
    },
    "fee": 2000,
    "tax": 0,
    "error_code": null,
    "error_description": null,
    "error_source": null,
    "error_step": null,
    "error_reason": null,
    "acquirer_data": {
        "auth_code": "159838"
    },
    "created_at": 1656929352
}
```
-------------------------------------------------------------------------------------------------------

### Fetch expanded card or emi details for payments

Request #1: Card

```go
data := map[string]interface{}{
     "expand[]":"card",
    }

body, err := client.Payment.All(data, nil)
```

Request #2: EMI

```go
data := map[string]interface{}{
      "expand[]":"emi",
    }

body, err := client.Payment.All(data, nil)
```

**Response:**<br>
For expanded card or emi details for payments response please click [here](https://razorpay.com/docs/api/payments/#fetch-expanded-card-or-emi-details-for-payments)

-------------------------------------------------------------------------------------------------------

### Fetch card details with paymentId

```go
paymentId := "pay_CBYy6tLmJTzn3Q"

body, err := client.Payment.FetchCardDetails(paymentId, nil, nil)
```

**Parameters:**

| Name        | Type    | Description                          |
|-------------|---------|--------------------------------------|
| paymentId* | string  | Id of the payment to update          |

**Response:**
```json
{
  "id": "card_JXPULjlKqC5j0i",
  "entity": "card",
  "name": "ankit",
  "last4": "4366",
  "network": "Visa",
  "type": "credit",
  "issuer": "UTIB",
  "international": false,
  "emi": true,
  "sub_type": "consumer",
  "token_iin": null
}
```
-------------------------------------------------------------------------------------------------------

### Fetch Payment Downtime Details

```go
body, err := client.Payment.FetchPaymentDowntime(nil, nil)
```
**Response:** <br>
For payment downtime response please click [here](https://razorpay.com/docs/api/payments/downtime/#fetch-payment-downtime-details)

-------------------------------------------------------------------------------------------------------

### Fetch Payment Downtime

```go
downtimeId := "down_F7LroRQAAFuswd"

body, err := client.Payment.FetchPaymentDowntimeById(downtimeId, nil)
```

**Parameters:**

| Name        | Type    | Description                          |
|-------------|---------|--------------------------------------|
| downtimeId* | string  | Id to fetch payment downtime         |

**Response:**
For payment downtime by id response please click [here](https://razorpay.com/docs/api/payments/downtime/#fetch-payment-downtime-details-by-id)

-------------------------------------------------------------------------------------------------------

### Payment capture settings API

```go
data := map[string]interface{}{
  "amount": 50000,
  "currency": "INR",
  "receipt": "some_receipt_id",
  "payment_capture": 1,
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
| payment_capture  | integer  | 1 if capture should be done automatically or else 0                                                             |
| notes           | object  | A key-value pair                                                             |

**Response:** <br>
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

### Create Payment Json

```go
para_attr := map[string]interface{}{
  "amount": 100,
  "currency": "INR",
  "order_id": "order_EAkbvXiCJlwhHR",
  "email": "gaurav.kumar@example.com",
  "contact": 9090909090,
  "method": "card",
  "card": map[string]interface{}{
    "number": "4111111111111111",
    "name": "Gaurav",
    "expiry_month": 11,
    "expiry_year": 23,
    "cvv": 100,
  },
}
body, err := client.Payment.CreatePaymentJson(para_attr, nil)
```

**Parameters:**
 please refer this [doc](https://razorpay.com/docs/payment-gateway/s2s-integration/payment-methods/) for params

**Response:** <br>
```json
{
  "razorpay_payment_id": "pay_FVmAstJWfsD3SO",
  "next": [
    {
      "action": "redirect",
      "url": "https://api.razorpay.com/v1/payments/FVmAtLUe9XZSGM/authorize"
    },
    {
      "action": "otp_generate",
      "url": "https://api.razorpay.com/v1/payments/pay_FVmAstJWfsD3SO/otp_generate?track_id=FVmAtLUe9XZSGM&key_id=<YOUR_KEY_ID>"
    }
  ]
}
```
-------------------------------------------------------------------------------------------------------

### Create Payment Json (Third party validation)

```go
para_attr := map[string]interface{}{
  "amount": "500",
  "currency": "INR",
  "email": "gaurav.kumar@example.com",
  "contact": "9123456789",
  "order_id": "order_GAWN9beXgaqRyO",
  "method": "netbanking",
  "bank": "HDFC",
}
body, err := client.Payment.CreatePaymentJson(para_attr, nil)
```

**Parameters:**
| Name        | Type    | Description                          |
|-------------|---------|--------------------------------------|
| amount*          | integer | Amount of the order to be paid  |
| currency*   | string  | The currency of the payment (defaults to INR)                                  |
| order_id*        | string  | The unique identifier of the order created. |
| email*        | string      | Email of the customer                       |
| contact*      | string      | Contact number of the customer              |
| method*      | string  | Possible value is `netbanking` |
| bank*      | string      | The customer's bank code.For example, `HDFC`.|

 please refer this [doc](https://razorpay.com/docs/payments/third-party-validation/s2s-integration/netbanking#step-3-create-a-payment) for params

**Response:** <br>
```json
{
  "razorpay_payment_id": "pay_GAWOYqPlvrtPSi",
  "next": [
    {
      "action": "redirect",
      "url": "https://api.razorpay.com/v1/payments/pay_GAWOYqPlvrtPSi/authorize"
    }
  ]
}
```
-------------------------------------------------------------------------------------------------------
### Create Payment UPI Using VPA token (Third party validation)

```go
para_attr := map[string]interface{}{
  "amount": 200,
  "currency": "INR",
  "order_id": "order_GAWRjlWkVcRh0V",
  "email": "gaurav.kumar@example.com",
  "contact": "9123456789",
  "method": "upi",
  "customer_id": "cust_EIW4T2etiweBmG",
  "save": 1,
  "ip": "192.168.0.103",
  "referer": "http",
  "user_agent": "Mozilla/5.0",
  "description": "Test flow",
  "notes": map[string]interface{}{
    "note_key": "value1",
  },
  "upi": map[string]interface{}{
    "flow": "collect",
    "vpa": "gauravkumar@exampleupi",
    "expiry_time": 5,
  },
}
body, err := client.Payment.CreateUpi(para_attr, nil)
```

**Parameters:**
| Name        | Type    | Description                          |
|-------------|---------|--------------------------------------|
| amount*          | integer | Amount of the order to be paid  |
| currency*   | string  | The currency of the payment (defaults to INR)                                  |
| order_id*        | string  | The unique identifier of the order created. |
| email*        | string      | Email of the customer                       |
| customer_id*   | string      | The id of the customer to be fetched |
| contact*      | string      | Contact number of the customer              |
| notes | object  | A key-value pair  |
| description | string  | Descriptive text of the payment. |
| save | boolean  |  Specifies if the VPA should be stored as tokens.Possible value is `0`, `1`  |
| callback_url   | string      | URL where Razorpay will submit the final payment status. |
| ip*   | string      | The client's browser IP address. For example `117.217.74.98` |
| referer*   | string      | Value of `referer` header passed by the client's browser. For example, `https://example.com/` |
| user_agent*   | string      | Value of `user_agent` header passed by the client's browser. For example, `Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36` |
| upi* (for Upi only)  | object      | All keys listed [here](https://razorpay.com/docs/payments/third-party-validation/s2s-integration/upi/collect#step-14-initiate-a-payment) are supported  |

**Response:** <br>
```json
{
  "razorpay_payment_id": "pay_EAm09NKReXi2e0"
}
```
-------------------------------------------------------------------------------------------------------

### Create Payment UPI s2s (Third party validation)

```go
para_attr := map[string]interface{}{
    "amount": 100,
    "currency": "INR",
    "order_id": "order_Ee0biRtLOqzRjP",
    "email": "gaurav.kumar@example.com",
    "contact": "9090909090",
    "method": "upi",
    "ip": "192.168.0.103",
    "referer": "http",
    "user_agent": "Mozilla/5.0",
    "description": "Test flow",
    "notes": map[string]interface{}{
      "purpose": "UPI test payment",
    },
    "upi": map[string]interface{}{
      "flow" : "intent",
    },
}
body, err := client.Payment.CreateUpi(para_attr, nil)
```

**Parameters:**
| Name        | Type    | Description                          |
|-------------|---------|--------------------------------------|
| amount*          | integer | Amount of the order to be paid  |
| currency*   | string  | The currency of the payment (defaults to INR)                                  |
| order_id*        | string  | The unique identifier of the order created. |
| email*        | string      | Email of the customer                       |
| customer_id*   | string      | The id of the customer to be fetched |
| contact*      | string      | Contact number of the customer              |
| notes | object  | A key-value pair  |
| description | string  | Descriptive text of the payment. |
| save | boolean  |  Specifies if the VPA should be stored as tokens.Possible value is `0`, `1`  |
| callback_url   | string      | URL where Razorpay will submit the final payment status. |
| ip*   | string      | The client's browser IP address. For example `117.217.74.98` |
| referer*   | string      | Value of `referer` header passed by the client's browser. For example, `https://example.com/` |
| user_agent*   | string      | Value of `user_agent` header passed by the client's browser. For example, `Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36` |
| upi* (for Upi only)  | object      | All keys listed [here](https://razorpay.com/docs/payments/third-party-validation/s2s-integration/upi/intent/#step-2-initiate-a-payment) are supported  |

**Response:** <br>
```json
{
  "razorpay_payment_id": "pay_CMeM6XvOPGFiF",
  "link": "upi://pay?pa=success@razorpay&pn=xyz&tr=xxxxxxxxxxx&tn=gourav&am=1&cu=INR&mc=xyzw"
}
```
-------------------------------------------------------------------------------------------------------

### OTP Generate

```go
body, err := client.Payment.OtpGenerate(paymentId, nil, nil)
```

**Parameters:**

| Name        | Type    | Description                          |
|-------------|---------|--------------------------------------|
| paymentId*    | integer | Unique identifier of the payment                                               |

Doc reference [doc](https://razorpay.com/docs/payments/payment-gateway/s2s-integration/json/v2/build-integration/cards/#otp-generation)

**Response:** <br>

```json
{
 "razorpay_payment_id": "pay_FVmAstJWfsD3SO",
 "next": [
  {
   "action": "otp_submit",
   "url": "https://api.razorpay.com/v1/payments/pay_FVmAstJWfsD3SO/otp_submit/ac2d415a8be7595de09a24b41661729fd9028fdc?key_id=<YOUR_KEY_ID>"
  },
  {
   "action": "otp_resend",
   "url": "https://api.razorpay.com/v1/payments/pay_FVmAstJWfsD3SO/otp_resend/json?key_id=<YOUR_KEY_ID>"
  }
 ],
 "metadata": {
  "issuer": "HDFC",
  "network": "MC",
  "last4": "1111",
  "iin": "411111"
 }
}
```
-------------------------------------------------------------------------------------------------------

### OTP Submit

```go
data := map[string]interface{}{
	"otp" : "123456",
}
body, err := client.Payment.OtpSubmit(paymentId, data, nil)
```

**Parameters:**

| Name        | Type    | Description                          |
|-------------|---------|--------------------------------------|
| paymentId*    | integer | Unique identifier of the payment                                               |
| data.otp*    | object | The customer receives the OTP using their preferred notification medium - SMS or email |

Doc reference [doc](https://razorpay.com/docs/payments/payment-gateway/s2s-integration/json/v2/build-integration/cards/#response-on-submitting-otp)

**Response:** <br>
Success
```json
{
 "razorpay_payment_id": "pay_D5jmY2H6vC7Cy3",
 "razorpay_order_id": "order_9A33XWu170gUtm",
 "razorpay_signature": "9ef4dffbfd84f1318f6739a3ce19f9d85851857ae648f114332d8401e0949a3d"
}
```
-------------------------------------------------------------------------------------------------------

### Valid VPA (Third party validation)

```go
para_attr := map[string]interface{}{
  "vpa": "gauravkumar@exampleupi",
}

body, err := client.Payment.ValidateVpa(para_attr, nil)
```

**Parameters:**
| Name        | Type    | Description                          |
|-------------|---------|--------------------------------------|
| vpa*          | string | The virtual payment address (VPA) you want to validate. For example,   `gauravkumar@exampleupi`  |

 please refer this [doc](https://razorpay.com/docs/payments/third-party-validation/s2s-integration/upi/collect#step-13-validate-the-vpa) for params

**Response:** <br>
```json
{
  "vpa": "gauravkumar@exampleupi",
  "success": true,
  "customer_name": "Gaurav Kumar"
}
```
-------------------------------------------------------------------------------------------------------

### Fetch payment methods (Third party validation)

```go
client := razorpay.NewClient("key", "")  // Use Only razorpay key

body, err := client.Payment.FetchMethods(nil, nil)
```

**Response:** <br>
 please refer this [doc](https://razorpay.com/docs/payments/third-party-validation/s2s-integration/methods-api/#fetch-payment-methods) for response

-------------------------------------------------------------------------------------------------------
### OTP Resend

```go
body, err := client.Payment.OtpResend(paymentId, nil, nil)
```

**Parameters:**

| Name        | Type    | Description                          |
|-------------|---------|--------------------------------------|
| paymentId*    | integer | Unique identifier of the payment                                               |

Doc reference [doc](https://razorpay.com/docs/payments/payment-methods/cards/authentication/native-otp/#otp-resend)

**Response:** <br>

```json
{
  "next": [
    "otp_submit",
    "otp_resend"
  ],
  "razorpay_payment_id": "pay_JWaNvYmrx75sXo"
}
```
-------------------------------------------------------------------------------------------------------
```go
tokenIin := "412345"

body, err := client.Iin.Fetch(tokenIin, nil, nil)
```

**Parameters:**

| Name       | Type   | Description                       |
|------------|--------|-----------------------------------|
| tokenIin* | string | The token IIN. |

**Response:**
```json
{
  "iin": "412345",
  "entity": "iin",
  "network": "Visa",
  "type": "credit",
  "sub_type": "business",
  "issuer_code": "HDFC",
  "issuer_name": "HDFC Bank Ltd",
  "international": false,
  "is_tokenized": true,
  "card_iin": "411111",
  "emi":{
     "available": true
     },
  "recurring": {
     "available": true
     },
  "authentication_types": [
   {
       "type":"3ds"
   },
   {
       "type":"otp"
   }
  ]
}
```
-------------------------------------------------------------------------------------------------------
**PN: * indicates mandatory fields**
<br>
<br>
**For reference click [here](https://razorpay.com/docs/api/payments/)**