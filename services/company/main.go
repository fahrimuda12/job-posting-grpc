package main

import (
	"fmt"

	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"job-posting/core/config"
	pb "job-posting/gen/go/protos/company"
	"job-posting/services/company/migrations"
	"job-posting/services/company/services"
)

func LoadEnvs() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}


func init() {
	LoadEnvs()
	config.DBInit()

}

func main() {
	l := log.New(os.Stdout, "company: ", log.LstdFlags)

	//inisialiasai Gin
	db := config.DBInit()

	// Run database migrations
	migrations.Migrate(db)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "2010"))
	if err != nil {
		log.Fatalf("Couldn't create connection tcp %v", err)
	}

	srv := grpc.NewServer()
	companyService := services.NewCompanyService(l, db)
	pb.RegisterCompanyServiceServer(srv, companyService)
	log.Printf("Server start at port %s", "2010")

	srv.Serve(lis)
}