package utils

import "project-app-bioskop-golang-homework-anas/internal/domain"

// CreateCreditCardDetails membuat payment details untuk credit card
func CreateCreditCardDetails(cardType, last4, bank string) domain.PaymentDetails {
	return domain.PaymentDetails{
		"type":      "credit_card",
		"card_type": cardType,
		"last4":     last4,
		"bank":      bank,
	}
}

// CreateEWalletDetails membuat payment details untuk e-wallet
func CreateEWalletDetails(provider, phone, transactionID string) domain.PaymentDetails {
	return domain.PaymentDetails{
		"type":           "e_wallet",
		"provider":       provider,
		"phone":          phone,
		"transaction_id": transactionID,
	}
}

// CreateBankTransferDetails membuat payment details untuk bank transfer
func CreateBankTransferDetails(bank, accountNumber, referenceNo string) domain.PaymentDetails {
	return domain.PaymentDetails{
		"type":           "bank_transfer",
		"bank":           bank,
		"account_number": accountNumber,
		"reference_no":   referenceNo,
	}
}
