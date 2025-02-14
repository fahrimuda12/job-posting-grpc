package handler

import (
	"context"
	"job-posting/gateway/helper"
	companypb "job-posting/gen/go/protos/company"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type CompanyHandler struct {}

var company CompanyHandler


type GetPayloadCompany struct {
	Page  int32 `form:"page"`
	Limit int32 `form:"limit"`
}

func (a *CompanyHandler) connGrpc() *grpc.ClientConn {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient("localhost:2010", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	
	return conn
}

func GetCompany(c *gin.Context) {
	var (
		paramsGetCompany GetPayloadCompany
	)

	// validate input
	if err := c.ShouldBindQuery(&paramsGetCompany); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error parameter": err.Error()})
		return
	}

	page := paramsGetCompany.Page
	if page == 0 {
		page = 1
	}
	limit := paramsGetCompany.Limit
	if limit == 0 {
		limit = 10
	}

	conn := company.connGrpc()
	defer conn.Close()
	u := companypb.NewCompanyServiceClient(conn)
	res, err := u.GetCompany(context.Background(), &companypb.GetCompanyRequest{
		Page:   page,
		Limit:  limit,
	})

	if err != nil {
		statusCode := status.Code(err)
		code := helper.SwitchStatusToCode(int(statusCode))
		message := status.Convert(err).Message()
		c.JSON(code, helper.ErrorResponse(message, statusCode.String()))
		return
	}

	c.JSON(200, res)
}

func GetCompanyByID(c *gin.Context) {
	id := c.Param("id")

	conn := company.connGrpc()
	defer conn.Close()
	u := companypb.NewCompanyServiceClient(conn)
	res, err := u.DetailCompany(context.Background(), &companypb.DetailCompanyRequest{Id: id})

	if err != nil {
		statusCode := status.Code(err)
		code := helper.SwitchStatusToCode(int(statusCode))
		message := status.Convert(err).Message()
		c.JSON(code, helper.ErrorResponse(message, statusCode.String()))
		return
	}

	c.JSON(200, res)
}

type CreatePayloadCompany struct {
	Name    string	`json:"name" binding:"required"`
}


func CompanyCreate(c *gin.Context) {
	var (
		paramCompanyCreate CreatePayloadCompany
	)

	// validate input
	if err := c.ShouldBindJSON(&paramCompanyCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error parameter": err.Error()})
		return
	}


	conn := company.connGrpc()
	defer conn.Close()
	u := companypb.NewCompanyServiceClient(conn)
	res, err := u.CreateCompany(context.Background(), &companypb.CreateCompanyRequest{
		Name:   paramCompanyCreate.Name,
	})

	if err != nil {
		statusCode := status.Code(err)
		code := helper.SwitchStatusToCode(int(statusCode))
		message := status.Convert(err).Message()
		c.JSON(code, helper.ErrorResponse(message, statusCode.String()))
		return
	}

	c.JSON(200, res)
}

type UpdatePayloadCompany struct {
	Name    	string	`json:"name" binding:"required"`
}


func CompanyUpdate(c *gin.Context) {
	var (
		paramCompanyUpdate UpdatePayloadCompany
	)

	// validate input
	if err := c.ShouldBindJSON(&paramCompanyUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error parameter": err.Error()})
		return
	}

	Id := c.Param("id")

	conn := company.connGrpc()
	defer conn.Close()
	u := companypb.NewCompanyServiceClient(conn)
	res, err := u.UpdateCompany(context.Background(), &companypb.UpdateCompanyRequest{
		Id:     Id,
		Name:   paramCompanyUpdate.Name,
	})

	if err != nil {
		statusCode := status.Code(err)
		code := helper.SwitchStatusToCode(int(statusCode))
		message := status.Convert(err).Message()
		c.JSON(code, helper.ErrorResponse(message, statusCode.String()))
		return
	}

	c.JSON(200, res)
}

func CompanyDelete(c *gin.Context) {
	id := c.Param("id")
	conn := company.connGrpc()
	defer conn.Close()
	u := companypb.NewCompanyServiceClient(conn)
	res, err := u.DeleteCompany(context.Background(), &companypb.DeleteCompanyRequest{Id: id})

	if err != nil {
		statusCode := status.Code(err)
		code := helper.SwitchStatusToCode(int(statusCode))
		message := status.Convert(err).Message()
		c.JSON(code, helper.ErrorResponse(message, statusCode.String()))
		return
	}

	c.JSON(200, res)
}
