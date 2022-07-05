## Invoices

### Create Invoice

Request #1
In this example, an invoice is created using the customer and item details. Here, the customer and item are created while creating the invoice.
```go

line_items := make(map[string]interface{})
line_items["0"] = map[string]interface{}{
      "name": "Master Cloud Computing in 30 Days",
      "description": "Book by Ravena Ravenclaw",
      "amount": 399,
      "currency": "INR",
      "quantity": 1,
    }

data := map[string]interface{}{
  "type": "invoice",
  "description": "Invoice for the month of January 2020",
  "partial_payment": true,
  "customer": map[string]interface{}{
    "name": "Gaurav Kumar",
    "contact": 9999999999,
    "email": "gaurav.kumar@example.com",
    "billing_address": map[string]interface{}{
      "line1": "Ground & 1st Floor, SJR Cyber Laskar",
      "line2": "Hosur Road",
      "zipcode": 560068,
      "city": "Bengaluru",
      "state": "Karnataka",
      "country": "in",
    },
    "shipping_address": map[string]interface{}{
      "line1": "Ground & 1st Floor, SJR Cyber Laskar",
      "line2": "Hosur Road",
      "zipcode": 560068,
      "city": "Bengaluru",
      "state": "Karnataka",
      "country": "in",
    },
  },
  "line_items": line_items,
  "sms_notify": 1,
  "email_notify": 1,
  "currency": "INR",
  "expire_by": 1589765167,
}
body, err := client.Invoice.Create(data, nil)
```
**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
|type*          | string | entity type (here its invoice)                                               |
|description        | string  | A brief description of the invoice.                      |
|customer_id           | string  | customer id for which invoice need be raised   |
|draft           | string  |  Invoice is created in draft state when value is set to `1`   |
| customer*     | object | All parameters listed [here](https://razorpay.com/docs/api/payments/invoices/#create-an-invoice) are supported           |
| line_items    | object | All parameters listed [here](https://razorpay.com/docs/api/payments/invoices/#create-an-invoice) are supported |
|expire_by           | integer  | Details of the line item that is billed in the invoice.  |
|sms_notify           | integer  | Details of the line item that is billed in the invoice.  |
|email_notify           | boolean  | Details of the line item that is billed in the invoice.  |
|partial_payment | boolean  | Indicates whether customers can make partial payments on the invoice . Possible values: true - Customer can make partial payments. false (default) - Customer cannot make partial payments. |
| currency*   | string  | The currency of the payment (defaults to INR)  |

Request #2
In this example, an invoice is created using existing `customer_id` and `item_id`
```go
line_items := make(map[string]interface{})
line_items["0"] = map[string]interface{}{
      "item_id": "item_DRt61i2NnL8oy6",
}

data := map[string]interface{}{
  "type": "invoice",
  "date": 1589994898,
  "customer_id": "cust_E7q0trFqXgExmT",
  "line_items": line_items,
}

body, err := client.Invoice.Create(data, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
|type*          | string | entity type (here its invoice)                                               |
|description        | string  | A brief description of the invoice.                      |
|customer_id           | string  | customer id for which invoice need be raised                     |
| customer*     | object | All parameters listed [here](https://razorpay.com/docs/api/payments/invoices/#create-an-invoice) are supported           |
| line_items    | object | All parameters listed [here](https://razorpay.com/docs/api/payments/invoices/#create-an-invoice) are supported |
| sms_notify  | boolean  | SMS notifications are to be sent by Razorpay (default : 1)  |
| currency*  (conditionally mandatory) | string  | The 3-letter ISO currency code for the payment. Currently, only `INR` is supported. |
| email_notify | boolean  | Email notifications are to be sent by Razorpay (default : 1)  |
| expire_by    | integer | The timestamp, in Unix format, till when the customer can make the authorization payment. |


**Response:**
For create invoice response please click [here](https://razorpay.com/docs/api/invoices/#create-an-invoice)

-------------------------------------------------------------------------------------------------------

### Fetch all invoices

```go
option := map[string]interface{}{
    "count" : 1,
}
body, err := client.Invoice.All(option, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
|type          | string | entity type (here its invoice)                                               |
|payment_id        | string  | The unique identifier of the payment made by the customer against the invoice.                      |
|customer_id           | string  | The unique identifier of the customer.                     |
|receipt           | string  |  The unique receipt number that you entered for internal purposes.                     |

**Response:**
For fetch all invoice response please click [here](https://razorpay.com/docs/api/invoices/#fetch-multiple-invoices)

-------------------------------------------------------------------------------------------------------

### Fetch invoice

```go
body, err := client.Invoice.Fetch("<invoiceId>", nil, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
| invoiceId*          | string | The id of the invoice to be fetched                         |

**Response:**
```json
{
  "id": "inv_E7q0tqkxBRzdau",
  "entity": "invoice",
  "receipt": null,
  "invoice_number": null,
  "customer_id": "cust_E7q0trFqXgExmT",
  "customer_details": {
    "id": "cust_E7q0trFqXgExmT",
    "name": "Gaurav Kumar",
    "email": "gaurav.kumar@example.com",
    "contact": "9999999999",
    "gstin": null,
    "billing_address": {
      "id": "addr_E7q0ttqh4SGhAC",
      "type": "billing_address",
      "primary": true,
      "line1": "Ground & 1st Floor, SJR Cyber Laskar",
      "line2": "Hosur Road",
      "zipcode": "560068",
      "city": "Bengaluru",
      "state": "Karnataka",
      "country": "in"
    },
    "shipping_address": {
      "id": "addr_E7q0ttKwVA1h2V",
      "type": "shipping_address",
      "primary": true,
      "line1": "Ground & 1st Floor, SJR Cyber Laskar",
      "line2": "Hosur Road",
      "zipcode": "560068",
      "city": "Bengaluru",
      "state": "Karnataka",
      "country": "in"
    },
    "customer_name": "Gaurav Kumar",
    "customer_email": "gaurav.kumar@example.com",
    "customer_contact": "9999999999"
  },
  "order_id": "order_E7q0tvRpC0WJwg",
  "line_items": [
    {
      "id": "li_E7q0tuPNg84VbZ",
      "item_id": null,
      "ref_id": null,
      "ref_type": null,
      "name": "Master Cloud Computing in 30 Days",
      "description": "Book by Ravena Ravenclaw",
      "amount": 399,
      "unit_amount": 399,
      "gross_amount": 399,
      "tax_amount": 0,
      "taxable_amount": 399,
      "net_amount": 399,
      "currency": "INR",
      "type": "invoice",
      "tax_inclusive": false,
      "hsn_code": null,
      "sac_code": null,
      "tax_rate": null,
      "unit": null,
      "quantity": 1,
      "taxes": []
    }
  ],
  "payment_id": null,
  "status": "issued",
  "expire_by": 1589765167,
  "issued_at": 1579765167,
  "paid_at": null,
  "cancelled_at": null,
  "expired_at": null,
  "sms_status": "pending",
  "email_status": "pending",
  "date": 1579765167,
  "terms": null,
  "partial_payment": true,
  "gross_amount": 399,
  "tax_amount": 0,
  "taxable_amount": 399,
  "amount": 399,
  "amount_paid": 0,
  "amount_due": 399,
  "currency": "INR",
  "currency_symbol": "₹",
  "description": "Invoice for the month of January 2020",
  "notes": [],
  "comment": null,
  "short_url": "https://rzp.io/i/2wxV8Xs",
  "view_less": true,
  "billing_start": null,
  "billing_end": null,
  "type": "invoice",
  "group_taxes_discounts": false,
  "created_at": 1579765167
}
```
-------------------------------------------------------------------------------------------------------

### Update invoice

```go
line_items := make(map[string]interface{})
line_items["0"] = map[string]interface{}{
      "id": "li_JKqakYLuD5qxvi",
      "name": "Book / English August - Updated name and quantity",
      "quantity": 1,
    }

data:= map[string]interface{}{
  "line_items": line_items,
  "notes": map[string]interface{}{
    "updated-key": "An updated note.",
  },
}
body, err := client.Invoice.Update("<invoiceId>", data, nil)

```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
| invoiceId*          | string | The id of the invoice to be fetched                         |
| line_items    | object | All parameters listed [here](https://razorpay.com/docs/api/payments/invoices/#update-an-invoice) are supported |
| notes         | object      | A key-value pair                            |

**Response:**

```json
{
    "id": "inv_Jm44BTJyzbquiy",
    "entity": "invoice",
    "receipt": "dfdsf",
    "invoice_number": "dfdsf",
    "customer_id": "cust_JjoxLvdd8gjgyi",
    "customer_details": {
        "id": "cust_JjoxLvdd8gjgyi",
        "name": "ypNsfYnGeGE",
        "email": "ypnsfyngege@gmail.com",
        "contact": "9999999999",
        "gstin": null,
        "billing_address": null,
        "shipping_address": null,
        "customer_name": "ypNsfYnGeGE",
        "customer_email": "ypnsfyngege@gmail.com",
        "customer_contact": "9999999999"
    },
    "order_id": null,
    "line_items": [
        {
            "id": "li_JnhMrLhv1PMAoH",
            "item_id": "item_JnQ2kRGq8Kbte3",
            "ref_id": null,
            "ref_type": null,
            "name": "Book / English August - Updated name and quantity",
            "description": "New descirption too. :).",
            "amount": 20000,
            "unit_amount": 20000,
            "gross_amount": 20000,
            "tax_amount": 0,
            "taxable_amount": 20000,
            "net_amount": 20000,
            "currency": "INR",
            "type": "invoice",
            "tax_inclusive": false,
            "hsn_code": null,
            "sac_code": null,
            "tax_rate": null,
            "unit": null,
            "quantity": 1,
            "taxes": []
        },
        {
            "id": "li_JnhXAk87ENeeo1",
            "item_id": null,
            "ref_id": null,
            "ref_type": null,
            "name": "Book / A Wild Sheep Chase",
            "description": null,
            "amount": 200,
            "unit_amount": 200,
            "gross_amount": 200,
            "tax_amount": 0,
            "taxable_amount": 200,
            "net_amount": 200,
            "currency": "INR",
            "type": "invoice",
            "tax_inclusive": false,
            "hsn_code": null,
            "sac_code": null,
            "tax_rate": null,
            "unit": null,
            "quantity": 1,
            "taxes": []
        }
    ],
    "payment_id": null,
    "status": "draft",
    "expire_by": null,
    "issued_at": null,
    "paid_at": null,
    "cancelled_at": null,
    "expired_at": null,
    "sms_status": null,
    "email_status": null,
    "date": 1656233665,
    "terms": null,
    "partial_payment": false,
    "gross_amount": 20200,
    "tax_amount": 0,
    "taxable_amount": 20200,
    "amount": 20200,
    "amount_paid": null,
    "amount_due": null,
    "currency": "INR",
    "currency_symbol": "₹",
    "description": "dsfdsfsdf",
    "notes": {
        "updated-key": "An updated note."
    },
    "comment": null,
    "short_url": null,
    "view_less": true,
    "billing_start": null,
    "billing_end": null,
    "type": "invoice",
    "group_taxes_discounts": false,
    "created_at": 1656233693,
    "idempotency_key": null
}
```
-------------------------------------------------------------------------------------------------------

### Issue an invoice

```go
body, err := client.Invoice.Issue("<invoiceId>", nil, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
| invoiceId*          | string | The id of the invoice to be issued                         |

**Response:**
```json
{
  "id": "inv_DAweOiQ7amIUVd",
  "entity": "invoice",
  "receipt": "#0961",
  "invoice_number": "#0961",
  "customer_id": "cust_DAtUWmvpktokrT",
  "customer_details": {
    "id": "cust_DAtUWmvpktokrT",
    "name": "Gaurav Kumar",
    "email": "gaurav.kumar@example.com",
    "contact": "9977886633",
    "gstin": null,
    "billing_address": {
      "id": "addr_DAtUWoxgu91obl",
      "type": "billing_address",
      "primary": true,
      "line1": "318 C-Wing, Suyog Co. Housing Society Ltd.",
      "line2": "T.P.S Road, Vazira, Borivali",
      "zipcode": "400092",
      "city": "Mumbai",
      "state": "Maharashtra",
      "country": "in"
    },
    "shipping_address": null,
    "customer_name": "Gaurav Kumar",
    "customer_email": "gaurav.kumar@example.com",
    "customer_contact": "9977886633"
  },
  "order_id": "order_DBG3P8ZgDd1dsG",
  "line_items": [
    {
      "id": "li_DAweOizsysoJU6",
      "item_id": null,
      "name": "Book / English August - Updated name and quantity",
      "description": "150 points in Quidditch",
      "amount": 400,
      "unit_amount": 400,
      "gross_amount": 400,
      "tax_amount": 0,
      "taxable_amount": 400,
      "net_amount": 400,
      "currency": "INR",
      "type": "invoice",
      "tax_inclusive": false,
      "hsn_code": null,
      "sac_code": null,
      "tax_rate": null,
      "unit": null,
      "quantity": 1,
      "taxes": []
    },
    {
      "id": "li_DAwjWQUo07lnjF",
      "item_id": null,
      "name": "Book / A Wild Sheep Chase",
      "description": null,
      "amount": 200,
      "unit_amount": 200,
      "gross_amount": 200,
      "tax_amount": 0,
      "taxable_amount": 200,
      "net_amount": 200,
      "currency": "INR",
      "type": "invoice",
      "tax_inclusive": false,
      "hsn_code": null,
      "sac_code": null,
      "tax_rate": null,
      "unit": null,
      "quantity": 1,
      "taxes": []
    }
  ],
  "payment_id": null,
  "status": "issued",
  "expire_by": 1567103399,
  "issued_at": 1566974805,
  "paid_at": null,
  "cancelled_at": null,
  "expired_at": null,
  "sms_status": null,
  "email_status": null,
  "date": 1566891149,
  "terms": null,
  "partial_payment": false,
  "gross_amount": 600,
  "tax_amount": 0,
  "taxable_amount": 600,
  "amount": 600,
  "amount_paid": 0,
  "amount_due": 600,
  "currency": "INR",
  "currency_symbol": "₹",
  "description": "This is a test invoice.",
  "notes": {
    "updated-key": "An updated note."
  },
  "comment": null,
  "short_url": "https://rzp.io/i/K8Zg72C",
  "view_less": true,
  "billing_start": null,
  "billing_end": null,
  "type": "invoice",
  "group_taxes_discounts": false,
  "created_at": 1566906474,
  "idempotency_key": null
}
```
-------------------------------------------------------------------------------------------------------

### Delete an invoice

```go
body, err := client.Invoice.Delete("<invoiceId>", nil, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
| invoiceId*          | string | The id of the invoice to be deleted                         |

**Response:**
```
[]
```
-------------------------------------------------------------------------------------------------------

### Cancel an invoice

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
  "id": "inv_E7q0tqkxBRzdau",
  "entity": "invoice",
  "receipt": null,
  "invoice_number": null,
  "customer_id": "cust_E7q0trFqXgExmT",
  "customer_details": {
    "id": "cust_E7q0trFqXgExmT",
    "name": "Gaurav Kumar",
    "email": "gaurav.kumar@example.com",
    "contact": "9972132594",
    "gstin": null,
    "billing_address": {
      "id": "addr_E7q0ttqh4SGhAC",
      "type": "billing_address",
      "primary": true,
      "line1": "Ground & 1st Floor, SJR Cyber Laskar",
      "line2": "Hosur Road",
      "zipcode": "560068",
      "city": "Bengaluru",
      "state": "Karnataka",
      "country": "in"
    },
    "shipping_address": {
      "id": "addr_E7q0ttKwVA1h2V",
      "type": "shipping_address",
      "primary": true,
      "line1": "Ground & 1st Floor, SJR Cyber Laskar",
      "line2": "Hosur Road",
      "zipcode": "560068",
      "city": "Bengaluru",
      "state": "Karnataka",
      "country": "in"
    },
    "customer_name": "Gaurav Kumar",
    "customer_email": "gaurav.kumar@example.com",
    "customer_contact": "9972132594"
  },
  "order_id": "order_E7q0tvRpC0WJwg",
  "line_items": [
    {
      "id": "li_E7q0tuPNg84VbZ",
      "item_id": null,
      "ref_id": null,
      "ref_type": null,
      "name": "Master Cloud Computing in 30 Days",
      "description": "Book by Ravena Ravenclaw",
      "amount": 399,
      "unit_amount": 399,
      "gross_amount": 399,
      "tax_amount": 0,
      "taxable_amount": 399,
      "net_amount": 399,
      "currency": "INR",
      "type": "invoice",
      "tax_inclusive": false,
      "hsn_code": null,
      "sac_code": null,
      "tax_rate": null,
      "unit": null,
      "quantity": 1,
      "taxes": []
    }
  ],
  "payment_id": null,
  "status": "cancelled",
  "expire_by": null,
  "issued_at": 1579765167,
  "paid_at": null,
  "cancelled_at": 1579768206,
  "expired_at": null,
  "sms_status": "sent",
  "email_status": "sent",
  "date": 1579765167,
  "terms": null,
  "partial_payment": false,
  "gross_amount": 399,
  "tax_amount": 0,
  "taxable_amount": 399,
  "amount": 399,
  "amount_paid": 0,
  "amount_due": 399,
  "currency": "INR",
  "currency_symbol": "₹",
  "description": null,
  "notes": [],
  "comment": null,
  "short_url": "https://rzp.io/i/2wxV8Xs",
  "view_less": true,
  "billing_start": null,
  "billing_end": null,
  "type": "invoice",
  "group_taxes_discounts": false,
  "created_at": 1579765167,
  "idempotency_key": null
}
```
-------------------------------------------------------------------------------------------------------

### Send notification

```go
body, err := client.Invoice.Notify("<invoiceId>", "<medium>", nil, nil)

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
| invoiceId*          | string | The id of the invoice to be notified                         |
| medium*          | string | `sms`/`email`, Medium through which notification should be sent.                         |
```
**Response:**

```json
{
    "success": true
}
```
-------------------------------------------------------------------------------------------------------

**PN: * indicates mandatory fields**
<br>
<br>
**For reference click [here](https://razorpay.com/docs/api/invoices)**