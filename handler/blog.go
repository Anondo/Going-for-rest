package handler

import (
	"fmt"
	"gorest/checker"
	"gorest/model"
	"gorest/serializer"
	"gorest/view"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func listBlogs(rw http.ResponseWriter, r *http.Request) {
	rw.Write(view.BlogGet())
}
func listBlog(rw http.ResponseWriter, r *http.Request) {
	istr := chi.URLParam(r, "id")
	i, err := strconv.Atoi(istr)
	checker.CheckErr(err)
	fmt.Fprintf(rw, "%s", view.BlogGet(i))
}
func createBlog(rw http.ResponseWriter, r *http.Request) {
	var blog model.Blog
	serializer.BlogDeserialize(&blog, r)
	view.BlogPost(blog)
	fmt.Fprintf(rw, "%s", []byte(`[]`))
}
func updateBlogPut(rw http.ResponseWriter, r *http.Request) {
	var blog model.Blog
	serializer.BlogDeserialize(&blog, r)
	istr := chi.URLParam(r, "id")
	i, err := strconv.Atoi(istr)
	checker.CheckErr(err)
	fmt.Fprintf(rw, "%s", view.BlogPut(i, blog))
}
func removeBlog(rw http.ResponseWriter, r *http.Request) {
	istr := chi.URLParam(r, "id")
	i, err := strconv.Atoi(istr)
	checker.CheckErr(err)
	fmt.Fprintf(rw, "%s", view.BlogDelete(i))
}
func updateBlogPatch(rw http.ResponseWriter, r *http.Request) {
	var blog model.Blog
	serializer.BlogDeserialize(&blog, r)
	istr := chi.URLParam(r, "id")
	i, err := strconv.Atoi(istr)
	checker.CheckErr(err)
	fmt.Fprintf(rw, "%s", view.BlogPatch(i, blog))
}

func BlogHandler() http.Handler {
	h := chi.NewRouter()
	h.Group(func(r chi.Router) {
		r.Get("/", listBlogs)
		r.Get("/{id}", listBlog)
		r.Post("/", createBlog)
		r.Put("/{id}", updateBlogPut)
		r.Patch("/{id}", updateBlogPatch)
		r.Delete("/{id}", removeBlog)
	})
	return h
}
