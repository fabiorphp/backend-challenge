package handler

import (
	"github.com/fabiorphp/backend-challenge/pkg/render"
	"net/http"
)

func Error(w http.ResponseWriter, statusCode int, message string) {
	payload := map[string]string{"message": message}
	render.JSON(w, statusCode, payload)
}
