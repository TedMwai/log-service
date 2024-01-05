package handler

import (
	"bytes"
	"encoding/json"
	"log-management/domain"
	"log-management/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestHandler_UpdateMicroservice(t *testing.T) {
	handler := &Handler{
		db: &mocks.MockDAO{},
	}

	// Create a sample request body
	microserviceReq := UpdateMicroserviceStruct{
		ID:          "sampleID",
		Name:        "Sample Name",
		Description: "Sample Description",
	}
	reqBody, err := json.Marshal(microserviceReq)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the sample request body
	req, err := http.NewRequest("PUT", "/microservice", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create a new router and register the handler
	r := chi.NewRouter()
	r.HandleFunc("/microservice", handler.UpdateMicroservice)

	// Serve the request using the router
	r.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Check the response body
	var microserviceRes domain.Microservice
	err = json.NewDecoder(rr.Body).Decode(&microserviceRes)
	if err != nil {
		t.Fatal(err)
	}

	// Compare the response body with the expected values
	if microserviceRes.ID != microserviceReq.ID {
		t.Errorf("handler returned unexpected ID: got %v want %v", microserviceRes.ID, microserviceReq.ID)
	}
	if microserviceRes.Name != microserviceReq.Name {
		t.Errorf("handler returned unexpected Name: got %v want %v", microserviceRes.Name, microserviceReq.Name)
	}
	if microserviceRes.Description != microserviceReq.Description {
		t.Errorf("handler returned unexpected Description: got %v want %v", microserviceRes.Description, microserviceReq.Description)
	}
}
