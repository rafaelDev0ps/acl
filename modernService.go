package main

import "fmt"

type ModernOrderService interface {
	CreateOrder(order *Order) (bool, error)
	GetOrderStatus(orderID string) string
}

type ModernService struct {
	Adapter   LegacyAdapter
	OrderRefs []string
}

func (svc *ModernService) CreateOrder(order *Order) (bool, error) {
	submit, err := svc.Adapter.SubmitOrder(order)
	if err != nil {
		return false, err
	}

	svc.OrderRefs = append(svc.OrderRefs, submit)

	return true, nil
}

func (svc *ModernService) GetOrderStatus(orderID string) (string, error) {
	for _, order := range svc.OrderRefs {
		if order == orderID {
			status, err := svc.Adapter.GetOrderStatus(order)
			if err != nil {
				return "", err
			}
			return status, nil
		}
	}

	return "", fmt.Errorf("Order status not found")
}
