package main

type ModernOrderService interface {
	CreateOrder(order *Order) bool
	GetOrderStatus(orderID string) string
}

type ModernService struct{}

func (svc *ModernService) CreateOrder() bool

func (svc *ModernService) GetOrderStatus() string
