package handler

import (
	"fmt"
	"gorest/checker"
	"gorest/model"
	"gorest/serializer"
	"gorest/view"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func BlogHandler(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if req.Method == "GET" {
		if len(ps) > 0 {
			id, err := strconv.Atoi(ps.ByName("id"))
			checker.CheckErr(err)
			fmt.Fprintf(rw, "%s", view.BlogGet(id))
		} else {
			fmt.Fprintf(rw, "%s", view.BlogGet())
		}

	}
	if req.Method == "POST" {
		var blog model.Blog
		serializer.BlogDeserialize(&blog, req)
		view.BlogPost(blog)
		fmt.Fprintf(rw, "%s", []byte(`[]`))
	}
	if req.Method == "PUT" {
		var blog model.Blog
		serializer.BlogDeserialize(&blog, req)
		id, err := strconv.Atoi(ps.ByName("id"))
		checker.CheckErr(err)
		fmt.Fprintf(rw, "%s", view.BlogPut(id, blog))
	}
	if req.Method == "DELETE" {
		id, err := strconv.Atoi(ps.ByName("id"))
		checker.CheckErr(err)
		fmt.Fprintf(rw, "%s", view.BlogDelete(id))
	}
	if req.Method == "PATCH" {
		var blog model.Blog
		serializer.BlogDeserialize(&blog, req)
		id, err := strconv.Atoi(ps.ByName("id"))
		checker.CheckErr(err)
		fmt.Fprintf(rw, "%s", view.BlogPatch(id, blog))
	}
}
