package main

import (
	"flag"
	"gophercises/ex4/HtmlLinkParser"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type loc struct {
	Value string `xml:"loc"`
}
type urlset struct {
	Urls []loc `xml:"url"`
}

func main() {
	urlFlag := flag.String("url", "http://gophercises.com", "url of a site you wan to build map")
	maxDepth := flag.Int("depth", 10, "The maximum number of links to traverse")
	flag.Parse()
	pages := bfs(*urlFlag, *maxDepth)

	for _, page := range pages {
		println(page)
	}
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()
	return filter(hrefs(resp.Body, base), withPrefix(base))

}

func bfs(urlStr string, maxDepth int) []string {
	seen := make(map[string]struct{})
	var q map[string]struct{}
	nq := map[string]struct{}{
		urlStr: struct{}{},
	}

	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})
		for url := range q {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = struct{}{}

			for _, link := range get(url) {
				nq[link] = struct{}{}
			}
		}
	}
	var result []string
	for url := range seen {
		result = append(result, url)
	}
	return result
}

func hrefs(r io.Reader, base string) []string {

	links, _ := HtmlLinkParser.ParseHtml(r)
	var result []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			result = append(result, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			result = append(result, l.Href)

		}
	}
	return result
}

func filter(links []string, keepFn func(string) bool) []string {
	var result []string
	for _, link := range links {
		if keepFn(link) {
			result = append(result, link)
		}
	}
	return result
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}
