## Items

### Create an item

```go
data:= map[string]interface{}{
    "name": "Book / English August",
    "description": "An indian story, Booker prize winner.",
    "amount": 20000,
    "currency": "INR",
}
body, err := client.Item.Create(data, nil)
```

**Parameters:**

| Name            | Type    | Description                                                                  |
|-----------------|---------|------------------------------------------------------------------------------|
| name*          | string | Name of the item.                    |
| description        | string  | A brief description of the item.  |
| amount*         | integer  | Price of the item     |
| currency*           | string  | The currency in which the amount should be charged.    |

**Response:**
```json
{
  "id": "item_7Oxp4hmm6T4SCn",
  "active": true,
  "name": "Book / English August",
  "description": "An indian story, Booker prize winner.",
  "amount": 20000,
  "currency": "INR"
}
```

-------------------------------------------------------------------------------------------------------

### Fetch all items

```go
options:= map[string]interface{}{
    "count": 3,
    "skip": 1,
}
body, err := client.Item.All(options, nil)
```
**Parameters:**

| Name  | Type      | Description                                      |
|-------|-----------|--------------------------------------------------|
| from  | timestamp | timestamp after which the item were created  |
| to    | timestamp | timestamp before which the item were created |
| count | integer   | number of item to fetch (default: 10)        |
| skip  | integer   | number of item to be skipped (default: 0)    |
| active   | boolean  | Possible values is `0` or `1` |

**Response:**
```json
{
  "entity": "collection",
  "count": 3,
  "items": [
    {
      "id": "item_7Oy8OMV6BdEAac",
      "active": true,
      "name": "Book / Ignited Minds",
      "description": null,
      "amount": 15000,
      "currency": "INR"
    },
    {
      "id": "item_7Oxp4hmm6T4SCn",
      "active": true,
      "name": "Book / English August",
      "description": "An indian story, Booker prize winner.",
      "amount": 20000,
      "currency": "INR"
    },
    {
      "id": "item_7OxoGnoxCuUKbo",
      "active": true,
      "name": "Book / English August",
      "description": null,
      "amount": 20000,
      "currency": "INR"
    }
  ]
}
```
-------------------------------------------------------------------------------------------------------
### Fetch particular item

```go
body, err := client.Item.Fetch("<itemId>", nil, nil)
```
**Parameters**

| Name     | Type   | Description                         |
|----------|--------|-------------------------------------|
| itemId* | string | The id of the item to be fetched |

**Response:**
```json
{
  "id": "item_7Oxp4hmm6T4SCn",
  "active": true,
  "name": "Book / English August",
  "description": "An indian story, Booker prize winner.",
  "amount": 20000,
  "currency": "INR"
}
```

-------------------------------------------------------------------------------------------------------

### Update an item

```go
data:= map[string]interface{}{
    "name": "Book / Ignited Minds - Updated name!",
    "description": "New descirption too.",
    "amount": 20000,
    "currency": "INR",
    "active": true,
}
body, err := client.Item.Update("<itemId>", data, nil)
```
**Parameters**

| Name     | Type   | Description                         |
|----------|--------|-------------------------------------|
| itemId* | string | The id of the item to be fetched |
| name       | string | Name of the item.                    |
| description  | string  | A brief description of the item.  |
| amount         | integer  | Amount of the item to be paid     |
| currency           | string  | The currency in which the amount should be charged.    |
| active   | boolean  | Possible values is `0` or `1` |

**Response:**
```json
{
  "id": "item_7Oy8OMV6BdEAac",
  "active": true,
  "name": "Book / Ignited Minds - Updated name!",
  "description": "New descirption too. :)",
  "amount": 15000,
  "currency": "INR"
}
```
-------------------------------------------------------------------------------------------------------
### Delete item

```go
~~body, err := client.Item.Delete("<itemId>", nil, nil)~~
```
**Parameters**

| Name     | Type   | Description                         |
|----------|--------|-------------------------------------|
| itemId* | string | The id of the item to be fetched |

**Response:**
```json
[]
```
-------------------------------------------------------------------------------------------------------

**PN: * indicates mandatory fields**
<br>
<br>
**For reference click [here](https://razorpay.com/docs/api/items)**