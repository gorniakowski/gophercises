package main

import (
	"flag"
	"fmt"
	"gophercises/ex5/sitemapper"
)

func main() {
	domain := flag.String("d", "https://www.calhoun.io", "domain to generate site map")
	flag.Parse()
	webmap := sitemapper.GetSite(*domain)
	fmt.Print(sitemapper.GenerateXml(webmap))

}
