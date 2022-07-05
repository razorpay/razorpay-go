## Paper NACH

### Create customer
```go
data := map[string]interface{}{
    "name": "Gaurav Kumar",
    "contact": 9123456780,
    "email": "gaurav.kumar@example.com",
    "fail_existing": 0,
    "notes": map[string]interface{}{
        "notes_key_1": "Tea, Earl Grey, Hot",
        "notes_key_2": "Tea, Earl Grey… decaf.",
    },
}
body, err := client.Customer.Create(data, nil)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| name*          | string      | Name of the customer                        |
| email        | string      | Email of the customer                       |
| contact      | string      | Contact number of the customer              |
| notes         | object      | A key-value pair                            |
| fail_existing | string | If a customer with the same details already exists, the request throws an exception by default. Possible value is `0` or `1`|

**Response:**
```json
{
  "id": "cust_1Aa00000000003",
  "entity": "customer",
  "name": "Gaurav Kumar",
  "email": "Gaurav.Kumar@example.com",
  "contact": "9000000000",
  "gstin": null,
  "notes": {
    "notes_key_1": "Tea, Earl Grey, Hot",
    "notes_key_2": "Tea, Earl Grey… decaf."
  },
  "created_at": 1582033731
}
```
-------------------------------------------------------------------------------------------------------

### Create order

```go
data := map[string]interface{}{
  "amount":0,
  "currency":"INR",
  "method":"nach",
  "customer_id": "cust_1Aa00000000001",
  "receipt":"Receipt No. 1",
  "notes": map[string]interface{}{
    "notes_key_1": "Beam me up Scotty",
    "notes_key_2": "Engage",
  },
  "token":map[string]interface{}{
    "auth_type":"physical",
    "max_amount":10000000,
    "expire_at":2709971120,
    "notes": map[string]interface{}{
      "notes_key_1": "Tea, Earl Grey, Hot",
      "notes_key_2": "Tea, Earl Grey… decaf.",
    },
    "bank_account":map[string]interface{}{
      "account_number":"11214311215411",
      "ifsc_code":"HDFC0000001",
      "beneficiary_name":"Gaurav Kumar",
      "account_type":"savings",
    },
    "nach":map[string]interface{}{
      "form_reference1":"Recurring Payment for Gaurav Kumar",
      "form_reference2":"Method Paper NACH",
      "description":"Paper NACH Gaurav Kumar",
    },
  },
}
body, err := client.Order.Create(data, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
| amount*   | integer      | The amount to be captured (should be equal to the authorized amount, in paise) |
| currency*   | string  | The currency of the payment (defaults to INR)  |
| customerId*   | string      | The id of the customer to be fetched |
| receipt      | string  | Your system order reference id.  |
| method*    | string  | The authorization method. In this case the value will be `nach`.  |
| token  | object  | All keys listed [here](https://razorpay.com/docs/api/payments/recurring-payments/paper-nach/create-authorization-transaction/#112-create-an-order) are supported |
| notes | object  | A key-value pair  |

**Response:**
```json
{
   "id":"order_1Aa00000000001",
   "entity":"order",
   "amount":0,
   "amount_paid":0,
   "amount_due":0,
   "currency":"INR",
   "receipt":"rcptid #10",
   "offer_id":null,
   "offers":{
      "entity":"collection",
      "count":0,
      "items":[

      ]
   },
   "status":"created",
   "attempts":0,
   "notes":{
      "notes_key_1":"Beam me up Scotty",
      "notes_key_2":"Engage"
   },
   "created_at":1579775420,
   "token":{
      "method":"nach",
      "notes":{
         "notes_key_1":"Tea, Earl Grey, Hot",
         "notes_key_2":"Tea, Earl Grey… decaf."
      },
      "recurring_status":null,
      "failure_reason":null,
      "currency":"INR",
      "max_amount":10000000,
      "auth_type":"physical",
      "expire_at":1580480689,
      "nach":{
         "create_form":true,
         "form_reference1":"Recurring Payment for Gaurav Kumar",
         "form_reference2":"Method Paper NACH",
         "prefilled_form":"https://rzp.io/i/bitw",
         "upload_form_url":"https://rzp.io/i/gts",
         "description":"Paper NACH Gaurav Kumar"
      },
      "bank_account":{
         "ifsc":"HDFC0000001",
         "bank_name":"HDFC Bank",
         "name":"Gaurav Kumar",
         "account_number":"11214311215411",
         "account_type":"savings",
         "beneficiary_email":"gaurav.kumar@example.com",
         "beneficiary_mobile":"9876543210"
      },
      "first_payment_amount":0
   }
}
```

-------------------------------------------------------------------------------------------------------

### Create an Authorization Payment

Please refer this [doc](https://razorpay.com/docs/api/recurring-payments/paper-nach/authorization-transaction/#113-create-an-authorization-payment) for authorization payment

-------------------------------------------------------------------------------------------------------

### Create registration link

```go
data:= map[string]interface{}{
  "customer":map[string]interface{}{
    "name":"Gaurav Kumar",
    "email":"gaurav.kumar@example.com",
    "contact":"9123456780",
  },
  "amount":0,
  "currency":"INR",
  "type":"link",
  "description":"12 p.m. Meals",
  "subscription_registration":map[string]interface{}{
    "method":"nach",
    "auth_type":"physical",
    "bank_account":map[string]interface{}{
      "beneficiary_name":"Gaurav Kumar",
      "account_number":"11214311215411",
      "account_type":"savings",
      "ifsc_code":"HDFC0001233",
    },
    "nach":map[string]interface{}{
      "form_reference1":"Recurring Payment for Gaurav Kumar",
      "form_reference2":"Method Paper NACH",
    },
    "expire_at":1947483647,
    "max_amount":50000,
  },
  "receipt":"Receipt No. 1",
  "sms_notify":1,
  "email_notify":1,
  "expire_by":1647483647,
  "notes":map[string]interface{}{
    "note_key 1":"Beam me up Scotty",
    "note_key 2":"Tea. Earl Gray. Hot.",
  },
}

body, err := client.Invoice.CreateRegistrationLink(data, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
| customer         | object  | All parameters listed [here](https://razorpay.com/docs/api/payments/recurring-payments/paper-nach/create-authorization-transaction/#121-create-a-registration-link) |
| type*        | string  | In this case, the value is `link`.                      |
| currency*        | string  | The 3-letter ISO currency code for the payment. Currently, only `INR` is supported. |
| amount*         | integer  | The payment amount in the smallest currency sub-unit.                 |
| description*    | string  | A description that appears on the hosted page. For example, `12:30 p.m. Thali meals (Gaurav Kumar`).                                                             |
| subscription_registration | object  | All parameters listed [here](https://razorpay.com/docs/api/payments/recurring-payments/paper-nach/create-authorization-transaction/#121-create-a-registration-link) |
| sms_notify  | boolean  | SMS notifications are to be sent by Razorpay (default : 1)  |
| email_notify | boolean  | Email notifications are to be sent by Razorpay (default : 1)  |
| expire_by    | integer | The timestamp, in Unix format, till when the customer can make the authorization payment. |
| receipt      | string  | Your system order reference id.  |
| notes           | object  | A key-value pair  |

**Response:**
```json
{
    "id": "inv_FHrZiAubEzDdaq",
    "entity": "invoice",
    "receipt": "Receipt No. 27",
    "invoice_number": "Receipt No. 27",
    "customer_id": "cust_BMB3EwbqnqZ2EI",
    "customer_details": {
        "id": "cust_BMB3EwbqnqZ2EI",
        "name": "Gaurav Kumar",
        "email": "gaurav.kumar@example.com",
        "contact": "9123456780",
        "gstin": null,
        "billing_address": null,
        "shipping_address": null,
        "customer_name": "Gaurav Kumar",
        "customer_email": "gaurav.kumar@example.com",
        "customer_contact": "9123456780"
    },
    "order_id": "order_FHrZiBOkWHZPOp",
    "line_items": [],
    "payment_id": null,
    "status": "issued",
    "expire_by": 1647483647,
    "issued_at": 1595491154,
    "paid_at": null,
    "cancelled_at": null,
    "expired_at": null,
    "sms_status": "sent",
    "email_status": "sent",
    "date": 1595491154,
    "terms": null,
    "partial_payment": false,
    "gross_amount": 0,
    "tax_amount": 0,
    "taxable_amount": 0,
    "amount": 0,
    "amount_paid": 0,
    "amount_due": 0,
    "currency": "INR",
    "currency_symbol": "₹",
    "description": "12 p.m. Meals",
    "notes": {
        "note_key 1": "Beam me up Scotty",
        "note_key 2": "Tea. Earl Gray. Hot."
    },
    "comment": null,
    "short_url": "https://rzp.io/i/bzDYbNg",
    "view_less": true,
    "billing_start": null,
    "billing_end": null,
    "type": "link",
    "group_taxes_discounts": false,
    "created_at": 1595491154,
    "idempotency_key": null,
    "token": {
        "method": "nach",
        "notes": {
            "note_key 1": "Beam me up Scotty",
            "note_key 2": "Tea. Earl Gray. Hot."
        },
        "recurring_status": null,
        "failure_reason": null,
        "currency": "INR",
        "max_amount": 50000,
        "auth_type": "physical",
        "expire_at": 1947483647,
        "nach": {
            "create_form": true,
            "form_reference1": "Recurring Payment for Gaurav Kumar",
            "form_reference2": "Method Paper NACH",
            "prefilled_form": "https://rzp.io/i/exdIzYN",
            "upload_form_url": "https://rzp.io/i/bzDYbNg",
            "description": "12 p.m. Meals"
        },
        "bank_account": {
            "ifsc": "HDFC0001233",
            "bank_name": "HDFC Bank",
            "name": "Gaurav Kumar",
            "account_number": "11214311215411",
            "account_type": "savings",
            "beneficiary_email": "gaurav.kumar@example.com",
            "beneficiary_mobile": "9123456780"
        },
        "first_payment_amount": 0
    },
    "nach_form_url": "https://rzp.io/i/exdIzYN"
}
```
-------------------------------------------------------------------------------------------------------

### Send/Resend notifications

```go
body, err := client.Invoice.Notify("<invoiceId>", "<medium>", nil, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
| invoiceId*          | string | The id of the invoice to be notified                         |
| medium*          | string | `sms`/`email`, Medium through which notification should be sent.                         |

**Response:**
```json
{
    "success": true
}
```
-------------------------------------------------------------------------------------------------------

### Cancel a registration link

```go
body, err := client.Invoice.Cancel("<invoiceId>", nil, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
| invoiceId*          | string | The id of the invoice to be cancelled                         |

**Response:**
```json
{
  "id": "inv_FHrZiAubEzDdaq",
  "entity": "invoice",
  "receipt": "Receipt No. 27",
  "invoice_number": "Receipt No. 27",
  "customer_id": "cust_BMB3EwbqnqZ2EI",
  "customer_details": {
    "id": "cust_BMB3EwbqnqZ2EI",
    "name": "Gaurav Kumar",
    "email": "gaurav.kumar@example.com",
    "contact": "9123456780",
    "gstin": null,
    "billing_address": null,
    "shipping_address": null,
    "customer_name": "Gaurav Kumar",
    "customer_email": "gaurav.kumar@example.com",
    "customer_contact": "9123456780"
  },
  "order_id": "order_FHrZiBOkWHZPOp",
  "line_items": [],
  "payment_id": null,
  "status": "cancelled",
  "expire_by": 1647483647,
  "issued_at": 1595491154,
  "paid_at": null,
  "cancelled_at": 1595491339,
  "expired_at": null,
  "sms_status": "sent",
  "email_status": "sent",
  "date": 1595491154,
  "terms": null,
  "partial_payment": false,
  "gross_amount": 0,
  "tax_amount": 0,
  "taxable_amount": 0,
  "amount": 0,
  "amount_paid": 0,
  "amount_due": 0,
  "currency": "INR",
  "currency_symbol": "₹",
  "description": "12 p.m. Meals",
  "notes": {
    "note_key 1": "Beam me up Scotty",
    "note_key 2": "Tea. Earl Gray. Hot."
  },
  "comment": null,
  "short_url": "https://rzp.io/i/bzDYbNg",
  "view_less": true,
  "billing_start": null,
  "billing_end": null,
  "type": "link",
  "group_taxes_discounts": false,
  "created_at": 1595491154,
  "idempotency_key": null,
  "token": {
    "method": "nach",
    "notes": {
      "note_key 1": "Beam me up Scotty",
      "note_key 2": "Tea. Earl Gray. Hot."
    },
    "recurring_status": null,
    "failure_reason": null,
    "currency": "INR",
    "max_amount": 50000,
    "auth_type": "physical",
    "expire_at": 1947483647,
    "nach": {
      "create_form": true,
      "form_reference1": "Recurring Payment for Gaurav Kumar",
      "form_reference2": "Method Paper NACH",
      "prefilled_form": "https://rzp.io/i/tSYd5aV",
      "upload_form_url": "https://rzp.io/i/bzDYbNg",
      "description": "12 p.m. Meals"
    },
    "bank_account": {
      "ifsc": "HDFC0001233",
      "bank_name": "HDFC Bank",
      "name": "Gaurav Kumar",
      "account_number": "11214311215411",
      "account_type": "savings",
      "beneficiary_email": "gaurav.kumar@example.com",
      "beneficiary_mobile": "9123456780"
    },
    "first_payment_amount": 0
  },
  "nach_form_url": "https://rzp.io/i/tSYd5aV"
}
```
-------------------------------------------------------------------------------------------------------

### Fetch Payment ID using Order ID

```go
body, err := client.Order.Payments("<orderId>", nil, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
| orderId*          | string | Order id for which payment id need to be fetched                         |

**Response:**
```json
{
  "entity":"collection",
  "count":1,
  "items":[
    {
      "id":"pay_1Aa00000000003",
      "entity":"payment",
      "amount":0,
      "currency":"INR",
      "status":"captured",
      "order_id":"order_1Aa00000000003",
      "invoice_id":"inv_1Aa00000000003",
      "international":false,
      "method":"nach",
      "amount_refunded":0,
      "refund_status":null,
      "captured":true,
      "description":"12 p.m. Meals",
      "card_id":null,
      "bank":"HDFC",
      "wallet":null,
      "vpa":null,
      "email":"gaurav.kumar@example.com",
      "contact":"99876543210",
      "customer_id":"cust_1Aa00000000002",
      "token_id":"token_1Aa00000000003",
      "notes":{
        "note_key 1":"Beam me up Scotty",
        "note_key 2":"Tea. Earl Gray. Hot."
      },
      "fee":0,
      "tax":0,
      "error_code":null,
      "error_description":null,
      "created_at":1580109147
    }
  ]
}
```
-------------------------------------------------------------------------------------------------------

### Fetch token by payment ID

```go
body, err := client.Payment.Fetch("<paymentId>", nil, nil)
```

**Parameters:**

| Name       | Type   | Description                       |
|------------|--------|-----------------------------------|
| paymentId* | string | Id of the payment to be retrieved |

**Response:**
```json
{
  "id": "pay_EnLNTjINiPkMEZ",
  "entity": "payment",
  "amount": 0,
  "currency": "INR",
  "status": "captured",
  "order_id": "order_EnLLfglmKksr4K",
  "invoice_id": "inv_EnLLfgCzRfcMuh",
  "international": false,
  "method": "nach",
  "amount_refunded": 0,
  "refund_status": null,
  "captured": true,
  "description": "Invoice #inv_EnLLfgCzRfcMuh",
  "card_id": null,
  "bank": "UTIB",
  "wallet": null,
  "vpa": null,
  "email": "gaurav.kumar@example.com",
  "contact": "+919876543210",
  "customer_id": "cust_DtHaBuooGHTuyZ",
  "token_id": "token_EnLNTnn7uyRg5V",
  "notes": {
    "note_key 1": "Beam me up Scotty",
    "note_key 2": "Tea. Earl Gray. Hot."
  },
  "fee": 0,
  "tax": 0,
  "error_code": null,
  "error_description": null,
  "error_source": null,
  "error_step": null,
  "error_reason": null,
  "acquirer_data": {},
  "created_at": 1588827564
}
```
-------------------------------------------------------------------------------------------------------

### Fetch tokens by customer ID

```go
body, err := client.Token.All("<customerId>", nil, nil)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| customerId*          | string      | The id of the customer for whom tokens are to be retrieved. |

**Response:**
```json
{
  "entity": "collection",
  "count": 1,
  "items": [
    {
      "id": "token_EhYgIE3pOyMQpD",
      "entity": "token",
      "token": "3mQ5Czc6APNppI",
      "bank": "HDFC",
      "wallet": null,
      "method": "nach",
      "vpa": null,
      "recurring": true,
      "recurring_details": {
        "status": "confirmed",
        "failure_reason": null
      },
      "auth_type": "physical",
      "mrn": null,
      "used_at": 1587564373,
      "created_at": 1587564373,
      "dcc_enabled": false
    }
  ]
}
```
-------------------------------------------------------------------------------------------------------

### Delete token

```go
body, err := client.Token.Delete("<customerId>", "<tokenId>", nil, nil)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| customerId*   | string      | The id of the customer with whom the token is linked |
| tokenId*   | string      | The id of the token to be deleted |

**Response:**
```json
{
    "deleted": true
}
```
-------------------------------------------------------------------------------------------------------

### Create an order to charge the customer

```go
data:= map[string]interface{}{
  "amount":1000,
  "currency":"INR",
  "payment_capture": true,
  "receipt":"Receipt No. 1",
  "notes": map[string]interface{}{
    "notes_key_1":"Tea, Earl Grey, Hot",
    "notes_key_2":"Tea, Earl Grey… decaf.",
  },
}
body, err := client.Order.Create(data, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
| amount*          | integer | Amount in currency subunits. For cards, the minimum value is `100` (₹1).                                              |
| currency*        | string  | Currency of the order. Currently only `INR` is supported.                      |
| payment_capture*  | boolean  | Indicates whether payment status should be changed to captured automatically or not. Possible values: true - Payments are captured automatically. false - Payments are not captured automatically. |
| receipt         | string  | Your system order reference id.                                              |
| notes           | object  | A key-value pair                                                             |
**Response:**
```json
{
   "id":"order_1Aa00000000002",
   "entity":"order",
   "amount":1000,
   "amount_paid":0,
   "amount_due":1000,
   "currency":"INR",
   "receipt":"Receipt No. 1",
   "offer_id":null,
   "status":"created",
   "attempts":0,
   "notes":{
      "notes_key_1":"Tea, Earl Grey, Hot",
      "notes_key_2":"Tea, Earl Grey… decaf."
   },
   "created_at":1579782776
}
```
-------------------------------------------------------------------------------------------------------

### Create a Recurring Payment

```go
data:= map[string]interface{}{
  "email": "gaurav.kumar@example.com",
  "contact": "9123456789",
  "amount": 1000,
  "currency": "INR",
  "order_id": "<orderId>",
  "customer_id": "<customerId>",
  "token": "<tokenId>",
  "recurring": "1",
  "description": "Creating recurring payment for Gaurav Kumar",
  "notes": map[string]interface{}{
    "note_key 1": "Beam me up Scotty",
    "note_key 2": "Tea. Earl Gray. Hot.",
  },
}
body, err := Client.Payment.CreateRecurringPayment(data, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
| email*          | string | The customer's email address.                                               |
| contact*        | string  | The customer's phone number.                      |
| amount*         | integer  | The amount you want to charge your customer. This should be the same as the amount in the order.                        |
| currency*        | string  | The 3-letter ISO currency code for the payment. Currently, only `INR` is supported. |
| order_id*        | string  | The unique identifier of the order created. |
| customer_id*        | string  | The `customer_id` for the customer you want to charge.  |
| token*        | string  | The `token_id` generated when the customer successfully completes the authorization payment. Different payment instruments for the same customer have different `token_id`.|
| recurring*        | string  | Determines if recurring payment is enabled or not. Possible values:<br>* `1` - Recurring is enabled.* `0` - Recurring is not enabled.|
| description        | string  | A user-entered description for the payment.|
| notes       | object  | Key-value pair that can be used to store additional information about the entity. Maximum 15 key-value pairs, 256 characters (maximum) each. |

**Response:**
```json
{
  "razorpay_payment_id" : "pay_1Aa00000000001",
  "razorpay_order_id" : "order_1Aa00000000001",
  "razorpay_signature" : "9ef4dffbfd84f1318f6739a3ce19f9d85851857ae648f114332d8401e0949a3d"
}
```
-------------------------------------------------------------------------------------------------------


**PN: * indicates mandatory fields**
<br>
<br>
**For reference click [here](https://razorpay.com/docs/api/recurring-payments/paper-nach/authorization-transaction/)**