package main

import (
	"encoding/json"
	"fmt"
)

type AccountDetail struct {
	id          string
	accountType string
}

type Account struct {
	detail      *AccountDetail
	CustoerName string
}

func (account *Account) setDetail(id string, accountType string) {
	account.detail = &AccountDetail{id: id, accountType: accountType}
}

func (account *Account) getId() string {
	return account.detail.id
}

func (account *Account) getAccountType() string {
	return account.detail.accountType
}

func main() {
	account := Account{CustoerName: "Bimal Sharma"}
	account.setDetail("23456", "saving account")
	jsonAccount, _ := json.Marshal(account)
	fmt.Println("Account info : ", string(jsonAccount))
	fmt.Println("Account Id", account.getId())
	fmt.Println("Account Type", account.getAccountType())

}
