package render

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, v interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	if v == nil {
		return
	}

	if statusCode == http.StatusNoContent {
		return
	}

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)

	err = enc.Encode(v)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return err
}
