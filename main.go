package main

import (
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mwelwankuta/goweb/routes"
)

func main() {

	app.Route("/", &routes.Home{Loading: true})
	app.Script().Src("https://cdn.tailwindcss.com")

	app.RunWhenOnBrowser()
	http.Handle("/", &app.Handler{
		Name:        "Mwelwa's App",
		Description: "Making an app",
	})

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
