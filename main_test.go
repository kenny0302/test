package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPay(t *testing.T) {

	var jsonStr = []byte(`{"Id": "001", "Vip": 1,"Points":100,"Coins":1000}`)

	req, err := http.NewRequest("POST", "/pay", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(pay)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
