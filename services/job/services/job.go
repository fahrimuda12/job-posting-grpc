package services

import (
	"context"
	"errors"
	"job-posting/services/job/models"
	"log"

	pb "job-posting/gen/go/protos/job"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gorm.io/gorm"
)

type JobService struct {
	log *log.Logger
	DB  *gorm.DB

}

func NewJobService(log *log.Logger, DB *gorm.DB) *JobService {
	return &JobService{log, DB}
}

func (a *JobService) GetJob(ctx context.Context, req *pb.GetJobRequest)  (*pb.GetJobResponse, error) {
	
	var (
		jobModel []models.Job
		totalJobs int64
	)

	page := req.GetPage()
	limit := req.GetLimit()
	offset := (page - 1) * limit

	// get all job with pagination
	err := a.DB.Limit(int(limit)).Offset(int(offset)).Preload("Company").Find(&jobModel).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error when query data job")
	}

	// count all data in tabel job
	err = a.DB.Model(&models.Job{}).Count(&totalJobs).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}
	
	var jobList []*pb.Job
	for _, job := range jobModel {
		jobList = append(jobList, &pb.Job{
			Id:   job.ID.String(),
			Title: job.Title,
			Description: job.Description,
			Company: &pb.Company{
				Id:   job.Company.ID.String(),
				Name: job.Company.Name,
			},
		})
	}
	

	// jika data job kosong
	if len(jobList) == 0 {
		return nil, status.Errorf(codes.NotFound, "Job not found")
	}

	data := pb.GetJobResponse{
		Status:   true,
		Message: "Job List",
		Jobs: jobList,
		PageNow: page,
		Limit: limit,
		TotalData: totalJobs,
	}

	return &data, nil
}

func (a *JobService) DetailJob(ctx context.Context, req *pb.DetailJobRequest)  (*pb.DetailJobResponse, error) {
	
	var (
		jobModel models.Job
	)

	err := a.DB.Preload("Company").First(&jobModel, "id = ?", req.GetId()).Error
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Job not found")
	}

	data := pb.DetailJobResponse{
		Status: true,
		Job: &pb.Job{
			Id:   jobModel.ID.String(),
			Title: jobModel.Title,
			Description: jobModel.Description,
			Company: &pb.Company{
				Id:   jobModel.Company.ID.String(),
				Name: jobModel.Company.Name,
			},
		},
	}

	return &data, nil
}

type JobCreate struct {
	Name     string	`json:"name" binding:"required"`
}

func (a *JobService) CreateJob(ctx context.Context, req *pb.CreateJobRequest)  (*pb.CreateJobResponse, error) {

	var (
		jobModel models.Job
		companyModel models.Company
	)

	jobModel.Title = req.GetTitle()
	jobModel.Description = req.GetDescription()
	companyID, err := uuid.Parse(req.GetCompanyId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid company ID")
	}

	// check apakah company id ada di tabel company
	err = a.DB.First(&companyModel, "id = ?", companyID).Error
	if err != nil {
		return nil, errors.New("company ID not found")
	}

	jobModel.CompanyID = companyID
	

	err = a.DB.Create(&jobModel).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error when create job")
	}
	
	data := pb.CreateJobResponse{
		Status:  true,
		Message: "Job Created!",
	}

	return &data, nil
}

func (a *JobService) UpdateJob(ctx context.Context, req *pb.UpdateJobRequest)  (*pb.UpdateJobResponse, error) {
	
	var (
		jobModel models.Job
		companyModel models.Company
	)

	err := a.DB.First(&jobModel,"id = ?", req.GetId()).Error
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Job not found")
	}

	jobModel.Title = req.GetTitle()
	jobModel.Description = req.GetDescription()
	companyID, err := uuid.Parse(req.GetCompanyId())
	if err != nil {
		return nil, errors.New("invalid company ID")
	}
	jobModel.CompanyID = companyID
	// check apakah company id ada di tabel company
	err = a.DB.First(&companyModel, "id = ?", companyID).Error
	if err != nil {
		return nil, errors.New("company ID not found")
	}

	err = a.DB.Save(&jobModel).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}
	
	data := pb.UpdateJobResponse{
		Status:  true,
		Message: "Job Updated!",
	}

	return &data, nil
}

func (a *JobService) DeleteJob(ctx context.Context, req *pb.DeleteJobRequest)  (*pb.DeleteJobResponse, error) {
	
	var (
		jobModel models.Job
	)

	err := a.DB.First(&jobModel, "id = ?",req.GetId()).Error
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Job not found")
	}

	err = a.DB.Delete(&jobModel).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error when delete job")
	}
	
	data := pb.DeleteJobResponse{
		Status:  true,
		Message: "Job Deleted!",
	}

	return &data, nil
}




