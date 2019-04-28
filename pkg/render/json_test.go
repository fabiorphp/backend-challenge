package render

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestJSONRenderWithoutPayload(t *testing.T) {
	rec := httptest.NewRecorder()

	JSON(rec, http.StatusOK, nil)

	body, err := ioutil.ReadAll(rec.Body)

	if err != nil {
		t.Fatal()
	}

	if len(body) > 0 {
		t.Error("body muste be empty")
	}

	payload := map[string]bool{"status": true}

	JSON(rec, http.StatusNoContent, payload)

	body, err = ioutil.ReadAll(rec.Body)

	if err != nil {
		t.Fatal()
	}

	if len(body) > 0 {
		t.Error("body must be empty")
	}
}

func TestJSONRender(t *testing.T) {
	rec := httptest.NewRecorder()

	payload := map[string]bool{"status": true}

	JSON(rec, http.StatusOK, payload)

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

	if res["status"] != payload["status"] {
		t.Error("Invalid body response")
	}
}

func TestJSONRenderReturnsError(t *testing.T) {
	rec := httptest.NewRecorder()

	payload := map[string]interface{}{"status": math.Inf(-1)}

	JSON(rec, http.StatusOK, payload)

	body, err := ioutil.ReadAll(rec.Body)

	if err != nil {
		t.Fatal()
	}

	if !strings.Contains(string(body), "unsupported") {
		t.Error("Invalid body response")
	}
}
