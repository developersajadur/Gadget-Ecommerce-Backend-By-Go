package helpers

import (
	"encoding/json"
	"net/http"
)

// Error defines the structure for error responses
type Error struct {
	Success    bool        `json:"status"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Error      interface{} `json:"error,omitempty"`
}

func SendError(w http.ResponseWriter, err interface{}, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	response := Error{
		Success:    false,
		StatusCode: statusCode,
		Message:    message,
		Error:      err,
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
