package main

import "time"

type OrderItem struct {
	ProductID string
	Quantity  int
	UnitPrice float64
}

type Order struct {
	ID          string
	Items       []OrderItem
	TotalAmount float64
	Status      string
	CreatedAt   time.Time
}
