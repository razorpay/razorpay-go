## Customer Fund Account

### Create a fund account
```go
data:= map[string]interface{}{
  "customer_id":"cust_Aa000000000001",
  "account_type":"bank_account",
  "bank_account":map[string]interface{}{
    "name":"Gaurav Kumar",
    "account_number":"11214311215411",
    "ifsc":"HDFC0000053",
  },
}
body, err := client.FundAccount.Create(data, nil)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| customer_id*   | string      | The id of the customer to be fetched  |
| account_type* | string      | The bank_account to be linked to the customer ID  |
| bank_account* | object      | All keys listed [here](https://razorpay.com/docs/payments/customers/customer-fund-account-api/#create-a-fund-account) are supported |


**Response:**
```json
{
  "id":"fa_Aa00000000001",
  "entity":"fund_account",
  "customer_id":"cust_Aa000000000001",
  "account_type":"bank_account",
  "bank_account":{
    "name":"Gaurav Kumar",
    "account_number":"11214311215411",
    "ifsc":"HDFC0000053",
    "bank_name":"HDFC Bank"
  },
  "active":true,
  "created_at":1543650891
}
```
-------------------------------------------------------------------------------------------------------

### Fetch all fund accounts

```go
data := map[string]interface{}{
    "customer_id":"cust_Aa000000000001",
}
body, err := client.FundAccount.All(data, nil)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| customer_id*   | string      | The id of the customer to be fetched  |

**Response:**
```json
{
  "entity":"collection",
  "count":1,
  "items":[
    {
      "id":"fa_Aa000000000001",
      "entity":"fund_account",
      "customer_id":"cust_Aa000000000001",
      "account_type":"bank_account",
      "bank_account":{
        "name":"Gaurav Kumar",
        "account_number":"11214311215411",
        "ifsc":"HDFC0000053",
        "bank_name":"HDFC Bank"
      },
      "active":true,
      "created_at":1543650891
    }
  ]
}
```
-------------------------------------------------------------------------------------------------------

**PN: * indicates mandatory fields**
<br>
<br>
**For reference click [here](https://razorpay.com/docs/payments/customers/customer-fund-account-api/)**