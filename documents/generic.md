## Generic Entity


### Method Signature
```go
body, _ := client.Entity.Do(url, method, data, headers)
```    

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| url*          | string      | The endpoint to which the request will be made. This should include the version information (e.g., "v1/contacts" or "v2/accounts")  |
| method*       | string      | The HTTP method for the request |
| data          | string      | The data to be sent with the request (e.g., 'GET', 'POST', 'PUT', 'PATCH', 'DELETE').|
| headers         | object      | The headers to be included in the request. |

-------------------------------------------------------------------------------------------------------

### Create a contacts

```go

data := map[string]interface{}{
  "name": "Gaurav Kumar",
  "email": "gaurav.kumar@example.com",
  "contact": "9123456789",
  "type": "employee",
  "reference_id":"Acme Contact ID 12345",
  "notes": map[string]interface{}{
    "notes_key_1":"Tea, Earl Grey, Hot",
    "notes_key_2":"Tea, Earl Grey… decaf.",
  },
}

body, _ := client.Entity.Do("v1/contacts", http.MethodPost, data, nil)
```

**Response:**

```json
{
  "id": "cont_00000000000001",
  "entity": "contact",
  "name": "Gaurav Kumar",
  "contact": "9123456789",
  "email": "gaurav.kumar@example.com",
  "type": "employee",
  "reference_id": "Acme Contact ID 12345",
  "batch_id": null,
  "active": true,
  "notes": {
    "notes_key_1": "Tea, Earl Grey, Hot",
    "notes_key_2": "Tea, Earl Grey… decaf."
  },
  "created_at": 1545320320
}
```

-------------------------------------------------------------------------------------------------------


**PN: * indicates mandatory fields**
<br>
<br>