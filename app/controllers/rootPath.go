package controllers

import "net/http"


func RootPath(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome to the Person API! Use /persons to get the list of persons."))
}
