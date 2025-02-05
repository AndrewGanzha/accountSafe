package account

import (
	"encoding/json"
	"passwordKeep/encrypter"
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
	enc   encrypter.Encrypter
}

func NewVault(db Db, enc encrypter.Encrypter) *VaultWithDb {
	file, err := db.Read()

	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}

	data := enc.Decrypt(file)

	var vault Vault
	err = json.Unmarshal(data, &vault)
	color.Cyan("Найдено %d аккаунтов", vault.Accounts)

	if err != nil {
		output.PrintError("Не удалось разобрать файл")
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
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

func (vault *VaultWithDb) SearchAccountByUrl(query string, checker func(Account, string) bool) ([]Account, error) {
	var includesAccount []Account

	for _, account := range vault.Accounts {
		isMatch := checker(account, query)

		if isMatch {
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
	encData := vault.enc.Encrypt(data)

	if err != nil {
		color.Red("Не удалось преобразовать файл")
	}

	vault.db.Write(encData)
}
