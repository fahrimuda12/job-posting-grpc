package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {}


func LoginUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Login",
	})
}

type UserRegister struct {
	Name	  string	`json:"name" binding:"required"`
	Email    string	`json:"email" binding:"required"`
	Password string	`json:"password" binding:"required"`
}

func RegisterUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Register",
	})
}
