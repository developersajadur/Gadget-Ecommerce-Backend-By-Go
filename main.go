package main

import (
	"ecommerce/app"
	"ecommerce/app/database"
)

// init runs automatically before main()
func init() {

	database.HandleInit()
	
}

func main() {
	utils.RunServer()
}
