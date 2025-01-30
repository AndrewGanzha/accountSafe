package main

import (
	"fmt"
	"passwordKeep/account"
	"passwordKeep/files"
)

func main() {
	createAccount()
}

func createAccount() {
	myAccount, err := account.NewAccount()

	if err != nil {
		fmt.Println("Неверный формат данных")
		return
	}

	file, err := myAccount.ToBytes()

	if err != nil {
		fmt.Println("Не удалось сохранить")
		return
	}

	files.WriteFile(file, "data.json")
}
