package main

import "fmt"

type LegacyOrderAdapter interface {
	SubmitOrder(order *Order) (string, error)
	GetOrderStatus(legacyOrderRef string) (string, error)
}

type LegacyAdapter struct {
	Translator    Translator
	LegacyService LegacyService
}

func (adap *LegacyAdapter) SubmitOrder(order *Order) (string, error) {
	trans, err := adap.Translator.ToLegacyXML(order)
	if err != nil {
		return "", fmt.Errorf("Error translating to legacy XML")
	}

	submit, err := adap.LegacyService.SubmitOrder(trans)
	if err != nil {
		return "", err
	}

	return submit, nil
}

func (adap *LegacyAdapter) GetOrderStatus(legacyOrderRef string) (string, error) {
	status, err := adap.LegacyService.GetOrderStatus(legacyOrderRef)
	if err != nil {
		return "", fmt.Errorf("Error fetching legacy order reference")
	}

	legacyStatus := adap.Translator.FromLegacyStatus(status)

	return legacyStatus, nil
}
