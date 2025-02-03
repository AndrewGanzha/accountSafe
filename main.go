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
	"2": findAccount,
	"3": deleteAccount,
}

func main() {
	showMenu()
}

func showMenu() {
	fmt.Println("Добро пожаловать в программу гененрации и получения паролей!")

	vault := account.NewVault(files.NewJsonDb("data.json"))
	PropmtData(vault)
}

func findAccount(vault *account.VaultWithDb) {
	findsAccount, err := vault.SearchAccountByUrl(askUrl(), func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})

	if err != nil {
		output.PrintError(err.Error())
	}

	for _, findAccount := range findsAccount {
		findAccount.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	vault.DeleteAccountByUrl(askUrl())
}

func createAccount(vault *account.VaultWithDb) {
	myAccount, err := account.NewAccount()

	if err != nil {
		output.PrintError("Неверный формат данных")
		return
	}
	vault.AddAccount(*myAccount)
}

func askUrl() string {
	var url string
	fmt.Println("Введите URL для поиска")
	fmt.Scanln(&url)

	return url
}

func PropmtData(vault *account.VaultWithDb) {
Menu:
	for {
		var variant string
		fmt.Println("1. Добавить аккаунт")
		fmt.Println("2. Найти аккаунт")
		fmt.Println("3. Удалить аккаунт")
		fmt.Println("4. Выход")

		fmt.Scanln(&variant)

		menuFunc := menu[variant]

		if menuFunc == nil {
			break Menu
		}

		menuFunc(vault)
	}
}
