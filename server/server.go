package server

import (
	"gorest/checker"
	"gorest/config"
	"gorest/routes"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func Start() {

	app := config.GetApp()

	port := app.Port
	host := app.Host

	router := httprouter.New()

	routes.SetBlogRoutes(router)

	log.Printf("HTTP Server running at %s:%d\n", host, port)

	err := http.ListenAndServe(":"+strconv.Itoa(port), router)
	checker.CheckErr(err)

}
