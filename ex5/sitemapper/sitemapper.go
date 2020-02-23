package sitemapper

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"gophercises/ex4/HtmlLinkParser"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type siteMap map[string]bool

func GetSite(url string) siteMap {
	webpageMap := make(siteMap)
	domain, err := domainExtractor(url)
	if err != nil {
		fmt.Println(err)
	}
	err = getPage(url, webpageMap, domain)
	if err != nil {
		fmt.Println(err)
	}

	return webpageMap
}

func getPage(url string, webpageMap siteMap, domain string) error {
	var siteBody []byte
	var counter int
	site, err := http.Get(url)
	if err != nil {
		return err
	}
	defer site.Body.Close()
	siteBody, err = ioutil.ReadAll(site.Body)
	bodyReader := bytes.NewReader(siteBody)
	counter, err = collectLinks(bodyReader, webpageMap, domain)
	if counter == 0 {
		return nil
	}
	for link := range webpageMap {
		getPage(link, webpageMap, domain)
	}
	return nil
}

func collectLinks(site io.Reader, webpageMap siteMap, domain string) (int, error) {
	result, err := HtmlLinkParser.ParseHtml(site)
	if err != nil {
		return 0, err
	}
	counter := 0
	for _, link := range result {
		if webpageMap[link.Href] == false && filter(link.Href, domain) {
			fmt.Println(link.Href)
			webpageMap[link.Href] = true
			counter++
		}

	}
	return counter, nil

}

func filter(link string, domain string) bool {
	if strings.HasPrefix(link, "/") {
		return true
	}
	if strings.Contains(link, domain) {
		return true
	}
	return false

}

func domainExtractor(url string) (string, error) {
	if strings.HasPrefix(url, "https://www.") {
		return url[12:], nil
	}
	if strings.HasPrefix(url, "http://www.") {
		return url[11:], nil
	}
	if strings.HasPrefix(url, "https://") {
		return url[8:], nil
	}
	if strings.HasPrefix(url, "http://") {
		return url[7:], nil
	}
	return "", errors.New("canot parse url to get domain")

}

const (
	Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
)

func GenerateXml(input siteMap) string {

	type Link struct {
		URL string `xml:"loc"`
	}
	type webMap struct {
		XMLName  xml.Name `xml:"urlset"`
		AllLinks []Link   `xml:"url"`
	}

	var links webMap
	for link := range input {
		url := Link{URL: link}
		links.AllLinks = append(links.AllLinks, url)
	}

	result, err := xml.MarshalIndent(links, " ", "  ")
	if err != nil {
		fmt.Println(err)
	}

	return Header + string(result)
}
