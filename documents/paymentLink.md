## Payment Links

### Create payment link

Request #1
Standard Payment Link

```go
data := map[string]interface{}{
    "amount": 1000,
    "currency": "INR",
    "accept_partial": true,
    "first_min_partial_amount": 100,
    "expire_by": 1691097057,
    "reference_id": "TS1989",
    "description": "Payment for policy no #23456",
    "customer": map[string]interface{}{
        "name": "Gaurav Kumar",
        "contact": "+919999999999",
        "email": "gaurav.kumar@example.com",
    },
    "notify": map[string]interface{}{
        "sms": true,
        "email": true,
    },
    "reminder_enable": true,
    "notes": map[string]interface{}{
        "policy_name": "Jeevan Bima",
    },
    "callback_url": "https://example-callback-url.com/",
    "callback_method": "get",
}
body, err := client.PaymentLink.Create(data, nil)
```

Request #2
UPI Payment Link

```go
data := map[string]interface{}{
    "upi_link": "true",
    "amount": 1000,
    "currency": "INR",
    "accept_partial": false,
    "expire_by": 1691097057,
    "reference_id": "TS1989",
    "description": "Payment for policy no #23456",
    "customer": map[string]interface{}{
        "name": "Gaurav Kumar",
        "contact": "+919999999999",
        "email": "mudududla.sony@razorpay.com",
    },
    "notify": map[string]interface{}{
        "sms": true,
        "email": true,
    },
    "reminder_enable": true,
    "notes": map[string]interface{}{
        "policy_name": "Jeevan Bima",
    },
}
body, err := client.PaymentLink.Create(data, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
|upi_link*          | boolean | boolean Must be set to true   //   to creating UPI Payment Link only                                     |
|amount*        | integer  | Amount to be paid using the Payment Link.                     |
|currency           | string  |  A three-letter ISO code for the currency in which you want to accept the payment. For example, INR.                     |
|accept_partial        | boolean  | Indicates whether customers can make partial payments using the Payment Link. Possible values: true - Customer can make partial payments. false (default) - Customer cannot make partial payments. // UPI Payment Link is not supported partial payment  |
|description           | string  | A brief description of the Payment Link                     |
|first_min_partial_amount           | integer  |Minimum amount, in currency subunits, that must be paid by the customer as the first partial payment. // UPI Payment Link is not supported partial payment  |
|reference_id           | string  | Reference number tagged to a Payment Link.                      |
|customer           | object  | All parameters listed [here](https://razorpay.com/docs/api/payments/payment-links/#sample-codes-for-upi-payment-links) are supported                 |
|expire_by           | integer  | Timestamp, in Unix, at which the Payment Link will expire. By default, a Payment Link will be valid for six months from the date of creation.                     |
|notify           | object  | sms or email (boolean)                     |
|notes           |  object  | Key-value pair that can be used to store additional information about the entity. Maximum 15 key-value pairs, 256 characters (maximum) each. For example, "note_key": "Beam me up Scotty”                     |
| callback_url | string | If specified, adds a redirect URL to the Payment Link. Once customers completes the payment, they are redirected to the specified URL. |
| callback_method | string | If callback_url parameter is passed, callback_method must be passed with the value `get`. |
| reminder_enable | boolean | Used to send reminders for the Payment Link. Possible values is `true` or `false` |

**Response:**
For create payment link response please click [here](https://razorpay.com/docs/api/payment-links/#create-payment-link)

-------------------------------------------------------------------------------------------------------

### Fetch all payment link

```go
body, err := client.PaymentLink.All(nil, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
|payment_id          | string | Unique identifier of the payment associated with the Payment Link.                                               |
|reference_id        | string  | The unique reference number entered by you while creating the Payment Link.                     |

**Response:**
For fetch all payment link response please click [here](https://razorpay.com/docs/api/payment-links/#all-payment-links)

-------------------------------------------------------------------------------------------------------

### Fetch specific payment link

```go
body, err := client.PaymentLink.Fetch("<paymentLinkId>", nil, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
| paymentLinkId*          | string |  Unique identifier of the Payment Link.                         |

**Response:**
For fetch specific payment link response please click [here](https://razorpay.com/docs/api/payment-links/#specific-payment-links-by-id)

-------------------------------------------------------------------------------------------------------

### Update payment link

```go
data := map[string]interface{}{
    "reference_id": "TS35",
    "expire_by": 1653347540,
    "reminder_enable":false,
    "notes":map[string]interface{}{
      "policy_name": "Jeevan Saral",
    },
}
body, err := client.PaymentLink.Update("<paymentLinkId>", data, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
| paymentLinkId*          | string | The unique identifier of the Payment Link that needs to be updated.                         |
| accept_partial         | boolean | Indicates whether customers can make partial payments using the Payment Link. Possible values: true - Customer can make partial payments. false (default) - Customer cannot make partial payments.                         |
| reference_id          | string | Adds a unique reference number to an existing link.                         |
| expire_by         | integer | Timestamp, in Unix format, when the payment links should expire.                         |
| notes          | object | object Key-value pair that can be used to store additional information about the entity. Maximum 15 key-value pairs, 256 characters (maximum) each. For example, "note_key": "Beam me up Scotty”.                         |

**Response:**
For updating payment link response please click [here](https://razorpay.com/docs/api/payment-links/#update-payment-link)

-------------------------------------------------------------------------------------------------------

### Cancel a payment link

```go
body, err := client.PaymentLink.Cancel("<paymentLinkId>", nil, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
| paymentLinkId*          | string | Unique identifier of the Payment Link.                         |

**Response:**
For canceling payment link response please click [here](https://razorpay.com/docs/api/payment-links/#cancel-payment-link)
-------------------------------------------------------------------------------------------------------

### Send notification

```go
body, err := client.PaymentLink.NotifyBy("<paymentLinkId>", "<medium>", nil, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
| paymentLinkId*          | string | Unique identifier of the Payment Link that should be resent.                         |
| medium*          | string | Medium through which the Payment Link must be resent. Allowed values are:  `sms`/`email`         |

**Response:**
```json
{
    "success": true
}
```
-------------------------------------------------------------------------------------------------------

### Transfer payments received using payment links

```go
data := map[string]interface{}{
     "amount": 1500,
     "currency": "INR",
     "accept_partial": false,
     "reference_id": "#aasasw8",
     "description": "Payment for policy no #23456",
     "customer": map[string]interface{}{
       "name": "Gaurav Kumar",
       "contact": "+919999999999",
       "email": "gaurav.kumar@example.com",
     },
     "notify": map[string]interface{}{
       "sms": true,
       "email": true,
     },
     "reminder_enable": true,
     "options": map[string]interface{}{
       "order": map[string]interface{}{
         "transfers": []interface{}{
           map[string]interface{}{
             "account": "acc_HjVXbtpSCIxENR",
             "amount": 500,
             "currency": "INR",
             "notes": map[string]interface{}{
               "branch": "Acme Corp Bangalore North",
               "name": "Bhairav Kumar",
             },
             "linked_account_notes": []string{
               "branch",
             },
           },
           map[string]interface{}{
             "account": "acc_HalyQGZh9ZyiGg",
             "amount": 500,
             "currency": "INR",
             "notes": map[string]interface{}{
               "branch": "Acme Corp Bangalore South",
               "name": "Saurav Kumar",
             },
             "linked_account_notes": []string{
               "branch",
             },
           },
         },
       },
     },
   }
body, err := client.PaymentLink.Create(data, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
|amount*        | integer  | Amount to be paid using the Payment Link.      |
|options*           | object  |  All parameters listed [here](https://razorpay.com/docs/api/payments/payment-links/transfer-payments#request-parameters) are supported |

**Response:**
```json
{
  "accept_partial": false,
  "amount": 1500,
  "amount_paid": 0,
  "cancelled_at": 0,
  "created_at": 1656659763,
  "currency": "INR",
  "customer": {
    "contact": "+919999999999",
    "email": "gaurav.kumar@example.com",
    "name": "Gaurav Kumar"
  },
  "description": "Payment for policy no #23456",
  "expire_by": 0,
  "expired_at": 0,
  "first_min_partial_amount": 0,
  "id": "plink_Jo13NauazdANzM",
  "notes": null,
  "notify": {
    "email": true,
    "sms": true
  },
  "payments": null,
  "reference_id": "#aasasw8",
  "reminder_enable": true,
  "reminders": [],
  "short_url": "https://rzp.io/i/itsF28W",
  "status": "created",
  "updated_at": 1656659763,
  "upi_link": false,
  "user_id": ""
}
```
-------------------------------------------------------------------------------------------------------

### Offers on payment links

```go
data := map[string]interface{}{
     "amount": 3400,
     "currency": "INR",
     "accept_partial": false,
     "reference_id": "#425",
     "description": "Payment for policy no #23456",
     "customer": map[string]interface{}{
       "name": "Gaurav Kumar",
       "contact": "+919999999999",
       "email": "gaurav.kumar@example.com",
     },
     "notify": map[string]interface{}{
       "sms": true,
       "email": true,
     },
     "reminder_enable": false,
     "options": map[string]interface{}{
       "order": map[string]interface{}{
         "offers": []string{
           "offer_JGQvQtvJmVDRIA",
         },
       },
     },
   }
body, err := client.PaymentLink.Create(data, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
|amount*        | integer  | Amount to be paid using the Payment Link.                     |
|currency           | string  |  A three-letter ISO code for the currency in which you want to accept the payment. For example, INR.                     |
|description           | string  | A brief description of the Payment Link                     |
|reference_id           | string  | AReference number tagged to a Payment Link.                      |
|customer           | object  | name, email, contact                 |
|expire_by           | integer  | Timestamp, in Unix, at which the Payment Link will expire. By default, a Payment Link will be valid for six months from the date of creation.                     |
|notify           | object  | sms or email (boolean)                     |
|options*        | object  | Options to associate the offer_id with the Payment Link. Parent parameter under which the order child parameter must be passed.                     |

**Response:**
```json
{
  "accept_partial": false,
  "amount": 3400,
  "amount_paid": 0,
  "cancelled_at": 0,
  "created_at": 1600183040,
  "currency": "INR",
  "customer": {
    "contact": "+919999999999",
    "email": "gaurav.kumar@example.com",
    "name": "Gaurav Kumar"
  },
  "description": "Payment for policy no #23456",
  "expire_by": 0,
  "expired_at": 0,
  "first_min_partial_amount": 0,
  "id": "plink_FdLt0WBldRyE5t",
  "notes": null,
  "notify": {
    "email": true,
    "sms": true
  },
  "payments": null,
  "reference_id": "#425",
  "reminder_enable": false,
  "reminders": [],
  "short_url": "https://rzp.io/i/CM5ohDC",
  "status": "created",
  "user_id": ""
}
```
-------------------------------------------------------------------------------------------------------

### Managing reminders for payment links

```go
data := map[string]interface{}{
      "amount": 1000,
      "currency": "INR",
      "accept_partial": true,
      "first_min_partial_amount": 100,
      "reference_id": "#425",
      "description": "Payment for policy no #23456",
      "customer": map[string]interface{}{
        "name": "Gaurav Kumar",
        "contact": "+919999999999",
        "email": "gaurav.kumar@example.com",
      },
      "notify": map[string]interface{}{
        "sms": true,
        "email": true,
      },
      "reminder_enable": false,
    }
body, err := client.PaymentLink.Create(data, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
|amount*        | integer  | Amount to be paid using the Payment Link.                     |
|accept_partial        | boolean  |  Indicates whether customers can make partial payments using the Payment Link. Possible values:true - Customer can make partial payments.false (default) - Customer cannot make partial payments.                     |
|currency           | string  |  A three-letter ISO code for the currency in which you want to accept the payment. For example, INR.                     |
|description           | string  | A brief description of the Payment Link                     |
|customer           | object  | name, email, contact                 |
|expire_by           | integer  | Timestamp, in Unix, at which the Payment Link will expire. By default, a Payment Link will be valid for six months from the date of creation.                     |
|notify           | object  | sms or email (boolean)                     |
|reminder_enable       | boolean  | To disable reminders for a Payment Link, pass reminder_enable as false                     |

**Response:**
```json
{
  "amount": 340000,
  "amount_due": 340000,
  "amount_paid": 0,
  "billing_end": null,
  "billing_start": null,
  "cancelled_at": null,
  "comment": null,
  "created_at": 1592579126,
  "currency": "INR",
  "currency_symbol": "₹",
  "customer_details": {
    "billing_address": null,
    "contact": "9900990099",
    "customer_contact": "9900990099",
    "customer_email": "gaurav.kumar@example.com",
    "customer_name": "Gaurav Kumar",
    "email": "gaurav.kumar@example.com",
    "gstin": null,
    "id": "cust_F4WNtqj1xb0Duv",
    "name": "Gaurav Kumar",
    "shipping_address": null
  },
  "customer_id": "cust_F4WNtqj1xb0Duv",
  "date": 1592579126,
  "description": "Salon at Home Service",
  "email_status": null,
  "entity": "invoice",
  "expire_by": 1608390326,
  "expired_at": null,
  "first_payment_min_amount": 0,
  "gross_amount": 340000,
  "group_taxes_discounts": false,
  "id": "inv_F4WfpZLk1ct35b",
  "invoice_number": null,
  "issued_at": 1592579126,
  "line_items": [],
  "notes": [],
  "order_id": "order_F4WfpxUzWmYOTl",
  "paid_at": null,
  "partial_payment": false,
  "payment_id": null,
  "receipt": "5757",
  "reminder_enable": false,
  "short_url": "https://rzp.io/i/vitLptM",
  "sms_status": null,
  "status": "issued",
  "tax_amount": 0,
  "taxable_amount": 0,
  "terms": null,
  "type": "link",
  "user_id": "",
  "view_less": true
}
```
-------------------------------------------------------------------------------------------------------

### Rename labels in checkout section

```go
data := map[string]interface{}{
   "amount": 1000,
   "currency": "INR",
   "accept_partial": true,
   "first_min_partial_amount": 100,
   "reference_id": "#421",
   "description": "Payment for policy no #23456",
   "customer": map[string]interface{}{
     "name": "Gaurav Kumar",
     "contact": "+919999999999",
     "email": "gaurav.kumar@example.com",
   },
   "notify": map[string]interface{}{
         "sms": true,
         "email": true,
   },
   "reminder_enable": true,
   "options": map[string]interface{}{
     "checkout": map[string]interface{}{
         "partial_payment": map[string]interface{}{
             "min_amount_label": "Minimum Money to be paid",
             "partial_amount_label": "Pay in parts",
             "partial_amount_description": "Pay at least ₹100",
             "full_amount_label": "Pay the entire amount",
         },
     },
   },
}
body, err := client.PaymentLink.Create(data, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
|amount*        | integer  | Amount to be paid using the Payment Link.                     |
|accept_partial        | boolean  |  Indicates whether customers can make partial payments using the Payment Link. Possible values:true - Customer can make partial payments.false (default) - Customer cannot make partial payments.                     |
|currency           | string  |  A three-letter ISO code for the currency in which you want to accept the payment. For example, INR.                     |
|description           | string  | A brief description of the Payment Link                     |
|customer           | object  | name, email, contact                 |
|expire_by           | integer  | Timestamp, in Unix, at which the Payment Link will expire. By default, a Payment Link will be valid for six months from the date of creation.                     |
|notify           | object  | sms or email (boolean)                     |
|reminder_enable       | boolean  | To disable reminders for a Payment Link, pass reminder_enable as false                     |
|options*       | object  | Options to rename the labels for partial payment fields in the checkout form. Parent parameter under which the checkout and partial_payment child parameters must be passed. |

**Response:**
```json
{
  "accept_partial": true,
  "amount": 1000,
  "amount_paid": 0,
  "callback_method": "",
  "callback_url": "",
  "cancelled_at": 0,
  "created_at": 1596193199,
  "currency": "INR",
  "customer": {
    "contact": "+919999999999",
    "email": "gaurav.kumar@example.com",
    "name": "Gaurav Kumar"
  },
  "deleted_at": 0,
  "description": "Payment for policy no #23456",
  "expire_by": 0,
  "expired_at": 0,
  "first_min_partial_amount": 100,
  "id": "plink_FL4vbXVKfW7PAz",
  "notes": null,
  "notify": {
    "email": true,
    "sms": true
  },
  "payments": null,
  "reference_id": "#42321",
  "reminder_enable": true,
  "reminders": [],
  "short_url": "https://rzp.io/i/F4GC9z1",
  "source": "",
  "source_id": "",
  "status": "created",
  "updated_at": 1596193199,
  "user_id": ""
}
```
-------------------------------------------------------------------------------------------------------

### Change Business name

```go
data := map[string]interface{}{
  "amount": 1000,
  "currency": "INR",
  "accept_partial": true,
  "first_min_partial_amount": 100,
  "reference_id": "#2234542",
  "description": "Payment for policy no #23456",
  "customer": map[string]interface{}{
    "name": "Gaurav Kumar",
    "contact": "+919999999999",
    "email": "gaurav.kumar@example.com",
  },
  "notify": map[string]interface{}{
    "sms": true,
    "email": true,
  },
  "reminder_enable": true,
  "options": map[string]interface{}{
    "checkout": map[string]interface{}{
      "name": "Lacme Corp",
    },
  },
}
body, err := client.PaymentLink.Create(data, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
|amount*        | integer  | Amount to be paid using the Payment Link.                     |
|currency           | string  |  A three-letter ISO code for the currency in which you want to accept the payment. For example, INR.                     |
|accept_partial        | boolean  |  Indicates whether customers can make partial payments using the Payment Link. Possible values:true - Customer can make partial payments.false (default) - Customer cannot make partial payments.                     |
|first_min_partial_amount        | integer  |                      |
|description           | string  | A brief description of the Payment Link                     |
|customer           | object  | name, email, contact                 |
|notify           | object  | sms or email (boolean)                     |
|reminder_enable       | boolean  | To disable reminders for a Payment Link, pass reminder_enable as false                     |
|options*       | object  | Option to customize the business name. Parent parameter under which the checkout child parameter must be passed.|

**Response:**
```json
{
  "accept_partial": true,
  "amount": 1000,
  "amount_paid": 0,
  "callback_method": "",
  "callback_url": "",
  "cancelled_at": 0,
  "created_at": 1596187657,
  "currency": "INR",
  "customer": {
    "contact": "+919999999999",
    "email": "gaurav.kumar@example.com",
    "name": "Gaurav Kumar"
  },
  "description": "Payment for policy no #23456",
  "expire_by": 0,
  "expired_at": 0,
  "first_min_partial_amount": 100,
  "id": "plink_FL3M2gJFs1Jkma",
  "notes": null,
  "notify": {
    "email": true,
    "sms": true
  },
  "payments": null,
  "reference_id": "#2234542",
  "reminder_enable": true,
  "reminders": [],
  "short_url": "https://rzp.io/i/at2OOsR",
  "source": "",
  "source_id": "",
  "status": "created",
  "updated_at": 1596187657,
  "user_id": ""
}
```
-------------------------------------------------------------------------------------------------------

### Prefill checkout fields

```go
data := map[string]interface{}{
  "amount": 1000,
  "currency": "INR",
  "accept_partial": true,
  "first_min_partial_amount": 100,
  "reference_id": "#417",
  "description": "Payment for policy no #23456",
  "customer": map[string]interface{}{
    "name": "Gaurav Kumar",
    "contact": "+919999999999",
    "email": "gaurav.kumar@example.com",
  },
  "notify": map[string]interface{}{
    "sms": true,
    "email": true,
  },
  "reminder_enable": true,
  "options": map[string]interface{}{
    "checkout": map[string]interface{}{
      "prefill": map[string]interface{}{
        "method": "card",
        "card[name]": "Gaurav Kumar",
        "card[number]": "4111111111111111",
        "card[expiry]": "12/21",
        "card[cvv]": "123",
      },
    },
  },
}
body, err := client.PaymentLink.Create(data, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
|amount*        | integer  | Amount to be paid using the Payment Link.                     |
|currency           | string  |  A three-letter ISO code for the currency in which you want to accept the payment. For example, INR.                     |
|accept_partial        | boolean  |  Indicates whether customers can make partial payments using the Payment Link. Possible values:true - Customer can make partial payments.false (default) - Customer cannot make partial payments.                     |
|first_min_partial_amount        | integer  |                      |
|description           | string  | A brief description of the Payment Link                     |
|customer           | object  | name, email, contact                 |
|notify           | object  | sms or email (boolean)                     |
|reminder_enable       | boolean  | To disable reminders for a Payment Link, pass reminder_enable as false                     |
|options*       | object  | Options to customize Checkout. Parent parameter under which the checkout and prefill child parameters must be passed.|

**Response:**
For prefill checkout fields response please click [here](https://razorpay.com/docs/payment-links/api/new/advanced-options/customize/prefill/)

-------------------------------------------------------------------------------------------------------

### Customize payment methods

```go
data := map[string]interface{}{
  "amount": 1000,
  "currency": "INR",
  "accept_partial": true,
  "first_min_partial_amount": 100,
  "reference_id": "#523442",
  "description": "Payment for policy no #23456",
  "customer": map[string]interface{}{
    "name": "Gaurav Kumar",
    "contact": "+919999999999",
    "email": "gaurav.kumar@example.com",
  },
  "notify": map[string]interface{}{
    "sms": true,
    "email": true,
  },
  "reminder_enable": true,
  "options": map[string]interface{}{
    "checkout": map[string]interface{}{
      "method": map[string]interface{}{
        "netbanking": "1",
        "card": "1",
        "upi": "0",
        "wallet": "0",
      },
    },
  },
}
body, err := client.PaymentLink.Create(data, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
|amount*        | integer  | Amount to be paid using the Payment Link.                     |
|currency           | string  |  A three-letter ISO code for the currency in which you want to accept the payment. For example, INR.                     |
|accept_partial        | boolean  |  Indicates whether customers can make partial payments using the Payment Link. Possible values:true - Customer can make partial payments.false (default) - Customer cannot make partial payments.                     |
|first_min_partial_amount        | integer  |                      |
|description           | string  | A brief description of the Payment Link                     |
|customer           | object  | name, email, contact                 |
|notify           | object  | sms or email (boolean)                     |
|reminder_enable       | boolean  | To disable reminders for a Payment Link, pass reminder_enable as false                     |
|options*       | object  | Options to display or hide payment methods on the Checkout section. Parent parameter under which the checkout and method child parameters must be passed.|

**Response:**
```json
{
  "accept_partial": true,
  "amount": 1000,
  "amount_paid": 0,
  "callback_method": "",
  "callback_url": "",
  "cancelled_at": 0,
  "created_at": 1596188371,
  "currency": "INR",
  "customer": {
    "contact": "+919999999999",
    "email": "gaurav.kumar@example.com",
    "name": "Gaurav Kumar"
  },
  "deleted_at": 0,
  "description": "Payment for policy no #23456",
  "expire_by": 0,
  "expired_at": 0,
  "first_min_partial_amount": 100,
  "id": "plink_FL3YbdvN2Cj6gh",
  "notes": null,
  "notify": {
    "email": true,
    "sms": true
  },
  "payments": null,
  "reference_id": "#543422",
  "reminder_enable": true,
  "reminders": [],
  "short_url": "https://rzp.io/i/wKiXKud",
  "source": "",
  "source_id": "",
  "status": "created",
  "updated_at": 1596188371,
  "user_id": ""
}
```
-------------------------------------------------------------------------------------------------------

### Set checkout fields as read-only

```go
data := map[string]interface{}{
  "amount": 1000,
  "currency": "INR",
  "accept_partial": true,
  "first_min_partial_amount": 100,
  "reference_id": "#20",
  "description": "Payment for policy no #23456",
  "customer": map[string]interface{}{
    "name": "Gaurav Kumar",
    "contact": "+919999999999",
    "email": "gaurav.kumar@example.com",
  },
  "notify": map[string]interface{}{
    "sms": true,
    "email": true,
  },
  "reminder_enable": true,
  "options": map[string]interface{}{
    "checkout": map[string]interface{}{
      "readonly": map[string]interface{}{
        "email": "1",
        "contact": "1",
      },
    },
  },
}
body, err := client.PaymentLink.Create(data, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
|amount*        | integer  | Amount to be paid using the Payment Link.                     |
|currency           | string  |  A three-letter ISO code for the currency in which you want to accept the payment. For example, INR.                     |
|accept_partial        | boolean  |  Indicates whether customers can make partial payments using the Payment Link. Possible values:true - Customer can make partial payments.false (default) - Customer cannot make partial payments.                     |
|first_min_partial_amount        | integer  |                      |
|description           | string  | A brief description of the Payment Link                     |
|customer           | object  | name, email, contact                 |
|notify           | object  | sms or email (boolean)                     |
|reminder_enable       | boolean  | To disable reminders for a Payment Link, pass reminder_enable as false                     |
|options*       | object  | Options to set contact and email as read-only fields on Checkout. Parent parameter under which the checkout and readonly child parameters must be passed.|

**Response:**
```json
{
  "accept_partial": true,
  "amount": 1000,
  "amount_paid": 0,
  "callback_method": "",
  "callback_url": "",
  "cancelled_at": 0,
  "created_at": 1596190845,
  "currency": "INR",
  "customer": {
    "contact": "+919999999999",
    "email": "gaurav.kumar@example.com",
    "name": "Gaurav Kumar"
  },
  "deleted_at": 0,
  "description": "Payment for policy no #23456",
  "expire_by": 0,
  "expired_at": 0,
  "first_min_partial_amount": 100,
  "id": "plink_FL4GA1t6FBcaVR",
  "notes": null,
  "notify": {
    "email": true,
    "sms": true
  },
  "payments": null,
  "reference_id": "#19129",
  "reminder_enable": true,
  "reminders": [],
  "short_url": "https://rzp.io/i/QVwUglR",
  "source": "",
  "source_id": "",
  "status": "created",
  "updated_at": 1596190845,
  "user_id": ""
}
```
-------------------------------------------------------------------------------------------------------

### Implement thematic changes in payment links checkout section

```go
data := map[string]interface{}{
  "amount": 1000,
  "currency": "INR",
  "accept_partial": true,
  "first_min_partial_amount": 100,
  "reference_id": "#423212",
  "description": "Payment for policy no #23456",
  "customer": map[string]interface{}{
    "name": "Gaurav Kumar",
    "contact": "+919999999999",
    "email": "gaurav.kumar@example.com",
  },
  "notify": map[string]interface{}{
    "sms": true,
    "email": true,
  },
  "reminder_enable": true,
  "options": map[string]interface{}{
    "checkout": map[string]interface{}{
      "theme": map[string]interface{}{
        "hide_topbar": true,
      },
    },
  },
}
body, err := client.PaymentLink.Create(data, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
|amount*        | integer  | Amount to be paid using the Payment Link.                     |
|currency           | string  |  A three-letter ISO code for the currency in which you want to accept the payment. For example, INR.                     |
|accept_partial        | boolean  |  Indicates whether customers can make partial payments using the Payment Link. Possible values:true - Customer can make partial payments.false (default) - Customer cannot make partial payments.                     |
|first_min_partial_amount        | integer  |                      |
|description           | string  | A brief description of the Payment Link                     |
|customer           | object  | name, email, contact                 |
|notify           | object  | sms or email (boolean)                     |
|reminder_enable       | boolean  | To disable reminders for a Payment Link, pass reminder_enable as false                     |
|options*       | object  | Options to show or hide the top bar. Parent parameter under which the checkout and theme child parameters must be passed.|

**Response:**
```json
{
  "accept_partial": true,
  "amount": 1000,
  "amount_paid": 0,
  "callback_method": "",
  "callback_url": "",
  "cancelled_at": 0,
  "created_at": 1596187814,
  "currency": "INR",
  "customer": {
    "contact": "+919999999999",
    "email": "gaurav.kumar@example.com",
    "name": "Gaurav Kumar"
  },
  "description": "Payment for policy no #23456",
  "expire_by": 0,
  "expired_at": 0,
  "first_min_partial_amount": 100,
  "id": "plink_FL3Oncr7XxXFf6",
  "notes": null,
  "notify": {
    "email": true,
    "sms": true
  },
  "payments": null,
  "reference_id": "#423212",
  "reminder_enable": true,
  "reminders": [],
  "short_url": "https://rzp.io/i/j45EmLE",
  "source": "",
  "source_id": "",
  "status": "created",
  "updated_at": 1596187814,
  "user_id": ""
}
```
-------------------------------------------------------------------------------------------------------

### Rename labels in payment details section

```go
data := map[string]interface{}{
  "amount": 1000,
  "currency": "INR",
  "accept_partial": true,
  "first_min_partial_amount": 100,
  "reference_id": "#412232",
  "description": "Payment for policy no #23456",
  "expire_by": 1599193801,
  "customer": map[string]interface{}{
    "name": "Gaurav Kumar",
    "contact": "+919999999999",
    "email": "gaurav.kumar@example.com",
  },
  "notify": map[string]interface{}{
    "sms": true,
    "email": true,
  },
  "reminder_enable": true,
  "options": map[string]interface{}{
    "hosted_page": map[string]interface{}{
      "label": map[string]interface{}{
        "receipt": "Ref No.",
        "description": "Course Name",
        "amount_payable": "Course Fee Payable",
        "amount_paid": "Course Fee Paid",
        "partial_amount_due": "Fee Installment Due",
        "partial_amount_paid": "Fee Installment Paid",
        "expire_by": "Pay Before",
        "expired_on": "Link Expired. Please contact Admin",
        "amount_due": "Course Fee Due",
      },
      "show_preferences": map[string]interface{}{
        "issued_to": false,
      },
    },
  },
}
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
|amount*        | integer  | Amount to be paid using the Payment Link.                     |
|currency           | string  |  A three-letter ISO code for the currency in which you want to accept the payment. For example, INR.                     |
|accept_partial        | boolean  |  Indicates whether customers can make partial payments using the Payment Link. Possible values:true - Customer can make partial payments.false (default) - Customer cannot make partial payments.                     |
|first_min_partial_amount        | integer  |                      |
|description           | string  | A brief description of the Payment Link                     |
|customer           | object  | name, email, contact                 |
|notify           | object  | sms or email (boolean)                     |
|reminder_enable       | boolean  | To disable reminders for a Payment Link, pass reminder_enable as false                     |
|options*       | object  | Parent parameter under which the hosted_page and label child parameters must be passed.|

**Response:**
For rename labels in payment details section response please click [here](https://razorpay.com/docs/payment-links/api/new/advanced-options/customize/rename-payment-details-labels/)

-------------------------------------------------------------------------------------------------------

**PN: * indicates mandatory fields**
<br>
<br>
**For reference click [here](https://razorpay.com/docs/api/payment-links/)**