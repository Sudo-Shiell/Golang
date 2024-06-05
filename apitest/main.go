package main

import (
	"log"
	"api/app"
)

func main() {
	router := app.SetRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}