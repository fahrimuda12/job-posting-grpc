package handler

import (
	"context"
	companypb "job-posting/gen/go/protos/company"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type CompanyHandler struct {}

var company CompanyHandler


type GetPayload struct {
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
		paramsGetCompany GetPayload
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

func GetCompanyByID(c *gin.Context) {
	id := c.Param("id")

	conn := company.connGrpc()
	defer conn.Close()
	u := companypb.NewCompanyServiceClient(conn)
	res, err := u.DetailCompany(context.Background(), &companypb.DetailCompanyRequest{Id: id})

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

type CreatePayload struct {
	Name    string	`json:"name" binding:"required"`
}


func CompanyCreate(c *gin.Context) {
	var (
		paramCompanyCreate CreatePayload
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

type UpdatePayload struct {
	Name    	string	`json:"name" binding:"required"`
}


func CompanyUpdate(c *gin.Context) {
	var (
		paramCompanyUpdate UpdatePayload
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

func CompanyDelete(c *gin.Context) {
	id := c.Param("id")
	conn := company.connGrpc()
	defer conn.Close()
	u := companypb.NewCompanyServiceClient(conn)
	res, err := u.DeleteCompany(context.Background(), &companypb.DeleteCompanyRequest{Id: id})

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
