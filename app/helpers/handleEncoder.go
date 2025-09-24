package helpers

import (
	"encoding/json"
	"net/http"
)

	func HandleEncoder(w http.ResponseWriter, data interface{}){
		encoded := json.NewEncoder(w)
		encoded.Encode(data)
	}