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

	// route group /job
	job := routes.Group("/job")
	job.GET("", middleware.CheckAuth, handler.GetJob)
	job.GET("/:id",  middleware.CheckAuth, handler.GetJobByID)
	job.POST("/create",  middleware.CheckAuth, handler.JobCreate)
	job.PUT("/update/:id",  middleware.CheckAuth, handler.JobUpdate)
	job.DELETE("/delete/:id",  middleware.CheckAuth, handler.JobDelete)
	
}