# Orders

Params Types:

- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#NotesUnionParam">NotesUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#NotesUnion">NotesUnion</a>
- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Order">Order</a>
- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#OrderListResponse">OrderListResponse</a>
- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#OrderListPaymentsResponse">OrderListPaymentsResponse</a>

Methods:

- <code title="post /orders">client.Orders.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#OrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#OrderNewParams">OrderNewParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /orders/{id}">client.Orders.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#OrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /orders/{id}">client.Orders.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#OrderService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#OrderUpdateParams">OrderUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /orders">client.Orders.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#OrderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#OrderListParams">OrderListParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#OrderListResponse">OrderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /orders/{id}/payments">client.Orders.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#OrderService.ListPayments">ListPayments</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#OrderListPaymentsParams">OrderListPaymentsParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#OrderListPaymentsResponse">OrderListPaymentsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Customers

Response Types:

- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Customer">Customer</a>
- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#CustomerListResponse">CustomerListResponse</a>

Methods:

- <code title="post /customers">client.Customers.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#CustomerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#CustomerNewParams">CustomerNewParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customers/{id}">client.Customers.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#CustomerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /customers/{id}">client.Customers.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#CustomerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#CustomerUpdateParams">CustomerUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customers">client.Customers.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#CustomerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#CustomerListParams">CustomerListParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#CustomerListResponse">CustomerListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Payments

Response Types:

- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#AcquirerData">AcquirerData</a>
- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Card">Card</a>
- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Payment">Payment</a>
- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Refund">Refund</a>
- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentListResponse">PaymentListResponse</a>

Methods:

- <code title="get /payments/{id}">client.Payments.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentGetParams">PaymentGetParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Payment">Payment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /payments/{id}">client.Payments.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentUpdateParams">PaymentUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Payment">Payment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments">client.Payments.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentListParams">PaymentListParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentListResponse">PaymentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payments/{id}/capture">client.Payments.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentService.Capture">Capture</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentCaptureParams">PaymentCaptureParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Payment">Payment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CardDetails

Methods:

- <code title="get /payments/{id}/card">client.Payments.CardDetails.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentCardDetailService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Card">Card</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## QrCodes

Response Types:

- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#QrCode">QrCode</a>
- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentQrCodeListResponse">PaymentQrCodeListResponse</a>

Methods:

- <code title="post /payments/qr_codes">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentQrCodeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentQrCodeNewParams">PaymentQrCodeNewParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#QrCode">QrCode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/qr_codes/{id}">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentQrCodeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentQrCodeGetParams">PaymentQrCodeGetParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#QrCode">QrCode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/qr_codes">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentQrCodeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentQrCodeListParams">PaymentQrCodeListParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentQrCodeListResponse">PaymentQrCodeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payments/qr_codes/{id}/close">client.Payments.QrCodes.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentQrCodeService.Close">Close</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#QrCode">QrCode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Refunds

Response Types:

- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#RefundList">RefundList</a>

Methods:

- <code title="post /payments/{id}/refund">client.Payments.Refunds.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentRefundService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentRefundNewParams">PaymentRefundNewParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/{id}/refunds/{refund_id}">client.Payments.Refunds.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentRefundService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, refundID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentRefundGetParams">PaymentRefundGetParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/{id}/refunds">client.Payments.Refunds.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentRefundService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentRefundListParams">PaymentRefundListParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#RefundList">RefundList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PaymentLinks

Response Types:

- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentLink">PaymentLink</a>
- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentLinkListResponse">PaymentLinkListResponse</a>
- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentLinkNotifyResponse">PaymentLinkNotifyResponse</a>

Methods:

- <code title="post /payment_links">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentLinkService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentLinkNewParams">PaymentLinkNewParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentLink">PaymentLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payment_links/{id}">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentLinkService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentLink">PaymentLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /payment_links/{id}">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentLinkService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentLinkUpdateParams">PaymentLinkUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentLink">PaymentLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payment_links">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentLinkService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentLinkListParams">PaymentLinkListParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentLinkListResponse">PaymentLinkListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payment_links/{id}/notify_by/{medium}">client.PaymentLinks.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentLinkService.Notify">Notify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, medium <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentLinkNotifyParamsMedium">PaymentLinkNotifyParamsMedium</a>, body <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentLinkNotifyParams">PaymentLinkNotifyParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PaymentLinkNotifyResponse">PaymentLinkNotifyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Refunds

Methods:

- <code title="get /refunds/{id}">client.Refunds.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#RefundService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /refunds">client.Refunds.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#RefundService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#RefundListParams">RefundListParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#RefundList">RefundList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /refunds/{id}">client.Refunds.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#RefundService.UpdateNotes">UpdateNotes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#RefundUpdateNotesParams">RefundUpdateNotesParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Refund">Refund</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Settlements

Response Types:

- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Settlement">Settlement</a>
- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#SettlementListResponse">SettlementListResponse</a>

Methods:

- <code title="get /settlements/{id}">client.Settlements.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#SettlementService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Settlement">Settlement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /settlements">client.Settlements.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#SettlementService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#SettlementListParams">SettlementListParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#SettlementListResponse">SettlementListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Recon

Response Types:

- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#SettlementReconGetResponse">SettlementReconGetResponse</a>

Methods:

- <code title="get /settlements/recon/combined">client.Settlements.Recon.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#SettlementReconService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#SettlementReconGetParams">SettlementReconGetParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#SettlementReconGetResponse">SettlementReconGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Ondemand

Response Types:

- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#SettlementOndemand">SettlementOndemand</a>

Methods:

- <code title="post /settlements/ondemand">client.Settlements.Ondemand.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#SettlementOndemandService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#SettlementOndemandNewParams">SettlementOndemandNewParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#SettlementOndemand">SettlementOndemand</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /settlements/ondemand/{id}">client.Settlements.Ondemand.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#SettlementOndemandService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#SettlementOndemand">SettlementOndemand</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Payouts

Response Types:

- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Payout">Payout</a>
- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PayoutListResponse">PayoutListResponse</a>

Methods:

- <code title="get /payouts/{id}">client.Payouts.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PayoutService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Payout">Payout</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payouts">client.Payouts.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PayoutService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PayoutListParams">PayoutListParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PayoutListResponse">PayoutListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Plans

Response Types:

- <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Plan">Plan</a>

Methods:

- <code title="post /plans">client.Plans.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PlanService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#PlanNewParams">PlanNewParams</a>) (<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2">gotue</a>.<a href="https://pkg.go.dev/github.com/razorpay/razorpay-go/v2#Plan">Plan</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
