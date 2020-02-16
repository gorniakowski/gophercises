package main

import (
	"flag"
	"gophercises/ex4/HtmlLinkParser"
	"os"
)

func main() {
	fileName := flag.String("f", "ex1.html", "html file to parse")
	flag.Parse()
	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	_, err = HtmlLinkParser.ParseHtml(file)
	if err != nil {
		panic(err)
	}

}
