package HtmlLinkParser

import (
	"io"

	"golang.org/x/net/html"
)

type DocumentLinks []Link

type Link struct {
	Href string
	Text string
}

func ParseHtml(input io.Reader) (DocumentLinks, error) {
	nodes, err := html.Parse(input)
	if err != nil {
		return nil, err
	}
	findLinks(nodes)
	return nil, nil
}
func findLinks(nodes *html.Node) []Link {

}
