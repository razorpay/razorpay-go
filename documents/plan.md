## Plans

### Create plan

```go
data:= map[string]interface{}{
  "period": "weekly",
  "interval": 1,
  "item": map[string]interface{}{
    "name": "Test plan - Weekly",
    "amount": 69900,
    "currency": "INR",
    "description": "Description for the test plan",
  },
  "notes": map[string]interface{}{
    "notes_key_1": "Tea, Earl Grey, Hot",
    "notes_key_2": "Tea, Earl Grey… decaf.",
  },
}
body, err := client.Plan.Create(data, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
| period*          | string | Used together with `interval` to define how often the customer should be charged.Possible values:<br>1.`daily` <br>2.`weekly`<br>3.`monthly` <br>4.`yearly`  |
| interval*          | string | Used together with `period` to define how often the customer should be charged  |
| item*          | object | Details of the plan. For more details please refer [here](https://razorpay.com/docs/api/subscriptions/#create-a-plan) |
| notes          | object | Notes you can enter for the contact for future reference.   |

**Response:**
```json
{
  "id":"plan_00000000000001",
  "entity":"plan",
  "interval":1,
  "period":"weekly",
  "item":{
    "id":"item_00000000000001",
    "active":true,
    "name":"Test plan - Weekly",
    "description":"Description for the test plan - Weekly",
    "amount":69900,
    "unit_amount":69900,
    "currency":"INR",
    "type":"plan",
    "unit":null,
    "tax_inclusive":false,
    "hsn_code":null,
    "sac_code":null,
    "tax_rate":null,
    "tax_id":null,
    "tax_group_id":null,
    "created_at":1580219935,
    "updated_at":1580219935
  },
  "notes":{
    "notes_key_1":"Tea, Earl Grey, Hot",
    "notes_key_2":"Tea, Earl Grey… decaf."
  },
  "created_at":1580219935
}
```
-------------------------------------------------------------------------------------------------------

### Fetch all plans

```go
options:= map[string]interface{}{
  "count": 1,
  "skip": 1,
}
body, err := client.Plan.All(options, nil)
```

**Parameters:**

| Name  | Type      | Description                                      |
|-------|-----------|--------------------------------------------------|
| from  | timestamp | timestamp after which the plans were created  |
| to    | timestamp | timestamp before which the plans were created |
| count | integer   | number of plans to fetch (default: 10)        |
| skip  | integer   | number of plans to be skipped (default: 0)    |

**Response:**
```json
{
  "entity": "collection",
  "count": 1,
  "items": [
    {
      "id": "plan_00000000000001",
      "entity": "plan",
      "interval": 1,
      "period": "weekly",
      "item": {
        "id": "item_00000000000001",
        "active": true,
        "name": "Test plan - Weekly",
        "description": "Description for the test plan - Weekly",
        "amount": 69900,
        "unit_amount": 69900,
        "currency": "INR",
        "type": "plan",
        "unit": null,
        "tax_inclusive": false,
        "hsn_code": null,
        "sac_code": null,
        "tax_rate": null,
        "tax_id": null,
        "tax_group_id": null,
        "created_at": 1580220492,
        "updated_at": 1580220492
      },
      "notes": {
        "notes_key_1": "Tea, Earl Grey, Hot",
        "notes_key_2": "Tea, Earl Grey… decaf."
      },
      "created_at": 1580220492
    }
  ]
}
```
-------------------------------------------------------------------------------------------------------

### Fetch particular plan

```go
body, err := client.Plan.Fetch("<planId>", nil, nil)
```

**Parameters:**

| Name  | Type      | Description                                      |
|-------|-----------|--------------------------------------------------|
| planId  | string | The id of the plan to be fetched  |

**Response:**
```json
{
  "id":"plan_00000000000001",
  "entity":"plan",
  "interval":1,
  "period":"weekly",
  "item":{
    "id":"item_00000000000001",
    "active":true,
    "name":"Test plan - Weekly",
    "description":"Description for the test plan - Weekly",
    "amount":69900,
    "unit_amount":69900,
    "currency":"INR",
    "type":"plan",
    "unit":null,
    "tax_inclusive":false,
    "hsn_code":null,
    "sac_code":null,
    "tax_rate":null,
    "tax_id":null,
    "tax_group_id":null,
    "created_at":1580220492,
    "updated_at":1580220492
  },
  "notes":{
    "notes_key_1":"Tea, Earl Grey, Hot",
    "notes_key_2":"Tea, Earl Grey… decaf."
  },
  "created_at":1580220492
}
```
-------------------------------------------------------------------------------------------------------

**PN: * indicates mandatory fields**
<br>
<br>
**For reference click [here](https://razorpay.com/docs/api/subscriptions/#plans)**