package main

type LegacyOrderService interface {
	SubmitOrder(orderXML string) (string, error)
	GetOrderStatus(legacyOrderRef string) (string, error)
}

type LegacyService struct{}

func (svc *LegacyService) SubmitOrder(orderXML string) (string, error) {
	return "ORDER SUBMITTED", nil
}

func (svc *LegacyService) GetOrderStatus(legacyOrderRef string) (string, error) {
	return "Random order ref", nil
}
