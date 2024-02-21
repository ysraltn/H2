package models

type IndividualAccount struct {
	ID      int
	Name    string
	Balance float64
	Fee     float64
}

func IndividualAccountInit(id int, name string, balance float64) *IndividualAccount {
	return &IndividualAccount{ID: id, Name: name, Balance: balance, Fee: 2.0}
}

func (individualAccount *IndividualAccount) GetId() int {
	return individualAccount.ID
}

func (individualAccount *IndividualAccount) WithdrawMoney(amount float64) *TransactionResult {
	if amount+2 <= individualAccount.Balance {
		individualAccount.Balance -= amount + 2
		transectionResult := TransactionResultInit(true, "withdraw completed successfully, balance:", individualAccount.Balance)
		return transectionResult
	} else {
		transectionResult := TransactionResultInit(false, "not enough money in your balance, balance:", individualAccount.Balance)
		return transectionResult
	}
}

func (individualAccount *IndividualAccount) DepositMoney(amount float64) *TransactionResult {
	if amount > 2 {
		individualAccount.Balance += amount - 2
		transectionResult := TransactionResultInit(true, "deposit completed successfully, balance:", individualAccount.Balance)
		return transectionResult
	} else {
		transectionResult := TransactionResultInit(false, "amount is lower than transaction fee, balance:", individualAccount.Balance)
		return transectionResult
	}

}
