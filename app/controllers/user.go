package controllers

import (
	"chat-box/app/configs"
	"chat-box/app/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
}

var userService = services.UserService{}

func (us *UserController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		// userService.Register()
		password := "abc1234!"
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("Lỗi khi băm mật khẩu:", err)
			return
		}
		http.ServeFile(c.Writer, c.Request, configs.PathResource+"signin.html")
	}
}
func (us *UserController) LoginPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, configs.PathResource+"signin.html")
	}
}
func (us *UserController) Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, configs.PathResource+"signup.html")
	}
}
func (us *UserController) Chat() gin.HandlerFunc {
	return func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, configs.PathResource+"index.html")
	}
}
