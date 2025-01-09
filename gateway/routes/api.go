package routes

import (
	"job-posting/gateway/handler"

	"github.com/gin-gonic/gin"
)

func Routes(routes *gin.Engine) {

	// register log
	
	// route get /auth kembalikan response random string

	routes.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	routes.POST("/auth/login", handler.LoginUser)
	routes.POST("/auth/register", handler.RegisterUser)

}