package serializer

import (
	"encoding/json"
	"going_rest/checker"
	"going_rest/model"
	"net/http"
)

func BlogDeserialize(blog *model.Blog, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&blog)
	checker.CheckErr(err)
}
func BlogSerialize(i interface{}) []byte {
	json_bytes, err := json.Marshal(i)
	checker.CheckErr(err)
	return json_bytes
}
