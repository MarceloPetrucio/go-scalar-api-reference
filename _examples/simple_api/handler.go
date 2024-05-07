package main

import (
	"encoding/json"
	"net/http"
)

type ResponseSuccess struct {
	Message string `json:"message,omitempty"`
}

// Exemple Doc
// @Summary      Get Hello
// @Description  Simple get hello word
// @Tags         Hello
// @Produce      json
// @Success      200				 {object}  ResponseSuccess
// @Router       / [get]
func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := ResponseSuccess{
		Message: "Hello Word",
	}

	json.NewEncoder(w).Encode(response)
}
