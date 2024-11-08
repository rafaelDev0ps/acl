package main

type OrderTanslator interface {
	ToLegacyXML(order *Order) string
	FromLegacyStatus(legacyStatus string) string
}

type Translator struct{}

func (trsl *Translator) ToLegacyXML(order *Order) string

func (trsl *Translator) FromLegacyStatus(legacyStatus string) string
