package handler

import (
	"context"
	"fmt"
	authpb "job-posting/gen/go/protos/auth"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type AuthHandler struct {}

var auth AuthHandler

func (a *AuthHandler) connGrpc() *grpc.ClientConn {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient("localhost:2000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	
	return conn
}

type UserLogin struct {
	Email    string	`json:"email" binding:"required"`
	Password string	`json:"password" binding:"required"`
}


func LoginUser(c *gin.Context) {
	var (
		paramUserCreate UserLogin
	)

	// validate input
	if err := c.ShouldBindJSON(&paramUserCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error parameter": err.Error()})
		return
	}

	fmt.Print(paramUserCreate)
	fmt.Println(c.PostForm("email"))
	fmt.Println(c.PostForm("password"))

	conn := auth.connGrpc()
	defer conn.Close()
	u := authpb.NewAuthServiceClient(conn)
	res, err := u.Login(context.Background(), &authpb.LoginRequest{
		Email:   paramUserCreate.Email,
		Password: paramUserCreate.Password,
	})

	if err != nil {
		if status.Code(err) != codes.InvalidArgument {
			log.Printf("Received unexpected error: %v", err)
			errorMessage := status.Convert(err).Message()
			c.AbortWithStatusJSON(500, gin.H{"status": false, "error": errorMessage})
			return
		} else if status.Code(err) == codes.Canceled {
			log.Printf("Received unexpected error: %v", err)
			c.AbortWithStatusJSON(500, gin.H{"status": false, "error": err.Error()})
			return	
		} else {
		log.Printf("Received error: %v", err)
		log.Print(err)
		c.AbortWithStatusJSON(500, gin.H{"status": false, "error": err.Error()})
			return
		}
	}

	c.JSON(200, res)
}

type UserRegister struct {
	Name	  string	`json:"name" binding:"required"`
	Email    string	`json:"email" binding:"required"`
	Password string	`json:"password" binding:"required"`
}

func RegisterUser(c *gin.Context) {
	var (
		paramUserCreate UserRegister
	)

	// validate input
	if err := c.ShouldBindJSON(&paramUserCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error parameter": err.Error()})
		return
	}

	conn := auth.connGrpc()
	defer conn.Close()
	u := authpb.NewAuthServiceClient(conn)
	res, err := u.Register(context.Background(), &authpb.RegisterRequest{
		Name:	  paramUserCreate.Name,
		Email:   paramUserCreate.Email,
		Password: paramUserCreate.Password,
	})

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"status": false, "error": err})
		return
	}

	c.JSON(200, res)
}
