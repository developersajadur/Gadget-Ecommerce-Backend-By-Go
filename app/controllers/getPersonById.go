package controllers

import (
	"ecommerce/app/database"
	"ecommerce/app/helpers"
	"net/http"
	"strconv"
	
)



func GetPersonById(w http.ResponseWriter, r *http.Request){

	productId := r.PathValue("personId")

	id, err := strconv.Atoi(productId)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid person ID", http.StatusBadRequest)
		return
	}

	for _, person := range database.People{
		if person.ID == id{
			 helpers.HandleEncoder(w, person)
			 return
		}
	}

	http.Error(w, "Person not found", http.StatusNotFound)


}