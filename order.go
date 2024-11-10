package main

import (
	"time"
)

type OrderItem struct {
	ProductID string  `xml:"product_id" json:"product_id"`
	Quantity  int     `xml:"quantity" json:"quantity"`
	UnitPrice float64 `xml:"unit_price" json:"unit_price"`
}

type Order struct {
	ID          string      `xml:"id" json:"id"`
	Items       []OrderItem `xml:"items" json:"items"`
	TotalAmount float64     `xml:"total_amount" json:"total_amount"`
	Status      string      `xml:"status" json:"status"`
	CreatedAt   time.Time   `xml:"created_at" json:"created_at"`
}

var OrderStatuses = map[string]string{
	"PROCESSED": "completed",
	"PENDING":   "processing",
	"REJECTED":  "failed",
	"UNKNOWN":   "unknown",
}
