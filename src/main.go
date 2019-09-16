package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// MyHandler is a object of this http server
type MyHandler struct {
}

func (MyHandler *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path[1:]
	log.Println(string(path))
	if string(path) == "" {
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
	} else if strings.HasSuffix(string(path), ".ico") {
		return
	} else {
		log.Println("download file")
		data, err := ioutil.ReadFile(string(path))
		if err != nil {
			log.Println(err)
		} else {
			info, err := os.Stat(string(path))
			if err != nil {
				log.Println("there was an error to get file stats.", err)
			}

			w.Header().Set("Content-Disposition", "attachment; filename="+string(path))
			w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
			w.Write(data)
		}
	}
}

func main() {
	http.Handle("/", new(MyHandler))

	fmt.Println("Open http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
