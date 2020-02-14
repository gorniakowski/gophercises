package main

import (
	"flag"
	"fmt"
	"gophercises/ex3/parser"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	fileName := flag.String("file", "gopher.json", "you give file with story")
	flag.Parse()
	arcs, err := parser.ParseJson(*fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", arcHandler(arcs))

}

func arcHandler(arcs map[string]parser.Arc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("template.html")
		path := strings.TrimSpace(r.URL.Path)
		if path == "" || path == "/" {
			path = "/intro"
		}

		if arc, ok := arcs[path[1:]]; ok {
			err := t.Execute(w, arc)
			if err != nil {
				log.Printf("%v", err)
				http.Error(w, "something went wrong", http.StatusInternalServerError)
			}
			return
		}
		http.Error(w, "Chapter not found", http.StatusNotFound)

	}
}
