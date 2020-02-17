package HtmlLinkParser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type DocumentLinks []Link

type Link struct {
	Href string
	Text string
}

func ParseHtml(input io.Reader) (DocumentLinks, error) {
	node, err := html.Parse(input)
	if err != nil {
		return nil, err
	}
	linkNodes := findLinkNodes(node)
	var links DocumentLinks
	for _, node := range linkNodes {
		links = append(links, buildLink(node))
	}

	return links, nil
}
func findLinkNodes(node *html.Node) []*html.Node {
	if node.Type == html.ElementNode && node.Data == "a" {
		return []*html.Node{node}
	}
	var result []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		result = append(result, findLinkNodes(c)...)
	}

	return result
}

func buildLink(node *html.Node) Link {
	var result Link
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			result.Href = attr.Val
			break
		}
	}
	result.Text = findText(node)
	return result
}

func findText(node *html.Node) string {
	if node.Type == html.TextNode {
		return node.Data
	}
	if node.Type != html.ElementNode {
		return ""
	}
	var result string
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		result += findText(c)
	}
	return strings.Join(strings.Fields(result), " ")

}
