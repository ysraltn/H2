package models

type CorporateAccount struct {
	ID      int
	Name    string
	Balance float64
	Fee     float64
}

func CorporateAccountInit(id int, name string, balance float64) *CorporateAccount {
	return &CorporateAccount{ID: id, Name: name, Balance: balance, Fee: 4.0}
}

func (corporateAccount *CorporateAccount) GetId() int{
	return corporateAccount.ID
}

func (corporateAccount *CorporateAccount) WithdrawMoney(amount float64) *TransactionResult {
	if amount+corporateAccount.Fee <= corporateAccount.Balance {
		corporateAccount.Balance -= (amount + corporateAccount.Fee)
		transectionResult := TransactionResultInit(true, "withdraw completed successfully, balance:", corporateAccount.Balance)
		return transectionResult
	} else {
		transectionResult := TransactionResultInit(false, "not enough money in your balance, balance:", corporateAccount.Balance)
		return transectionResult
	}
}

func (corporateAccount *CorporateAccount) DepositMoney(amount float64) *TransactionResult {
	if amount > corporateAccount.Fee {
		corporateAccount.Balance += (amount - corporateAccount.Fee)
		transectionResult := TransactionResultInit(true, "deposit completed successfully, balance:", corporateAccount.Balance)
		return transectionResult
	} else {
		transectionResult := TransactionResultInit(false, "amount is lower than transaction fee, balance:", corporateAccount.Balance)
		return transectionResult
	}

}
