package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func htmlHandler(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("index.html"))

	m := map[string]string{
		"title": "golang file server",
		"text":  "this page will be a simple file server",
	}

	if err := t.ExecuteTemplate(w, "index.html", m); err != nil {
		log.Fatal(err)
	}
}

func serveimage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "images/gopher.jpg")
}

func main() {
	http.HandleFunc("/", htmlHandler)
	http.HandleFunc("/gopher.jpg", serveimage)

	fmt.Print("Open http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
