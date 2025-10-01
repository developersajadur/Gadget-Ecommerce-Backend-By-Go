package controllers

import (
	"ecommerce/app/database"
	"ecommerce/app/helpers"
	"net/http"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {

	var newPerson database.Person

	err := helpers.HandleDecoder(r, &newPerson)
	if err != nil {
		helpers.SendError(w, err, http.StatusBadRequest, "failed to decode")
		return
	}
	for _, person := range database.People {
		if person.Email == newPerson.Email {
			helpers.SendError(w, nil, http.StatusConflict, "Email already exists")
			return
		}
	}
	len := len(database.People)
	newPerson.ID = len + 1
	database.People = append(database.People, newPerson)
	helpers.SendResponse(w, newPerson, http.StatusCreated, "Person Created Successfully")

}
