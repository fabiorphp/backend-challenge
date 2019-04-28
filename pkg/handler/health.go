package handler

import (
	"github.com/fabiorphp/backend-challenge/pkg/render"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	payload := map[string]bool{"alive": true}

	render.JSON(w, http.StatusOK, payload)
}
