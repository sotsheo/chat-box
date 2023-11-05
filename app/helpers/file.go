package helpers

import (
	"chat-box/app/configs"
	"fmt"
	"io/ioutil"
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

func (f *FileHelper) CreateUser(fileName string) {
	data, _ := f.GetDataFile(configs.PathFileUser)
	if data != "" {

	}
}
