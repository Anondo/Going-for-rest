package routes

import (
	"fmt"
	"gorest/handler"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

/*func SetBlogRoutes(router *httprouter.Router) {
	router.GET("/blogs", handler.BlogHandler)
	router.GET("/blogs/:id", handler.BlogHandler)
	router.POST("/blogs", handler.BlogHandler)
	router.PUT("/blogs/:id", handler.BlogHandler)
	router.DELETE("/blogs/:id", handler.BlogHandler)
	router.PATCH("/blogs/:id", handler.BlogHandler)
}*/

var router = chi.NewRouter()

func Router() http.Handler {
	router.NotFound(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "%s", []byte(`message : Route Not Found`))
	})
	router.MethodNotAllowed(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "%s", []byte(`message : Method Not Allowed`))
	})

	router.Use(middleware.Timeout(60 * time.Second))

	registerRoutes()

	return router

}

func registerRoutes() {
	router.Route("/v1/", func(r chi.Router) {
		r.Mount("/blogs", handler.BlogHandler())
	})
}
