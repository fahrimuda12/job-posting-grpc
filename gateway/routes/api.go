package routes

import (
	"job-posting/gateway/handler"
	"job-posting/gateway/middleware"

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

	// route group /company
	company := routes.Group("/company")
	company.GET("", middleware.CheckAuth, handler.GetCompany)
	company.GET("/:id",  middleware.CheckAuth, handler.GetCompanyByID)
	company.POST("/create",  middleware.CheckAuth, handler.CompanyCreate)
	company.PUT("/update/:id",  middleware.CheckAuth, handler.CompanyUpdate)
	company.DELETE("/delete/:id",  middleware.CheckAuth,  middleware.CheckAuth, handler.CompanyDelete)
	
	
}