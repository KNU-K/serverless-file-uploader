package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlePing(t *testing.T) {
    // Test for a successful GET request
    req, err := http.NewRequest("GET", "/api/ping", nil)
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handlePing)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    expected := "pong"
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }

    // Test for a non-GET request
    req, err = http.NewRequest("POST", "/api/ping", nil)
    if err != nil {
        t.Fatal(err)
    }
    rr = httptest.NewRecorder()
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusMethodNotAllowed {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
    }
}