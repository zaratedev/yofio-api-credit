package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"yofio-api/credit/handlers"
)

func TestIntegrationAssigmentInvesment(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handlers.AssigmentInvestment))
	defer server.Close()

	requestBody := map[string]int32{"Investment": 3000}
	bodyBytes, _ := json.Marshal(requestBody)
	bodyReader := bytes.NewReader(bodyBytes)

	resp, err := http.Post(server.URL+"/credit-assigment", "application/json", bodyReader)
	if err != nil {
		t.Fatalf("Error al realizar la solicitud: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Código de estado incorrecto. Esperado %d pero recibido %d", http.StatusOK, resp.StatusCode)
	}
}
