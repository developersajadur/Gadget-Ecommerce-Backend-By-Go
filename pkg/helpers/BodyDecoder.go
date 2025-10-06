package helpers

import (
	"encoding/json"
	"net/http"
)

func BodyDecoder(w http.ResponseWriter, r *http.Request, data interface{}) error {
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		SendError(w, err, http.StatusBadRequest, "Can't decode Data from body")
		return err

	}
	return nil
}
