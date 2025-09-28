package main

import (
	"ecommerce/app/database"
	"ecommerce/app"
)

func init() {
	database.HandleInit()
}

func main() {
	utils.RunServer()
}
