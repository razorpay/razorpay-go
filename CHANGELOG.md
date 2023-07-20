# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.2.0] - 2023-07-20

### Added
feat: Added new API endpoints

- Added account onboarding API Create, Fetch, Edit, Delete
- Added stakeholders API Create, Fetch, All, Edit
- Added product configuration API RequestProductConfiguration, Fetch, Edit, FetchTnc
- Added webhooks API Create, Fetch, All, Edit, Delete
- Added Documents API UploadAccountDoc, FetchAccountDoc, UploadStakeholderDoc UploadStakeholderDoc
- Added token sharing API Create, FetchCardPropertiesByToken, DeleteToken, ProcessPaymentOnAlternatePAorPG

## [1.1.0] - 2022-07-05

### Added
- Added Third party validation & Otp API for Payment (CreateUpi, ValidateVpa, OtpGenerate, OtpSubmit, OtpResend)
- Update Documention

## [1.0.0] - 2022-04-29

### Added

- QR code end point API
- Settlement end point API
- Fund Account end point API
- Item end point API
- New APIs for Invoices (Edit, Cancel, Issue, Delete, Send/resend)
- New API for Customers (Fetch all customers)
- New APIs for PaymentLinks (Cancel, Update, Send/Resend)
- New APIs for Subscriptions (Update, Pause, Resume, Cancel, Pending update, Delete offer)
- New API for Addons (Fetch all Addons)
- New API for Refund (Update refund)
- New APIs for Payments (Update, Create recurring, Create Json, Payment downtime details, refunds of a payment)
- New APIs for Virtual Account (Add receiver, add an allowed payer account, delete an allowed payer account)
- Updated Testcases
- Updated Readme 

### Changed

- Request method of update order from `put` to `patch`
- Request method of update customer from `post` to `put`

## 0.1

### Added

- Initial Release
