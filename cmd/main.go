package main

import (
	"log"

	"github.com/FlyKarlik/news_skilfactory/internal/app"
)

func main() {
	app := app.NewApp()

	if err := app.RunApp(); err != nil {
		log.Fatal(err)
	}
}
