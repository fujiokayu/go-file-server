package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type renderData struct {
	Title string
}

func htmlHandler0(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))

	m := map[string]string{
		"title": "golang file server",
		"text":  "this page will be a simple file server",
	}

	// テンプレートを描画
	if err := t.ExecuteTemplate(w, "index.html", m); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", htmlHandler0)

	fmt.Print("Open http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
