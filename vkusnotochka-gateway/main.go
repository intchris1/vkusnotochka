package main

import (
	pb "github.com/intchris1/common/api"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"

	common "github.com/intchris1/common"
)

var (
	err                 = godotenv.Load()
	httpAddr            = common.EnvString("HTTP_ADDR", ":8080")
	orderServiceAddress = "localhost:2000"
)

func main() {
	conn, err := grpc.Dial(orderServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}

	log.Println("Dialing orders service at ", orderServiceAddress)

	c := pb.NewOrderServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)

	log.Println("Starting server on " + httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server")
	}
}
