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
			helpers.SendError(w, err, http.StatusBadRequest, "Invalid person ID")
		return
	}

	for _, person := range database.People{
		if person.ID == id{
			 helpers.SendResponse(w, person, http.StatusOK, "fetched the person Successfully")
			 return
		}
	}
		helpers.SendError(w, nil, http.StatusNotFound, "Person not found")
		


}