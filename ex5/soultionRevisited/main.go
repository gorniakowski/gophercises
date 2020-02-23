package main

import (
	"flag"
	"fmt"
	"gophercises/ex4/HtmlLinkParser"
	"net/http"
)

func main() {
	urlFlag := flag.String("url", "http://gophercises.com", "url of a site you wan to build map")
	flag.Parse()
	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	links, _ := HtmlLinkParser.ParseHtml(resp.Body)

	for _, l := range links {
		fmt.Println(l)
	}

}
