package routes

import (
	"gorest/handler"

	"github.com/julienschmidt/httprouter"
)

func SetBlogRoutes(router *httprouter.Router) {
	router.GET("/blogs", handler.BlogHandler)
	router.GET("/blogs/:id", handler.BlogHandler)
	router.POST("/blogs", handler.BlogHandler)
	router.PUT("/blogs/:id", handler.BlogHandler)
	router.DELETE("/blogs/:id", handler.BlogHandler)
	router.PATCH("/blogs/:id", handler.BlogHandler)
}
