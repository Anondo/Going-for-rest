package view

import (
	"going_rest/conn"
	"going_rest/model"
	"going_rest/serializer"
)

func BlogPost(blog model.Blog) {
	conn.DB.Create(&blog)
	conn.DB.NewRecord(blog)
}

func BlogGet(id ...int) []byte {
	if len(id) == 0 {
		var blog []model.Blog
		conn.DB.Find(&blog)
		return serializer.BlogSerialize(blog)
	} else {
		var blog model.Blog
		conn.DB.First(&blog, id[0])
		if blog.ID == 0 {
			return serializer.BlogSerialize(struct{ Detail string }{"Not Found"})
		}
		return serializer.BlogSerialize(blog)
	}
}

func BlogPut(id int, updatedBlog model.Blog) []byte {
	var blog model.Blog
	conn.DB.First(&blog, id)
	blog.Title = updatedBlog.Title
	blog.Body = updatedBlog.Body
	conn.DB.Save(&blog)
	return BlogGet(id)
}
func BlogDelete(id int) []byte {
	var blog model.Blog
	conn.DB.First(&blog, id).Delete(&blog)
	if blog.ID == 0 {
		return serializer.BlogSerialize(struct{ Detail string }{"Not Found"})
	}
	return []byte(`[]`)
}

func BlogPatch(id int, updtBlg model.Blog) []byte {
	var blog model.Blog
	conn.DB.First(&blog, id)
	if updtBlg.Title != "" {
		blog.Title = updtBlg.Title
	}
	if updtBlg.Body != "" {
		blog.Body = updtBlg.Body
	}
	conn.DB.Save(&blog)
	return BlogGet(id)
}
