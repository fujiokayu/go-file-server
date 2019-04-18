package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func htmlHandler(w http.ResponseWriter, r *http.Request) {

	http.StripPrefix("/", http.FileServer(http.Dir("/static/")))
	t := template.Must(template.ParseFiles("static/index.html"))

	files, _ := ioutil.ReadDir("./contents")
	fileNames := []string{}
	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}

	if err := t.ExecuteTemplate(w, "index.html", fileNames); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", htmlHandler)

	fmt.Print("Open http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
