## payment verification

```go
import (
utils "github.com/razorpay/razorpay-go/utils"
)
```

### Verify payment verification

```go
params := map[string]interface{}{
	"razorpay_order_id": "order_IEIaMR65cu6nz3",
	"razorpay_payment_id": "pay_IH4NVgf4Dreq1l",
}

signature := "0d4e745a1838664ad6c9c9902212a32d627d68e917290b0ad5f08ff4561bc50f";
secret := "EnLs21M47BllR3X8PSFtjtbd";
utils.VerifyPaymentSignature(params, signature, secret)
```

**Parameters:**


| Name  | Type      | Description                                      |
|-------|-----------|--------------------------------------------------|
| param.razorpay_order_id*  | string | The id of the order to be fetched  |
| param.razorpay_payment_id*    | string | The id of the payment to be fetched |
| signature* | string   | Signature returned by the Checkout. This is used to verify the payment. |
| secret* | string   | your api secret as secret |

-------------------------------------------------------------------------------------------------------
### Verify subscription verification

```go
params := map[string]interface{}{
	"razorpay_subscription_id": "sub_ID6MOhgkcoHj9I",
	"razorpay_payment_id": "pay_IDZNwZZFtnjyym",
}

signature := "601f383334975c714c91a7d97dd723eb56520318355863dcf3821c0d07a17693";
secret := "EnLs21M47BllR3X8PSFtjtbd";
utils.VerifySubscriptionSignature(params, signature, secret)
```

**Parameters:**


| Name  | Type      | Description                                      |
|-------|-----------|--------------------------------------------------|
| param.razorpay_subscription_id*  | string | The id of the subscription to be fetched  |
| param.razorpay_payment_id*    | string | The id of the payment to be fetched |
| signature* | string   | Signature returned by the Checkout. This is used to verify the payment. |
| secret* | string   | your api secret as secret |

-------------------------------------------------------------------------------------------------------
### Verify paymentlink verification

```go
params := map[string]interface{} {
	"payment_link_id": "plink_IH3cNucfVEgV68",
	"razorpay_payment_id": "pay_IH3d0ara9bSsjQ",
	"payment_link_reference_id": "TSsd1989",
	"payment_link_status": "paid",
}
signature := "07ae18789e35093e51d0a491eb9922646f3f82773547e5b0f67ee3f2d3bf7d5b";
secret := "EnLs21M47BllR3X8PSFtjtbd";
utils.VerifyPaymentLinkSignature(params, signature, secret)
```

**Parameters:**


| Name  | Type      | Description                                      |
|-------|-----------|--------------------------------------------------|
| param.payment_link_id*  | string | The id of the paymentlink to be fetched  |
| param.razorpay_payment_id*  | string | The id of the payment to be fetched  |
| param.payment_link_reference_id*  | string |  A reference number tagged to a Payment Link |
| param.payment_link_status*  | string | Current status of the link  |
| signature* | string   | The hash signature is calculated using HMAC with SHA256 algorithm; with your webhook secret set as the key and the webhook request body as the message.. |
| secret* | string   | your api secret as secret |

-------------------------------------------------------------------------------------------------------

### Verify webhook verification

```go
webhookbody := `{"entity":"event","account_id":"acc_HVFD0PFnHPAzKj","event":"payment.authorized","contains":["payment"],"payload":{"payment":{"entity":{"id":"pay_JUEM4c0pSLpFEW","entity":"payment","amount":12300,"currency":"INR","status":"authorized","order_id":"order_JUELuT6cFaHkvd","invoice_id":null,"international":false,"method":"netbanking","amount_refunded":0,"refund_status":null,"captured":false,"description":"#JUELZ1z1EC0pwi","card_id":null,"bank":"SBIN","wallet":null,"vpa":null,"email":"gaurav.kumar@example.com","contact":"+919999999999","notes":[],"fee":null,"tax":null,"error_code":null,"error_description":null,"error_source":null,"error_step":null,"error_reason":null,"acquirer_data":{"bank_transaction_id":"6416615"},"created_at":1652339804}}},"created_at":1652339806}`;

    signature := "56df260ab4647b07b729b0b48d2a95b71de42f109884949e5ec387a2bb8b6919";
    webhook_secret := "123456";
	body := utils.VerifyWebhookSignature(webhookbody, signature, webhook_secret)
```

**Parameters:**


| Name  | Type      | Description                                      |
|-------|-----------|--------------------------------------------------|
| webhookbody*  | string | raw webhook request body  |
| signature*  | string | The hash signature is calculated using HMAC with SHA256 algorithm; with your webhook secret set as the key and the webhook request body as the message.  |
| webhook_secret* | string   | your webhook secret |

-------------------------------------------------------------------------------------------------------


**PN: * indicates mandatory fields**
<br>
<br>