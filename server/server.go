package server

import (
	"gorest/checker"
	"gorest/config"
	"gorest/routes"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func Start() {

	app := config.GetApp()

	port := app.Port
	host := app.Host

	log.Printf("HTTP Server running at %s:%d\n", host, port)

	r := chi.NewMux()
	r.Mount("/api/", routes.Router())

	err := http.ListenAndServe(":"+strconv.Itoa(port), r)
	checker.CheckErr(err)

}
