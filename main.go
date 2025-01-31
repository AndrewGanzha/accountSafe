package main

import (
	"fmt"
	"passwordKeep/account"
)

func main() {
	showMenu()
	//createAccount()
}

func showMenu() {
	fmt.Println("Добро пожаловать в программу гененрации и получения паролей!")

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
			createAccount()
		case "2":
			findAccount()
		case "3":
			deleteAccount()
		case "4":
			fmt.Println("До свидания!")
			break Menu
		}
	}
}

func findAccount() {

}

func deleteAccount() {

}

func createAccount() {
	myAccount, err := account.NewAccount()

	if err != nil {
		fmt.Println("Неверный формат данных")
		return
	}

	vault := account.NewVault()
	vault.AddAccount(*myAccount)
}
