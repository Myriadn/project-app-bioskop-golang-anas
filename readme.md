# System Ticket Bioskop

Aplikasi Cinema Booking System berbasis RESTful API menggunakan bahasa pemrograman Golang.

Aplikasi ini bertujuan untuk membantu pengguna (customer) dalam melakukan pendaftaran akun, login, memilih bioskop, mengecek ketersediaan kursi, melakukan pemesanan kursi, serta melakukan pembayaran tiket.

## Payment Doc

1. Credit Card

`{
    "booking_id": ,
    "payment_method": "CREDIT_CARD",
    "payment_details": {
        "card_type": "Visa",
        "last4": "1234",
        "bank": "BCA",
        "cardholder": "John Doe"
    }
}`

2. E-Wallet

`{
    "booking_id": ,
    "payment_method": "GOPAY",
    "payment_details": {
        "provider": "GoPay",
        "phone": "081234567890",
        "transaction_id": "GP2026011112345",
        "account_name": "John Doe"
    }
}`

3. Bank Transfer

`{
    "booking_id": ,
    "payment_method": "BANK_TRANSFER",
    "payment_details": {
        "bank": "BCA",
        "account_number": "1234567890",
        "account_name": "John Doe",
        "reference_number": "TRF2026011112345",
        "transfer_date": "2026-01-11"
    }
}`

4. Minimal

`{
    "booking_id": ,
    "payment_method": "CREDIT_CARD",
    "payment_details": {
        "note": "Payment via Credit Card"
    }
}`

5. Empty

`{
  "booking_id": 5,
  "payment_method": "GOPAY",
  "payment_details": {}
}`

## Template for pay

`{
    "booking_id": <GANTI_DENGAN_BOOKING_ID>,
    "payment_method": "<PILIH: CREDIT_CARD|DEBIT_CARD|GOPAY|OVO|DANA|SHOPEEPAY|BANK_TRANSFER>",
    "payment_details": {
        "key1": "value1",
        "key2": "value2",
        "key3": "value3"
    }
}`
