# Razorpay Go Client

Golang bindings for interacting with the Razorpay API

This is primarily meant for merchants who wish to perform interactions with the Razorpay API programatically

Read up here for getting started and understanding the payment flow with Razorpay: <https://docs.razorpay.com/docs/getting-started>

## Documentation

Documentation of Razorpay's API and their usage is available at <https://docs.razorpay.com>

## Usage
You need to setup your key and secret using the following:
You can find your API keys at <https://dashboard.razorpay.com/#/app/keys>.

```go
import (
razorpay "github.com/razorpay/razorpay-go"
)

client := razorpay.NewClient("<YOUR_API_KEY>", "<YOUR_API_SECRET>")

```

Note: All methods below return a `map[string]interface{}` and `error`. All methods accept an `extraHeaders` param of type `map[string]string`, allowing you to optinally set extra HTTP headers on the request.

## Supported Resources

- [Account](documents/account.md)
- [Customer](documents/customers.md)
- [Token](documents/token.md)
- [Order](documents/order.md)
- [Payments](documents/payment.md)
- [Product Configuration](documents/productConfiguration.md)
- [Settlements](documents/settlement.md)
- [Stakeholder](documents/stakeholder.md)
- [Fund](documents/fundAccount.md)
- [Refunds](documents/refund.md)
- [Invoice](documents/invoice.md)
- [Item](documents/item.md)
- [Plan](documents/plan.md)
- [PaymentVerfication](documents/paymentVerification.md)
- [Subscriptions](documents/subscription.md)
- [Add-on](documents/addon.md)
- [Payment Links](documents/paymentLink.md)
- [Smart Collect](documents/virtualAccount.md)
- [Route](documents/transfer.md)
- [QR Code](documents/qrcode.md)
- [Emandate](documents/emandate.md)
- [Cards](documents/card.md)
- [Paper NACH](documents/papernach.md)
- [UPI](documents/upi.md)
- [Register Emandate and Charge First Payment Together](documents/registerEmandate.md)
- [Register NACH and Charge First Payment Together](documents/registerNach.md)
- [Token](documents/token.md)
- [Webhook](documents/webhook.md)

## License

The Razorpay Go SDK is released under the MIT License. See [LICENSE](LICENSE) file for more details.
