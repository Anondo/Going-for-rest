package server

import (
	"going_rest/checker"
	"going_rest/routes"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

const (
	port = 8080
)

func Start() {
	router := httprouter.New()

	routes.SetBlogRoutes(router)

	log.Println("Server running at localhost: " + strconv.Itoa(port))

	err := http.ListenAndServe(":"+strconv.Itoa(port), router)
	checker.CheckErr(err)

}
