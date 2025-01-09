package services

import (
	"context"
	"job-posting/services/company/models"
	"log"

	pb "job-posting/gen/go/protos/company"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type CompanyService struct {
	log *log.Logger
	DB  *gorm.DB

}

func NewCompanyService(log *log.Logger, DB *gorm.DB) *CompanyService {
	return &CompanyService{log, DB}
}

func (a *CompanyService) GetCompany(ctx context.Context, req *pb.GetCompanyRequest)  (*pb.GetCompanyResponse, error) {
	
	var (
		companyModel []models.Company
		totalCompanies int64
	)

	page := req.GetPage()
	limit := req.GetLimit()
	offset := (page - 1) * limit

	// get all company with pagination
	err := a.DB.Limit(int(limit)).Offset(int(offset)).Find(&companyModel).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error when query data company")
	}

	// count all data in tabel company
	err = a.DB.Model(&models.Company{}).Count(&totalCompanies).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error when count data company")
	}
	
	var companyList []*pb.Company
	for _, company := range companyModel {
		companyList = append(companyList, &pb.Company{
			Id:   company.ID.String(),
			Name: company.Name,
		})
	}

	data := pb.GetCompanyResponse{
		Status:   true,
		Message: "Company List",
		Companies: companyList,
		PageNow: page,
		Limit: limit,
		TotalData: totalCompanies,
	}

	return &data, nil
}

func (a *CompanyService) DetailCompany(ctx context.Context, req *pb.DetailCompanyRequest)  (*pb.DetailCompanyResponse, error) {
	
	var (
		companyModel models.Company
	)

	err := a.DB.First(&companyModel, "id = ?", req.GetId()).Error
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Company not found")
	}

	data := pb.DetailCompanyResponse{
		Status: true,
		Company: &pb.Company{
			Id:   companyModel.ID.String(),
			Name: companyModel.Name,
		},
	}

	return &data, nil
}

type CompanyCreate struct {
	Name     string	`json:"name" binding:"required"`
}

func (a *CompanyService) CreateCompany(ctx context.Context, req *pb.CreateCompanyRequest)  (*pb.CreateCompanyResponse, error) {

	var (
		companyModel models.Company
	)

	companyModel.Name = req.GetName()
	

	err := a.DB.Create(&companyModel).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error when create company")
	}
	
	data := pb.CreateCompanyResponse{
		Status:  true,
		Message: "Company Created!",
	}

	return &data, nil
}

func (a *CompanyService) UpdateCompany(ctx context.Context, req *pb.UpdateCompanyRequest)  (*pb.UpdateCompanyResponse, error) {
	
	var (
		companyModel models.Company
	)

	err := a.DB.First(&companyModel,"id = ?", req.GetId()).Error
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Company not found")
	}

	companyModel.Name = req.GetName()
	

	err = a.DB.Save(&companyModel).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error when update company")
	}
	
	data := pb.UpdateCompanyResponse{
		Status:  true,
		Message: "Company Updated!",
	}

	return &data, nil
}

func (a *CompanyService) DeleteCompany(ctx context.Context, req *pb.DeleteCompanyRequest)  (*pb.DeleteCompanyResponse, error) {
	
	var (
		companyModel models.Company
	)

	err := a.DB.First(&companyModel, "id = ?",req.GetId()).Error
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Company not found")
	}

	err = a.DB.Delete(&companyModel).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error when delete company")
	}
	
	data := pb.DeleteCompanyResponse{
		Status:  true,
		Message: "Company Deleted!",
	}

	return &data, nil
}




