package routes

import (
	// "github.com/gin-contrib/cors"
	// "github.com/gin-contrib/cors"
	"chat-box/app/configs"
	"chat-box/app/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "golang.org/x/time/rate"
)

// UserRoutes function
func Web(incomingRoutes *gin.Engine) {
	incomingRoutes.LoadHTMLGlob(configs.PathResource + "*.html")
	incomingRoutes.Static("/static", "./app/public/assets")
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Nguồn gốc của Next.js
	config.AllowMethods = []string{"*"}
	config.AllowHeaders = []string{"*"}
	incomingRoutes.Use(cors.New(config))
	userController := controllers.UserController{}
	incomingRoutes.GET("/", userController.Login())
	incomingRoutes.GET("/signup", userController.Signup())
	incomingRoutes.GET("/chat", userController.Chat())
}
