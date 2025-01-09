package main

import (
	"fmt"
	"job-posting/services/auth/config"
	"job-posting/services/auth/migrations"
	"job-posting/services/auth/services"

	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	pb "job-posting/gen/go/protos/auth"
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
	l := log.New(os.Stdout, "auth: ", log.LstdFlags)

	//inisialiasai Gin
	db := config.DBInit()

	// Run database migrations
	migrations.Migrate(db)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "2000"))
	if err != nil {
		log.Fatalf("Couldn't create connection tcp %v", err)
	}

	srv := grpc.NewServer()
	authSRV := services.NewAuthService(l, db)
	pb.RegisterAuthServiceServer(srv, authSRV)
	log.Printf("Server start at port %s", "2000")

	srv.Serve(lis)
}