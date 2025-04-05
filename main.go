package main

import (
	"habit-tracker-api/internal/app"
	"log"
)

func main() {
	app := app.NewApp()
	log.Fatal(app.Run(":8080"))
}
