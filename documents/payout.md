
# Payout 

## Fetch All Payouts

```go
queryParams := map[string]interface{}{
    "account_number": "7878780080316316",
}
payouts, err := client.Payout.All(queryParams, nil)
```

### Parameters

| Name           | Type    | Description                              |
|----------------|---------|------------------------------------------|
| account_number | string  | Filter payouts by account number         |
| count          | integer | Number of payouts to fetch (default: 10) |
| skip           | integer | Number of payouts to skip (default: 0)   |

### Response

```json
{
  "entity": "collection",
  "count": 2,
  "items": [
    {
      "id": "pout_00000000000001",
      "entity": "payout",
      "fund_account_id": "fa_00000000000001",
      "amount": 1000000,
      "currency": "INR",
      "notes": {
        "notes_key_1": "Tea, Earl Grey, Hot",
        "notes_key_2": "Tea, Earl Grey… decaf."
      },
      "fees": 590,
      "tax": 90,
      "status": "processed",
      "purpose": "payout",
      "utr": null,
      "mode": "NEFT",
      "reference_id": "Acme Transaction ID 12345",
      "narration": "Acme Corp Fund Transfer",
      "batch_id": null,
      "status_details": {
        "description": "Payout is processed and the money has been credited into the beneficiaries account",
        "source": "beneficiary_bank",
        "reason": "payout_processed"
      },
      "created_at": 1545382870,
      "fee_type": ""
    },
    {
      "id": "pout_00000000000002",
      "entity": "payout",
      "fund_account_id": "fa_00000000000002",
      "amount": 1000000,
      "currency": "INR",
      "notes": {
        "notes_key_1": "Tea, Earl Grey, Hot",
        "notes_key_2": "Tea, Earl Grey… decaf."
      },
      "fees": 590,
      "tax": 90,
      "status": "reversed",
      "purpose": "refund",
      "utr": null,
      "mode": "NEFT",
      "reference_id": "Acme Transaction ID 123456",
      "narration": "Acme Corp Fund Transfer",
      "batch_id": null,
      "status_details": {
        "description": "The NEFT 24*7 limits for your account has been exhausted. Please retry after sometime",
        "source": "business",
        "reason": "amount_limit_exhausted"
      },
      "created_at": 1545382870,
      "fee_type": ""
    }
  ]
}
```

---

## Fetch a Payout By ID

```go
payout, err := client.Payout.Fetch("pout_00000000000001", nil, nil)
```

### Parameters

| Name     | Type   | Description                     |
|----------|--------|---------------------------------|
| payoutID | string | The ID of the payout to fetch   |

### Response

```json
{
  "id": "pout_00000000000001",
  "entity": "payout",
  "fund_account_id": "fa_00000000000001",
  "amount": 1000000,
  "currency": "INR",
  "notes": {
    "note_key": "Beam me up Scotty"
  },
  "fees": 590,
  "tax": 90,
  "status": "processed",
  "purpose": "payout",
  "utr": null,
  "mode": "NEFT",
  "reference_id": "Acme Transaction ID 12345",
  "narration": "Acme Corp Fund Transfer",
  "batch_id": null,
  "status_details": {
    "description": "Payout is processed and the money has been credited into the beneficiaries account",
    "source": "beneficiary_bank",
    "reason": "payout_processed"
  },
  "created_at": 1545382870,
  "fee_type": ""
}
```
