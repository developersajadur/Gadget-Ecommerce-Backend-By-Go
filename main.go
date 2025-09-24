package main

import (
	"ecommerce/app/database"
	"ecommerce/app"
)

// init runs automatically before main()
func init() {
	database.HandleInit()
}

func main() {
	utils.RunServer()
}
