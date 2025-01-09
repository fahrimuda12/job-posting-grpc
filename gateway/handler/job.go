package handler

import (
	"context"
	"job-posting/gateway/helper"
	jobpb "job-posting/gen/go/protos/job"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type JobHandler struct {}

var job JobHandler


type GetPayloadJob struct {
	Page  int32 `form:"page"`
	Limit int32 `form:"limit"`
}

func (a *JobHandler) connGrpc() *grpc.ClientConn {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient("localhost:2020", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	
	return conn
}

func GetJob(c *gin.Context) {
	var (
		paramsGetJob GetPayloadJob
	)

	// validate input
	if err := c.ShouldBindQuery(&paramsGetJob); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error parameter": err.Error()})
		return
	}

	page := paramsGetJob.Page
	if page == 0 {
		page = 1
	}
	limit := paramsGetJob.Limit
	if limit == 0 {
		limit = 10
	}

	conn := job.connGrpc()
	defer conn.Close()
	u := jobpb.NewJobServiceClient(conn)
	res, err := u.GetJob(context.Background(), &jobpb.GetJobRequest{
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

func GetJobByID(c *gin.Context) {
	id := c.Param("id")

	conn := job.connGrpc()
	defer conn.Close()
	u := jobpb.NewJobServiceClient(conn)
	res, err := u.DetailJob(context.Background(), &jobpb.DetailJobRequest{Id: id})

	if err != nil {
		statusCode := status.Code(err)
		code := helper.SwitchStatusToCode(int(statusCode))
		message := status.Convert(err).Message()
		c.JSON(code, helper.ErrorResponse(message, statusCode.String()))
		return
	}

	c.JSON(200, res)
}

type CreatePayloadJob struct {
	Title   	string	`json:"title" binding:"required"`
	Description string	`json:"description" binding:"required"`
	CompanyId 	string	`json:"company_id" binding:"required"`
}


func JobCreate(c *gin.Context) {
	var (
		paramJobCreate CreatePayloadJob
	)

	// validate input
	if err := c.ShouldBindJSON(&paramJobCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error parameter": err.Error()})
		return
	}


	conn := job.connGrpc()
	defer conn.Close()
	u := jobpb.NewJobServiceClient(conn)
	res, err := u.CreateJob(context.Background(), &jobpb.CreateJobRequest{
		Title:   paramJobCreate.Title,
		Description: paramJobCreate.Description,
		CompanyId: paramJobCreate.CompanyId,
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

type UpdatePayload struct {
	Title  	string	`json:"title" binding:"required"`
	Description string	`json:"description" binding:"required"`
	CompanyId 	string	`json:"company_id" binding:"required"`
}


func JobUpdate(c *gin.Context) {
	var (
		paramJobUpdate UpdatePayload
	)

	// validate input
	if err := c.ShouldBindJSON(&paramJobUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error parameter": err.Error()})
		return
	}

	Id := c.Param("id")

	conn := job.connGrpc()
	defer conn.Close()
	u := jobpb.NewJobServiceClient(conn)
	res, err := u.UpdateJob(context.Background(), &jobpb.UpdateJobRequest{
		Id:     Id,
		Title:  paramJobUpdate.Title,
		Description: paramJobUpdate.Description,
		CompanyId: paramJobUpdate.CompanyId,
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

func JobDelete(c *gin.Context) {
	id := c.Param("id")
	conn := job.connGrpc()
	defer conn.Close()
	u := jobpb.NewJobServiceClient(conn)
	res, err := u.DeleteJob(context.Background(), &jobpb.DeleteJobRequest{Id: id})

	if err != nil {
		statusCode := status.Code(err)
		code := helper.SwitchStatusToCode(int(statusCode))
		message := status.Convert(err).Message()
		c.JSON(code, helper.ErrorResponse(message, statusCode.String()))
		return

	}

	c.JSON(200, res)
}
