package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

type LegacyOrderService interface {
	SubmitOrder(orderXML string) (string, error)
	GetOrderStatus(legacyOrderRef string) (string, error)
}

type LegacyService struct{}

func (svc *LegacyService) SubmitOrder(orderXML string) (string, error) {
	client := &http.Client{} // i should create a separate function to setup http client, but not today

	req, err := http.NewRequest("POST", os.Getenv("LEGACY_ENDPOINT"), bytes.NewBuffer([]byte(orderXML)))
	if err != nil {
		return "", fmt.Errorf("Error creating request to legacy system")
	}

	req.Header.Add("Content-Type", "application/xml; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error sending request to legacy system")
	}

	// I know I should threat all the status codes of the legacy API, but not today again, this is a reference example, bear with me!
	if resp.Status != "200" {
		return "", fmt.Errorf("Order not succeeded")
	}

	return "ORDER SUBMITTED", nil
}

func (svc *LegacyService) GetOrderStatus(legacyOrderRef string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", os.Getenv("LEGACY_ENDPOINT")+legacyOrderRef, nil)
	if err != nil {
		return "", fmt.Errorf("Error creating request to legacy system")
	}

	req.Header.Add("Content-Type", "application/xml; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error sending request to legacy system")
	}

	if resp.Status != "200" {
		return "", fmt.Errorf("Order not succeeded")
	}
	return "The order ref", nil
}
