package main

import (
	"gopost/content/post"
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl/index.html")
	if err != nil {
		panic(err)
	}
	ctx := make(map[string]*[]post.Post)
	ctx["posts"] = post.GetAll()
	t.Execute(w, ctx)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
