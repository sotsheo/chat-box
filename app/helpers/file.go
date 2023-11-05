package helpers

import (
	"fmt"
	"io/ioutil"
	"os"
)

type FileHelper struct {
}

func (f *FileHelper) GetDataFile(fileName string) (string, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Không thể đọc tệp:", err)
		return "", nil
	}
	return string(content), nil
}

func (f *FileHelper) SaveDataFile(fileName string, data []byte) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return
	}
}
