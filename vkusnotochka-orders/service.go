package main

import (
	"context"
	"errors"
	pb "github.com/intchris1/common/api"
	"log"
)

type service struct {
	store OrdersStore
}

func (s *service) ValidateOrder(ctx context.Context, rq *pb.CreateOrderRequest) error {
	if len(rq.Items) == 0 {
		return errors.New("items must have at least one item")
	}

	for _, i := range rq.Items {
		if i.Id == "" {
			return errors.New("item must have an ID")
		}
		if i.Quantity <= 0 {
			return errors.New("item must have a positive quantity")
		}
	}
	rq.Items = mergeItemsQuantity(rq.Items)
	log.Println(rq.Items)
	return nil
}

func NewService(store OrdersStore) *service {
	return &service{store}
}

func (s *service) CreateOrder(ctx context.Context) error {
	return nil
}

func mergeItemsQuantity(items []*pb.ItemsWithQuantity) []*pb.ItemsWithQuantity {
	merged := make(map[string]*pb.ItemsWithQuantity)
	for _, item := range items {
		current := merged[item.Id]
		if current == nil {
			merged[item.Id] = item
		} else {
			current.Quantity += item.Quantity
		}
	}
	var values []*pb.ItemsWithQuantity
	for _, v := range merged {
		values = append(values, v)
	}
	return values
}
