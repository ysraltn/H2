package models

type TransactionResult struct {
	Success bool
	Message string
	Balance float64
}

func TransactionResultInit(success bool, message string, balance float64) *TransactionResult {
	return &TransactionResult{Success: success, Message: message, Balance: balance}
}
