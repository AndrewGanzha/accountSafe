package main

import (
	"fmt"
	"passwordKeep/account"
	"passwordKeep/files"
	"passwordKeep/output"
	"strings"
)

var menu = map[string]func(db *account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

func main() {
	showMenu()
}

func showMenu() {
	fmt.Println("Добро пожаловать в программу гененрации и получения паролей!")

	vault := account.NewVault(files.NewJsonDb("data.json"))
	PropmtData(vault)
}

func findAccountByUrl(vault *account.VaultWithDb) {
	findsAccount, err := vault.SearchAccountByUrl(askField("URL"), func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})

	if err != nil {
		output.PrintError(err.Error())
	}

	for _, findAccount := range findsAccount {
		findAccount.Output()
	}
}

func findAccountByLogin(vault *account.VaultWithDb) {
	findsAccount, err := vault.SearchAccountByUrl(askField("логин"), func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})

	if err != nil {
		output.PrintError(err.Error())
	}

	for _, findAccount := range findsAccount {
		findAccount.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	vault.DeleteAccountByUrl(askField("url"))
}

func createAccount(vault *account.VaultWithDb) {
	myAccount, err := account.NewAccount()

	if err != nil {
		output.PrintError("Неверный формат данных")
		return
	}
	vault.AddAccount(*myAccount)
}

func askField(field string) string {
	var searchField string
	fmt.Println("Введите ", field)
	fmt.Scanln(&searchField)

	return searchField
}

func PropmtData(vault *account.VaultWithDb) {
Menu:
	for {
		var variant string
		fmt.Println("1. Добавить аккаунт")
		fmt.Println("2. Найти аккаунт по URL")
		fmt.Println("3. Найти аккаунт по логину")
		fmt.Println("4. Удалить аккаунт")
		fmt.Println("5. Выход")

		fmt.Scanln(&variant)

		menuFunc := menu[variant]

		if menuFunc == nil {
			break Menu
		}

		menuFunc(vault)
	}
}
