package services

import "chat-box/app/helpers"

type UserService struct {
}

var file = helpers.FileHelper{}

func (us *UserService) Register() {
	file.GetDataFile("./app/data/user.txt")
}
