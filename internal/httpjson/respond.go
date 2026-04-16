package httpjson

import (
	"encoding/json"
	"net/http"
)

type ResponseBody struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type RB = ResponseBody

func Respond(w http.ResponseWriter, response ResponseBody, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(response)
}
