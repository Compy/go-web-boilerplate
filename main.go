package main

import "go-web-boilerplate/database"

func main() {
	database.Connect("boilerplate:boilerplate@tcp(localhost:3306)/boilerplate?parseTime=true")
	database.Migrate()

	router := initRouter()
	router.Run(":8080")
}
