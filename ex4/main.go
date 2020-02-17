package main

import (
	"flag"
	"fmt"
	"gophercises/ex4/HtmlLinkParser"
	"os"
)

func main() {
	var result []HtmlLinkParser.Link
	fileName := flag.String("f", "ex1.html", "html file to parse")
	flag.Parse()
	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	result, err = HtmlLinkParser.ParseHtml(file)
	if err != nil {
		panic(err)
	}
	for _, link := range result {
		fmt.Printf("%+v", link)
	}
}
