## Document

### Fetch All Disputes

```go
body, err := client.Dispute.All(nil, nil)
```

**Response:**
```json
{
  "entity": "collection",
  "count": 1,
  "items": [
    {
      "id": "disp_Esz7KAitoYM7PJ",
      "entity": "dispute",
      "payment_id": "pay_EsyWjHrfzb59eR",
      "amount": 10000,
      "currency": "INR",
      "amount_deducted": 0,
      "reason_code": "pre_arbitration",
      "respond_by": 1590604200,
      "status": "open",
      "phase": "pre_arbitration",
      "created_at": 1590059211,
      "evidence": {
        "amount": 10000,
        "summary": null,
        "shipping_proof": null,
        "billing_proof": null,
        "cancellation_proof": null,
        "customer_communication": null,
        "proof_of_service": null,
        "explanation_letter": null,
        "refund_confirmation": null,
        "access_activity_log": null,
        "refund_cancellation_policy": null,
        "term_and_conditions": null,
        "others": null,
        "submitted_at": null
      }
    }
  ]
}
```
-------------------------------------------------------------------------------------------------------

### Fetch a Dispute

```go
disputeId := "disp_0000000000000"

body, err := client.Dispute.Fetch(disputeId, nil, nil)
```

**Parameters:**

| Name  | Type      | Description                                      |
|-------|-----------|--------------------------------------------------|
| disputeId*  | string | The unique identifier of the dispute.  |

**Response:**
```json
{
  "id": "disp_AHfqOvkldwsbqt",
  "entity": "dispute",
  "payment_id": "pay_EsyWjHrfzb59eR",
  "amount": 10000,
  "currency": "INR",
  "amount_deducted": 0,
  "reason_code": "pre_arbitration",
  "respond_by": 1590604200,
  "status": "open",
  "phase": "pre_arbitration",
  "created_at": 1590059211,
  "evidence": {
    "amount": 10000,
    "summary": "goods delivered",
    "shipping_proof": null,
    "billing_proof": null,
    "cancellation_proof": null,
    "customer_communication": null,
    "proof_of_service": null,
    "explanation_letter": null,
    "refund_confirmation": null,
    "access_activity_log": null,
    "refund_cancellation_policy": null,
    "term_and_conditions": null,
    "others": null,
    "submitted_at": null
  }
}
```
-------------------------------------------------------------------------------------------------------

### Fetch a Dispute

```go
disputeId := "disp_0000000000000"

body, err := client.Dispute.Accept(disputeId, nil, nil)
```

**Parameters:**

| Name  | Type      | Description                                      |
|-------|-----------|--------------------------------------------------|
| disputeId*  | string | The unique identifier of the dispute.  |

**Response:**
```json
{
  "id": "disp_AHfqOvkldwsbqt",
  "entity": "dispute",
  "payment_id": "pay_EsyWjHrfzb59eR",
  "amount": 10000,
  "currency": "INR",
  "amount_deducted": 10000,
  "reason_code": "pre_arbitration",
  "respond_by": 1590604200,
  "status": "lost",
  "phase": "pre_arbitration",
  "created_at": 1590059211,
  "evidence": {
    "amount": 10000,
    "summary": null,
    "shipping_proof": null,
    "billing_proof": null,
    "cancellation_proof": null,
    "customer_communication": null,
    "proof_of_service": null,
    "explanation_letter": null,
    "refund_confirmation": null,
    "access_activity_log": null,
    "refund_cancellation_policy": null,
    "term_and_conditions": null,
    "others": null,
    "submitted_at": null
  }
}
```
-------------------------------------------------------------------------------------------------------
### Contest a Dispute

```go
// Use this API sample code for draft

disputeId := "disp_0000000000000";

shipping_proof := [2]string{"doc_EFtmUsbwpXwBH9", "doc_EFtmUsbwpXwBH8"}
doc := [2]string{"doc_EFtmUsbwpXwBH1", "doc_EFtmUsbwpXwBH7"}

other := make(map[string]interface{})
other["0"] = map[string]interface{}{
  "type":         "receipt_signed_by_customer",
  "document_ids": doc,
}

params := map[string]interface{}{
  "amount":         5000,
  "summary":        "goods delivered",
  "shipping_proof": shipping_proof,
  "others":         other,
  "action":         "draft",
}

body, err := client.Dispute.Contest(disputeId, params, nil)
```

**Parameters:**

| Name  | Type      | Description                                      |
|-------|-----------|--------------------------------------------------|
| disputeId*  | string | The unique identifier of the dispute.  |
| amount  | integer | The amount being contested. If the contest amount is not mentioned, we will assume it to be a full dispute contest.  |
| summary  | string | The explanation provided by you for contesting the dispute. It can have a maximum length of 1000 characters. |
| shipping_proof  | array | List of document ids which serves as proof that the product was shipped to the customer at their provided address. It should show their complete shipping address, if possible. |
| others  | array | All keys listed [here](https://razorpay.com/docs/api/disputes/contest) are supported |

```go
// Use this API sample code for submit

billing_proof := [2]string{"doc_EFtmUsbwpXwBH9", "doc_EFtmUsbwpXwBH8"}

params := map[string]interface{}{
  "billing_proof": billing_proof,
  "action":         "submit",
}

body, err := client.Dispute.Contest(disputeId, params, nil)
```

**Response:**
```json
// Draft
{
  "id": "disp_AHfqOvkldwsbqt",
  "entity": "dispute",
  "payment_id": "pay_EsyWjHrfzb59eR",
  "amount": 10000,
  "currency": "INR",
  "amount_deducted": 0,
  "reason_code": "chargeback",
  "respond_by": 1590604200,
  "status": "open",
  "phase": "chargeback",
  "created_at": 1590059211,
  "evidence": {
    "amount": 5000,
    "summary": "goods delivered",
    "shipping_proof": [
      "doc_EFtmUsbwpXwBH9",
      "doc_EFtmUsbwpXwBH8"
    ],
    "billing_proof": null,
    "cancellation_proof": null,
    "customer_communication": null,
    "proof_of_service": null,
    "explanation_letter": null,
    "refund_confirmation": null,
    "access_activity_log": null,
    "refund_cancellation_policy": null,
    "term_and_conditions": null,
    "others": [
      {
        "type": "receipt_signed_by_customer",
        "document_ids": [
          "doc_EFtmUsbwpXwBH1",
          "doc_EFtmUsbwpXwBH7"
        ]
      }
    ],
    "submitted_at": null
  }
}

//Submit 
{
  "id": "disp_AHfqOvkldwsbqt",
  "entity": "dispute",
  "payment_id": "pay_EsyWjHrfzb59eR",
  "amount": 10000,
  "currency": "INR",
  "amount_deducted": 0,
  "reason_code": "chargeback",
  "respond_by": 1590604200,
  "status": "under_review",
  "phase": "chargeback",
  "created_at": 1590059211,
  "evidence": {
    "amount": 5000,
    "summary": "goods delivered",
    "shipping_proof": [
      "doc_EFtmUsbwpXwBH9",
      "doc_EFtmUsbwpXwBH8"
    ],
    "billing_proof": [
      "doc_EFtmUsbwpXwBG9",
      "doc_EFtmUsbwpXwBG8"
    ],
    "cancellation_proof": null,
    "customer_communication": null,
    "proof_of_service": null,
    "explanation_letter": null,
    "refund_confirmation": null,
    "access_activity_log": null,
    "refund_cancellation_policy": null,
    "term_and_conditions": null,
    "others": [
      {
        "type": "receipt_signed_by_customer",
        "document_ids": [
          "doc_EFtmUsbwpXwBH1",
          "doc_EFtmUsbwpXwBH7"
        ]
      }
    ],
    "submitted_at": 1590603200
  }
}
```
-------------------------------------------------------------------------------------------------------
**PN: * indicates mandatory fields**
<br>
<br>
**For reference click [here](https://razorpay.com/docs/api/documents)**