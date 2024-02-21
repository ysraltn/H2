package models

type BankService struct {
	Accounts         []Account
	nextCorporateId  int
	nextIndividualId int
}

func BankServiceInit() *BankService {
	return &BankService{Accounts: make([]Account, 0), nextCorporateId: 110001, nextIndividualId: 220001}
}

func (bankService *BankService) GetCorporateId() int {
	corporateId := bankService.nextCorporateId
	bankService.nextCorporateId += 1
	return corporateId
}

func (bankService *BankService) GetIndividualId() int {
	individualId := bankService.nextIndividualId
	bankService.nextIndividualId += 1
	return individualId
}

func (bankService *BankService) AppendAccountToList(account Account) {
	bankService.Accounts = append(bankService.Accounts, account)

}

func (bankService *BankService) WithdrawService(account Account, amount float64) *TransactionResult {
	switch acc := account.(type) {
	case *CorporateAccount:
		transactionResult := acc.WithdrawMoney(amount)
		return transactionResult
	case *IndividualAccount:
		transactionResult := acc.WithdrawMoney(amount)
		return transactionResult
	default:
		transactionResult := TransactionResultInit(false, "invalid account", 0)
		return transactionResult
	}

}

func (bankService *BankService) DepositService(account Account, amount float64) *TransactionResult {
	return account.DepositMoney(amount)
	/* switch acc := account.(type) {
	case *CorporateAccount:
		transactionResult := acc.DepositMoney(amount)
		return transactionResult
	case *IndividualAccount:
		transactionResult := account.DepositMoney(amount)
		return transactionResult
	default:
		transactionResult := TransactionResultInit(false, "invalid account", 0)
		return transactionResult
	} */

}
