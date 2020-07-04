package services

import (
	"context"
	"fmt"
)

type OrderService struct {

}

func (this *OrderService)NewOrder(ctx context.Context, order *Order) (*OrderResponse, error){

	fmt.Println(order)
	return &OrderResponse{Msg: "123",Status: "1"},nil
}