package controllers

import (
	"ecommerce/app/helpers"
	"net/http"
)


func RootPath(w http.ResponseWriter, r *http.Request){
		 helpers.SendResponse(w, nil, http.StatusOK, "Welcome Home")
}
