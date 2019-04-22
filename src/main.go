package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func htmlHandler(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path[1:]
	fmt.Println("string(path)")
	fmt.Println(string(path))
	if strings.HasSuffix(string(path), ".jpg") {
		fmt.Println("jpg")
		fmt.Println(r.URL.Path)
		data, err := ioutil.ReadFile(string(path))
		if err == nil {
			fmt.Println("write")

			w.Header().Set("Content-Disposition", "attachment; filename="+string(path))
			w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
			w.Write(data)
		}
	} else {
		fmt.Println("else")

		fmt.Println(r.URL.Path)
		http.StripPrefix("/", http.FileServer(http.Dir("./contents")))
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
}

func main() {
	http.HandleFunc("/", htmlHandler)

	fmt.Println("Open http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
