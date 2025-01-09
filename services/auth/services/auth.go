package services

import (
	"context"
	"errors"
	"job-posting/services/auth/models"
	"log"
	"os"
	"time"

	pb "job-posting/gen/go/protos/auth"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type AuthService struct {
	log *log.Logger
	DB  *gorm.DB

}

func NewAuthService(log *log.Logger, DB *gorm.DB) *AuthService {
	return &AuthService{log, DB}
}

func (a *AuthService) GetHelloAuth(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Auth",
	})
}

type UserCreate struct {
	Name     string	`json:"name" binding:"required"`
	Email    string	`json:"email" binding:"required"`
	Password string	`json:"password" binding:"required"`
}
func (a *AuthService) Register(ctx context.Context, req *pb.RegisterRequest)  (*pb.RegisterResponse, error) {

	var (
		userModel models.User
	)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New(err.Error())		
	}

	userModel.Name = req.GetName()
	userModel.Email = req.GetEmail()
	userModel.Password = string(passwordHash)
	
	// set role id to 2
	userModel.RoleID = 2

	err = a.DB.Create(&userModel).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}
	
	data := pb.RegisterResponse{
		Status:  true,
		Message: "Register success!",
	}

	return &data, nil
	
}

type UserLogin struct {
	Email    string	`json:"email" binding:"required"`
	Password string	`json:"password" binding:"required"`
}

// func (a *AuthService) Login(ctx context.Context, req *pb.LoginRequest)  (*pb.LoginResponse, error) {
func (a *AuthService) Login(ctx context.Context, req *pb.LoginRequest)  (*pb.LoginResponse, error) {
	
	var (
		userModel models.User
	)

	err := a.DB.Where("email = ?", req.GetEmail()).First(&userModel).Error
	if err != nil {
		return nil,  status.Errorf(codes.Canceled, "user not found")

	}

	err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(req.GetPassword()))
	if err != nil {
		// return nil, errors.New("password not match")
		return nil,  status.Errorf(codes.Canceled, "password not match")

	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userModel.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		// return nil, errors.New("failed to generate token")
		return nil,  status.Errorf(codes.Unknown, "failed to generate token")
	}


	data := pb.LoginResponse{
		Status:  true,
		Message: "Login success!",
		Token:   token,
	}

	return &data, nil
}


