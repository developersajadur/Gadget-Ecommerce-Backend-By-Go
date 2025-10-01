package controllers

import (
	"ecommerce/app/database"
	"ecommerce/app/helpers"
	"net/http"
)

func GetPersons(w http.ResponseWriter, r *http.Request){

	 helpers.SendResponse(w, database.People, http.StatusOK, "fetched all peoples Successfully")
}