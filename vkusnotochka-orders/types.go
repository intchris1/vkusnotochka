package main

import (
	"context"
	pb "github.com/intchris1/common/api"
)

type OrdersService interface {
	CreateOrder(ctx context.Context) error
	ValidateOrder(ctx context.Context, request *pb.CreateOrderRequest) error
}

type OrdersStore interface {
	Create(ctx context.Context) error
}
