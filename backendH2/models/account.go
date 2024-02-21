package models

type Account interface {
	WithdrawMoney(amount float64) *TransactionResult
	DepositMoney(amount float64) *TransactionResult
	GetId() int
}
