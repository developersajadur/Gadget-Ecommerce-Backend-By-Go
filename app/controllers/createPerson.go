package controllers

import (
	"ecommerce/app/database"
	"ecommerce/app/helpers"
	"net/http"
)


func CreatePerson(w http.ResponseWriter, r *http.Request){

	var newPerson database.Person

	err:= helpers.HandleDecoder(r, &newPerson)
	if err!= nil{
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	database.People = append(database.People, newPerson)
	w.WriteHeader(http.StatusCreated)
	helpers.HandleEncoder(w, newPerson)

}