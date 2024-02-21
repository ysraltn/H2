package main

import (
	"backendH2/models"
	"fmt"
)

func main() {
	bankService := models.BankServiceInit()

	//test icin
	// account1 := models.CorporateAccountInit(bankService.GetCorporateId(), "msi", 100)
	// account2 := models.IndividualAccountInit(bankService.GetIndividualId(), "veli", 200)
	// fmt.Println(bankService.DepositService(account1, 1))
	// fmt.Println(bankService.WithdrawService(account1, 500))
	// fmt.Println(bankService.DepositService(account2, 1))
	// fmt.Println(bankService.WithdrawService(account2, 500))
	// fmt.Println(account1)
	// fmt.Println(account2)

	account3 := models.CorporateAccountInit(bankService.GetCorporateId(), "htc", 100)
	fmt.Println(account3)
	var input int
	for {
		fmt.Println("1. cikis\n2. hesap ekle\n3. para yatir\n4. para cek\n5. hesaplari listele")
		fmt.Scanln(&input)
		switch input {
		case 1:
			return
		case 2:
			var name string
			var balance float64
			var accountType int
			fmt.Println("hesap ismi: ")
			fmt.Scanln(&name)
			fmt.Println("bakiye: ")
			fmt.Scanln(&balance)
			fmt.Println("hesap turu (1: corporate, 2: individual): ")
			fmt.Scanln(&accountType)
			switch accountType {
			case 1:
				account := models.CorporateAccountInit(bankService.GetCorporateId(), name, balance)
				bankService.AppendAccountToList(account)
				fmt.Println("hesabiniz olusturuldu:", account)
			case 2:
				account := models.IndividualAccountInit(bankService.GetIndividualId(), name, balance)
				bankService.AppendAccountToList(account)
				fmt.Println("hesabiniz olusturuldu:", account)
			default:
				fmt.Println("gecersiz secim")
			}
		case 3:
			var id int
			var amount float64
			fmt.Println("hesap id: ")
			fmt.Scanln(&id)
			fmt.Println("yatirilacak tutar: ")
			fmt.Scanln(&amount)
			found := false
			for _, account := range bankService.Accounts {
				if account.GetId() == id {
					fmt.Println(account.DepositMoney(amount))
					found = true
					break
				}
			}
			if !found {
				fmt.Println("hesap id gecersiz")
			}
		case 4:
			var id int
			var amount float64
			fmt.Println("hesap id: ")
			fmt.Scanln(&id)
			fmt.Println("cekilecek tutar: ")
			fmt.Scanln(&amount)
			for _, account := range bankService.Accounts {
				if account.GetId() == id {
					fmt.Println(account.WithdrawMoney(amount))
				} else {
					fmt.Println("hesap id gecersiz")
				}
			}
		case 5:
			for _, account := range bankService.Accounts {
				fmt.Println(account)
			}
		}
	}
}
