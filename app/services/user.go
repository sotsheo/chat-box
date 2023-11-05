package services

import (
	"chat-box/app/configs"
	"chat-box/app/helpers"
)

type UserService struct {
}

var file = helpers.FileHelper{}

func (us *UserService) Register() {
	file.GetDataFile(configs.PathFileUser)
}
