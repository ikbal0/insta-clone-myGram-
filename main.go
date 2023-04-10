package main

import (
	"insta-clone/database"
	"insta-clone/routers"
)

func main() {
	database.StartDB()

	var PORT = "8080"
	r := routers.StartApp()
	r.Run(":" + PORT)
}
