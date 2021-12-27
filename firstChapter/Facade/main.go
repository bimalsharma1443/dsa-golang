package main

import "fmt"

type Account struct {
	id          string
	accountType string
}

func (account *Account) create(accountType string) {
	fmt.Println("creating an account")
	account.accountType = accountType
}

func (account *Account) getById(id string) *Account {
	fmt.Println("get account by id")
	return account
}

func (account *Account) deleteById(id string) {
	fmt.Println("delete account by id")
}

type Customer struct {
	name string
	id   int
}

func (customer *Customer) create(name string) *Customer {
	fmt.Println("creating a customer")
	customer.name = name
	return customer
}

type Transaction struct {
	id            string
	amount        float32
	srcAccountId  string
	destAccountId string
}

func (transaction *Transaction) create(srcAccountId, destAccountId string, amount float32) {
	fmt.Println("creating a transaction")
	transaction.srcAccountId = srcAccountId
	transaction.destAccountId = destAccountId
	transaction.amount = amount
}

type BranchManagerFacade struct {
	account     *Account
	customer    *Customer
	transaction *Transaction
}

func NewBranchManagerFacade() *BranchManagerFacade {
	fmt.Println("new branch manager Facade")
	return &BranchManagerFacade{&Account{}, &Customer{}, &Transaction{}}
}

func (facade *BranchManagerFacade) createCustomerAccount(customerName string, accountType string) (*Customer, *Account) {
	facade.customer.create(customerName)
	facade.account.create(accountType)
	return facade.customer, facade.account
}
func (facade *BranchManagerFacade) createTransaction(srcAccountId, destAccountId string, amount float32) *Transaction {
	facade.transaction.create(srcAccountId, destAccountId, amount)
	return facade.transaction

}

func main() {
	var facade = NewBranchManagerFacade()
	var customer, account = facade.createCustomerAccount("John", "saving")
	fmt.Println("Customer name is :", customer.name)
	fmt.Println("account type is :", account.accountType)
	var transaction = facade.createTransaction("21345", "98646", 100)
	fmt.Println("Transaction Amount is", transaction.amount)
}
