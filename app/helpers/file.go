package helpers

import (
	"fmt"
	"io/ioutil"
)

type FileHelper struct {
}

func (f *FileHelper) GetDataFile(fileName string) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Không thể đọc tệp:", err)
		return
	}

	// In nội dung tệp
	fmt.Println(string(content))
}
