package files

import (
	"fmt"
	"os"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(filename string) *JsonDb {
	return &JsonDb{filename}
}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.filename)

	if err != nil {
		fmt.Println(err)
	}

	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		file.Close()
		fmt.Println(err)
		return
	}

	fmt.Println("Запись успешна изменена")
}
