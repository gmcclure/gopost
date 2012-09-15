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
	posts := post.GetAll()
	t.Execute(w, posts)
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8000", nil)
}
