package account

import (
	"encoding/json"
	"passwordKeep/output"
	"strings"
	"time"

	"github.com/fatih/color"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type Db interface {
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

type VaultWithDb struct {
	Vault `json:"vault"`
	db    Db
}

func NewVault(db Db) *VaultWithDb {
	file, err := db.Read()

	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)

	if err != nil {
		output.PrintError("Не удалось разобрать файл data.json")
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}

	return &VaultWithDb{
		Vault: vault,
		db:    db,
	}
}

func (vault *VaultWithDb) AddAccount(account Account) {
	vault.Accounts = append(vault.Accounts, account)
	vault.saveVault()
}

func (vault *VaultWithDb) SearchAccount(search string) ([]Account, error) {
	var includesAccount []Account

	for _, account := range vault.Accounts {

		if strings.Contains(account.Url, search) {
			includesAccount = append(includesAccount, account)
		}
	}

	if len(includesAccount) == 0 {
		output.PrintError("Аккаунтов не найдено")
	}

	return includesAccount, nil
}

func (vault *VaultWithDb) DeleteAccountByUrl(url string) bool {
	var accounts []Account
	isDeleted := false
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if !isMatched {
			accounts = append(accounts, account)
			continue
		}
		isDeleted = true
	}

	vault.Accounts = accounts
	vault.saveVault()

	return isDeleted
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func (vault *VaultWithDb) saveVault() {
	vault.UpdatedAt = time.Now()
	data, err := vault.Vault.ToBytes()

	if err != nil {
		color.Red("Не удалось преобразовать файл data.json")
	}

	vault.db.Write(data)
}
