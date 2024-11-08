package main

type LegacyOrderService interface {
	SubmitOrder(orderXML string) string
	GetOrderStatus(legacyOrderRef string) string
}

type LegacyService struct{}

func (svc *LegacyService) SubmitOrder(orderXML string) string

func (svc *LegacyService) GetOrderStatus(legacyOrderRef string) string
