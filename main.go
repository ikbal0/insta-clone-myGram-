package main

import (
	"insta-clone/database"
	"insta-clone/internals/protocols/http/routers"
)

func main() {
	database.StartDB()

	var PORT = "8082"
	r := routers.StartApp()
	r.Run(":" + PORT)
}
