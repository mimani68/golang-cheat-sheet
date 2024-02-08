package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsersEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUsersHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedResponse := "Hello, users!"
	if rr.Body.String() != expectedResponse {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expectedResponse)
	}
}