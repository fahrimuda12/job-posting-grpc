package main

import (
	"fmt"

	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"job-posting/core/config"
	pb "job-posting/gen/go/protos/job"
	"job-posting/services/job/migrations"
	"job-posting/services/job/services"
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
	l := log.New(os.Stdout, "job: ", log.LstdFlags)

	//inisialiasai Gin
	db := config.DBInit()

	// Run database migrations
	migrations.Migrate(db)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "2020"))
	if err != nil {
		log.Fatalf("Couldn't create connection tcp %v", err)
	}

	srv := grpc.NewServer()
	jobService := services.NewJobService(l, db)
	pb.RegisterJobServiceServer(srv, jobService)
	log.Printf("Server start at port %s", "2020")

	srv.Serve(lis)
}