package account

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"net/url"
	"time"
)

type Account struct {
	Login     string    `json:"login" xml:"login"`
	Password  string    `json:"password" xml:"password"`
	Url       string    `json:"url" xml:"url"`
	CreatedAt time.Time `json:"createdAt" xml:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" xml:"updatedAt"`
}

func (acc *Account) generatePassword() {
	var length int

	fmt.Println("Введите желаемую длину пароля")
	fmt.Scanln(&length)

	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	runes := []rune(alphabet)
	password := make([]rune, length)

	for i := range password {
		password[i] = runes[rand.Intn(len(runes))]
	}

	acc.Password = string(password)
}

func (acc Account) Output() {
	color.Cyan(acc.Login)
	color.Cyan(acc.Password)
	color.Cyan(acc.Url)
}

func NewAccount() (*Account, error) {
	newAccount := &Account{}
	fmt.Println("Введите логин")
	fmt.Scanln(&newAccount.Login)

	if newAccount.Login == "" {
		return nil, errors.New("Неверный формат логина")
	}

	fmt.Println("Введите url")
	fmt.Scanln(&newAccount.Url)

	_, err := url.ParseRequestURI(newAccount.Url)
	if err != nil {
		return nil, errors.New("Неверный формат URL")
	}

	fmt.Println("Введите пароль или нажимте enter для генерации пароля")

	fmt.Scanln(&newAccount.Password)

	if newAccount.Password == "" {
		newAccount.generatePassword()
	}

	newAccount.CreatedAt = time.Now()
	newAccount.UpdatedAt = time.Now()

	return newAccount, nil
}
