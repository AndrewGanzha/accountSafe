package main

import (
	"fmt"
	"passwordKeep/account"
	"passwordKeep/files"
	"passwordKeep/output"
)

func main() {
	showMenu()
}

func showMenu() {
	fmt.Println("Добро пожаловать в программу гененрации и получения паролей!")
	vault := account.NewVault(files.NewJsonDb("data.json"))

Menu:
	for {
		var input string
		fmt.Println("1. Добавить аккаунт")
		fmt.Println("2. Найти аккаунт")
		fmt.Println("3. Удалить аккаунт")
		fmt.Println("4. Выход")

		fmt.Scanln(&input)

		switch input {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		case "4":
			fmt.Println("До свидания!")
			break Menu
		}
	}
}

func findAccount(vault *account.VaultWithDb) {

	findsAccount, err := vault.SearchAccount(askUrl())

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

func askUrl() string {
	var url string
	fmt.Println("Введите URL для поиска")
	fmt.Scanln(&url)

	return url
}

func createAccount(vault *account.VaultWithDb) {
	myAccount, err := account.NewAccount()

	if err != nil {
		output.PrintError("Неверный формат данных")
		return
	}
	vault.AddAccount(*myAccount)
}
