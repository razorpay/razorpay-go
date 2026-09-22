# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed
- Base API URL is now read from the `RAZORPAY_BASE_URL` environment variable, falling back to `https://api.razorpay.com` when unset
- Renamed `constants.BASE_URL` to `constants.DEFAULT_BASE_URL`; use `constants.GetBaseURL()` to resolve the effective base URL

## [1.4.1] - 2026-04-23

### Changed
- Upgraded Go version from 1.14 to 1.26
- Updated dependencies (testify v1.5.1 → v1.11.1)
- Migrated from deprecated `ioutil` to `io` package

### Security
- Fixed timing attack vulnerability in signature verification using `hmac.Equal()` for constant-time comparison

### Added
- OAuth token support
- Unit tests for client initialization and configuration
- Unit tests for error handling across resources
- Unit tests for signature verification edge cases

### Fixed
- Boolean corrections in documentation

## [1.4.0] - 2025-06-15

### Changed
- Removed global request variable for better concurrency support

### Fixed
- Initialise headers for request

## [1.3.4] - 2025-05-07

### Added
- Unit tests for `buildURLWithParams` function in requests package

### Fixed
- Query parameter handling in URLs
- Custom User Agent in headers

## [1.3.3] - 2025-05-02
feat: Added new API endpoints
- Added Payout API:
  - Fetch all payouts
  - Fetch specific payout by ID
  
## [1.3.2] - 2024-04-17
feat: Added new API endpoints

Added support for AddBankAccount, DeleteBankAccount, RequestEligibilityCheck & FetchEligibility on customer
Added support for Dispute
Added support for Document
Added support for fetch all IINs Supporting native otps & fetch all IINs with business sub-type using `All`
Added support for ViewRtoReview & EditFulfillment on order

## [1.3.1] - 2023-12-19
Rollback: Generic access point due to some performance concern

## [1.3.0] - 2023-12-14
feat: Added generic access point

## [1.2.0] - 2023-07-20

### Added
feat: Added new API endpoints

- Added account on boarding API Create, Fetch, Edit, Delete
- Added stakeholders API Create, Fetch, All, Edit
- Added product configuration API RequestProductConfiguration, Fetch, Edit, FetchTnc
- Added webhooks API Create, Fetch, All, Edit, Delete
- Added Documents API UploadAccountDoc, FetchAccountDoc, UploadStakeholderDoc UploadStakeholderDoc
- Added token sharing API Create, FetchCardPropertiesByToken, DeleteToken, ProcessPaymentOnAlternatePAorPG

## [1.1.0] - 2022-07-05

### Added
- Added Third party validation & Otp API for Payment (CreateUpi, ValidateVpa, OtpGenerate, OtpSubmit, OtpResend)
- Update Documentation

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
