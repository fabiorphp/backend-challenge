package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/health", nil)

	if err != nil {
		t.Fail()
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(Health)

	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned invalid status code: got %v want %v",
			status,
			http.StatusOK,
		)
	}

	res := make(map[string]bool)

	if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
		t.Fail()
	}

	if res["alive"] != true {
		t.Errorf("Invalid health check status")
	}
}
