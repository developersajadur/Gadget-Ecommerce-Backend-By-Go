package main

import (
	"ecommerce/app"
	"ecommerce/app/database"
)

func init() {

	database.HandleInit()
	
}

func main() {
	app.RunServer()
}
