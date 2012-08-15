package main

import (
	"fmt"
	"gopost/content/post"
	"html/template"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl/index.html")
	p := &Post{Title: "Falling Rockets", Body: []byte("Hello, Professor McClure.")}
	if err != nil {
		cwd, _ := os.Getwd()
		fmt.Fprintf(w, "Template not available: %v. Current working dir: %v", err, cwd)
	} else {
		t.Execute(w, p)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
