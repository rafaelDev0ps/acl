package main

type LegacyOrderAdapter interface {
	SubmitOrder(order *Order) string
	GetOrderStatus(legacyOrderRef string) string
}

type LegacyAdapter struct{}

func (adap *LegacyAdapter) SubmitOrder() string

func (adap *LegacyAdapter) GetOrderStatus(legacyOrderRef string) string
