package helpers

import (
	"encoding/json"
	"net/http"
)


	func HandleDecoder(r *http.Request, data interface{}) error{
		err:= json.NewDecoder(r.Body).Decode(data)
		if err!= nil{
			return err
		}
		return nil
	}