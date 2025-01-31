package files

import (
	"fmt"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile("data.json")

	if err != nil {
		return nil, err
	}

	return data, nil
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)

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

	fmt.Println("Запись успешна")
}
