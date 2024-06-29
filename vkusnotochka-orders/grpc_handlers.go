package main

import (
	"context"
	pb "github.com/intchris1/common/api"
	"google.golang.org/grpc"
	"log"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
	service OrdersService
}

func NewGrpcHandler(grpcServer *grpc.Server, service OrdersService) *grpcHandler {
	handler := &grpcHandler{service: service}
	pb.RegisterOrderServiceServer(grpcServer, handler)
	return handler
}

func (h *grpcHandler) CreateOrder(ctx context.Context, rq *pb.CreateOrderRequest) (*pb.Order, error) {

	log.Printf("New order received! Order %v", rq)
	o := &pb.Order{
		Id: "42",
	}
	return o, nil
}
