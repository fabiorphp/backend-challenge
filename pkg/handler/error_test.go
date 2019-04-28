package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestErrorMessage(t *testing.T) {
	rec := httptest.NewRecorder()

	Error(rec, http.StatusNotFound, "not found")

	if status := rec.Code; status != http.StatusNotFound {
		t.Errorf(
			"handler returned invalid status code: got %v want %v",
			status,
			http.StatusOK,
		)
	}

	res := make(map[string]string)

	if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
		t.Fail()
	}

	if res["message"] != "not found" {
		t.Error("Invalid body response")
	}
}
