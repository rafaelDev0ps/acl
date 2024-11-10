package main

import (
	"encoding/xml"
	"fmt"
)

type OrderTanslator interface {
	ToLegacyXML(order *Order) (string, error)
	FromLegacyStatus(legacyStatus string) (string, error)
}

type Translator struct{}

func (trsl *Translator) ToLegacyXML(order *Order) (string, error) {
	orderXML, err := xml.MarshalIndent(order, " ", "  ")
	if err != nil {
		return "", fmt.Errorf("Error converting data to XML")
	}

	return string(orderXML), nil
}

func (trsl *Translator) FromLegacyStatus(legacyStatus string) string {
	_, exists := OrderStatuses[legacyStatus]
	if exists {
		return OrderStatuses[legacyStatus]
	}
	return OrderStatuses["UNKNOWN"]
}
