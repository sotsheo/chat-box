package services

import (
	"chat-box/app/configs"
	"chat-box/app/helpers"
	"chat-box/app/models"
	"encoding/json"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
}

var file = helpers.FileHelper{}

func (us *UserService) Register(user models.UserModels) {
	data, _ := file.GetDataFile(configs.PathFileUser)
	var users []models.UserModels
	if data != "" {
		err := json.Unmarshal([]byte(data), &users)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	pw, _ := bcrypt.GenerateFromPassword([]byte(user.Pw), bcrypt.DefaultCost)
	user.Pw = string(pw)
	users = append(users, user)
	content, _ := json.Marshal(users)
	file.SaveDataFile(configs.PathFileUser, content)
}
