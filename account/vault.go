package account

import (
	"encoding/json"
	"github.com/fatih/color"
	"passwordKeep/files"
	"strings"
	"time"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewVault() *Vault {
	file, err := files.ReadFile("data.json")

	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)

	if err != nil {
		color.Red("Не удалось разобрать файл data.json")
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}

	return &vault
}

func (vault *Vault) AddAccount(account Account) {
	vault.Accounts = append(vault.Accounts, account)
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()

	if err != nil {
		color.Red("Не удалось преобразовать файл data.json")
	}
	files.WriteFile(data, "data.json")
}

func (vault *Vault) SearchAccount(search string) ([]Account, error) {
	var includesAccount []Account

	for _, account := range vault.Accounts {

		if strings.Contains(account.Url, search) {
			includesAccount = append(includesAccount, account)
		}
	}

	return includesAccount, nil
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)

	if err != nil {
		return nil, err
	}

	return file, nil
}
