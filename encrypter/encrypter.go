package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	key := os.Getenv("KEY")
	if key == "" {
		panic("KEY environment variable not set")
	}
	return &Encrypter{Key: key}
}

func (encrypt *Encrypter) Encrypt(encryptString []byte) []byte {
	block, err := aes.NewCipher([]byte(encrypt.Key))

	if err != nil {
		panic(err.Error())
	}

	aesGSM, err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, aesGSM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)

	if err != nil {
		panic(err.Error())
	}

	return aesGSM.Seal(nonce, nonce, encryptString, nil)
}

func (encrypt *Encrypter) Decrypt(decryptString []byte) []byte {
	block, err := aes.NewCipher([]byte(encrypt.Key))

	if err != nil {
		panic(err.Error())
	}

	aesGSM, err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}
	nonceSize := aesGSM.NonceSize()
	nonce, cipherText := decryptString[:nonceSize], decryptString[nonceSize:]
	plainText, err := aesGSM.Open(nil, nonce, cipherText, nil)

	if err != nil {
		panic(err.Error())
	}

	return plainText
}
