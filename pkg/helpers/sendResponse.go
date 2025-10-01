package helpers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func SendResponse(w http.ResponseWriter, data interface{}, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Success:    true,
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)

}
