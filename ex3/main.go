package main

import (
	"fmt"
	"gophercises/ex3/parser"
	"html/template"
	"net/http"
	"os"
)

const fileName = "gopher.json"

func main() {
	arcs, err := parser.ParseJson(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", arcHandler(arcs))

}

func arcHandler(arcs map[string]parser.Arc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var arc parser.Arc
		if r.URL.Path[1:] == "" {
			arc = arcs["intro"]
		} else {
			arc = arcs[r.URL.Path[1:]]
		}
		t, _ := template.ParseFiles("template.html")
		t.Execute(w, arc)

	}
}
