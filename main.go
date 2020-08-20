package main

import (
	"log"
	"net/http"

	"webengine/controllers"
)

func main() {
	api := controllers.NewAPI()

	if err := http.ListenAndServe(":8080", api); err != nil {
		log.Println(err)
	}
}
