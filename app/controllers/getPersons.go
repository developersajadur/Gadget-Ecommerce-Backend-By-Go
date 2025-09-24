package controllers

import (
	"ecommerce/app/database"
	"ecommerce/app/helpers"
	"net/http"
)

func GetPersons(w http.ResponseWriter, r *http.Request){

	helpers.HandleEncoder(w, database.People)
}