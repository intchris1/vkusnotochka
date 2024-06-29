package main

import (
	"context"
	common "github.com/intchris1/common"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {

	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer l.Close()
	store := NewStore()
	svc := NewService(store)
	NewGrpcHandler(grpcServer, svc)

	svc.CreateOrder(context.Background())

	log.Println("Grpc server started at ", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err)
	}
}
