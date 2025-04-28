# Payout

The Payout API allows you to create and manage payouts to your customers' bank accounts.

## Create a Payout

```go
params := map[string]interface{}{
    "account_number": "7878780080316316",
    "fund_account_id": "fa_00000000000001",
    "amount": 1000000,
    "currency": "INR",
    "mode": "IMPS",
    "purpose": "refund",
    "queue_if_low_balance": true,
    "reference_id": "Acme Transaction ID 12345",
    "narration": "Acme Corp Fund Transfer",
    "notes": map[string]interface{}{
        "notes_key_1": "Tea, Earl Grey, Hot",
        "notes_key_2": "Tea, Earl Grey… decaf.",
    },
}
payout, err := client.Payout.Create(params, nil)
```

### Parameters

| Name | Type | Description |
|------|------|-------------|
| account_number | string | The account number to which the payout should be made |
| fund_account_id | string | The fund account ID to which the payout should be made |
| amount | integer | The amount to be paid out (in paise) |
| currency | string | The currency of the payout (default: INR) |
| mode | string | The mode of payout (IMPS, NEFT, RTGS) |
| purpose | string | The purpose of the payout |
| queue_if_low_balance | boolean | Whether to queue the payout if the account has low balance |
| reference_id | string | A unique reference ID for the payout |
| narration | string | A description of the payout |
| notes | object | Additional notes for the payout |

### Response

```json
{
    "id": "pout_00000000000001",
    "entity": "payout",
    "fund_account_id": "fa_00000000000001",
    "fund_account": {
        "id": "fa_00000000000001",
        "entity": "fund_account",
        "contact_id": "cont_00000000000001",
        "account_type": "bank_account",
        "bank_account": {
            "name": "Gaurav Kumar",
            "ifsc": "RAZR0000001",
            "account_number": "7878780080316316"
        },
        "active": true
    },
    "amount": 1000000,
    "currency": "INR",
    "notes": {
        "notes_key_1": "Tea, Earl Grey, Hot",
        "notes_key_2": "Tea, Earl Grey… decaf."
    },
    "fees": 0,
    "tax": 0,
    "status": "processing",
    "utr": null,
    "mode": "IMPS",
    "purpose": "refund",
    "reference_id": "Acme Transaction ID 12345",
    "narration": "Acme Corp Fund Transfer",
    "created_at": 1566985740
}
```

## Fetch All Payouts

```go
queryParams := map[string]interface{}{
    "account_number": "7878780080316316",
}
payouts, err := client.Payout.All(queryParams, nil)
```

### Parameters

| Name | Type | Description |
|------|------|-------------|
| account_number | string | Filter payouts by account number |
| count | integer | Number of payouts to fetch (default: 10) |
| skip | integer | Number of payouts to skip (default: 0) |

### Response

```json
{
    "entity": "collection",
    "count": 1,
    "items": [
        {
            "id": "pout_00000000000001",
            "entity": "payout",
            "fund_account_id": "fa_00000000000001",
            "fund_account": {
                "id": "fa_00000000000001",
                "entity": "fund_account",
                "contact_id": "cont_00000000000001",
                "account_type": "bank_account",
                "bank_account": {
                    "name": "Gaurav Kumar",
                    "ifsc": "RAZR0000001",
                    "account_number": "7878780080316316"
                },
                "active": true
            },
            "amount": 1000000,
            "currency": "INR",
            "notes": {
                "notes_key_1": "Tea, Earl Grey, Hot",
                "notes_key_2": "Tea, Earl Grey… decaf."
            },
            "fees": 0,
            "tax": 0,
            "status": "processing",
            "utr": null,
            "mode": "IMPS",
            "purpose": "refund",
            "reference_id": "Acme Transaction ID 12345",
            "narration": "Acme Corp Fund Transfer",
            "created_at": 1566985740
        }
    ]
}
```

## Fetch a Payout By ID

```go
payout, err := client.Payout.Fetch("pout_00000000000001", nil, nil)
```

### Parameters

| Name | Type | Description |
|------|------|-------------|
| payoutID | string | The ID of the payout to fetch |

### Response

```json
{
    "id": "pout_00000000000001",
    "entity": "payout",
    "fund_account_id": "fa_00000000000001",
    "fund_account": {
        "id": "fa_00000000000001",
        "entity": "fund_account",
        "contact_id": "cont_00000000000001",
        "account_type": "bank_account",
        "bank_account": {
            "name": "Gaurav Kumar",
            "ifsc": "RAZR0000001",
            "account_number": "7878780080316316"
        },
        "active": true
    },
    "amount": 1000000,
    "currency": "INR",
    "notes": {
        "notes_key_1": "Tea, Earl Grey, Hot",
        "notes_key_2": "Tea, Earl Grey… decaf."
    },
    "fees": 0,
    "tax": 0,
    "status": "processing",
    "utr": null,
    "mode": "IMPS",
    "purpose": "refund",
    "reference_id": "Acme Transaction ID 12345",
    "narration": "Acme Corp Fund Transfer",
    "created_at": 1566985740
}
```

### Response Fields

| Name | Type | Description |
|------|------|-------------|
| id | string | The unique identifier of the payout |
| entity | string | The type of entity (payout) |
| fund_account_id | string | The ID of the fund account |
| fund_account | object | Details of the fund account |
| amount | integer | The amount of the payout (in paise) |
| currency | string | The currency of the payout |
| notes | object | Additional notes for the payout |
| fees | integer | The fees charged for the payout |
| tax | integer | The tax charged on the fees |
| status | string | The status of the payout |
| utr | string | The UTR (Unique Transaction Reference) number |
| mode | string | The mode of payout |
| purpose | string | The purpose of the payout |
| reference_id | string | The reference ID for the payout |
| narration | string | The narration for the payout |
| created_at | integer | The timestamp when the payout was created |