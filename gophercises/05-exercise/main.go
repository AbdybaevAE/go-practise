package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)
const currLevel = 10
//?>
//<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
//  <url>
//    <loc>http://www.example.com/</loc>
//  </url>
//  <url>
//    <loc>http://www.example.com/dogs</loc>
//  </url>
//</urlset>
type UrlXml struct {
	LocationXml string `xml:"urlset:url:loc"`
}
type xmlStr struct {
	singUrl []UrlXml `xml:"urlset:url"`
}
const link =
	"https://www.mediawiki.org/wiki/Parser_tests"
func main() {
	myMap := make(map[string]bool)
	ProcessUrl(link, myMap, currLevel)
	allLinks := make([]string, 0)
	for key, _:= range myMap  {
		allLinks = append(allLinks, key)
	}
	file, _ := xml.MarshalIndent(allLinks, "", " ")
	_ = ioutil.WriteFile("notes1.xml", file, 0644)

}
func ProcessUrl(currUrl string, visited map[string]bool, level int) {
	if visited[currUrl] || level == 0 {
		return
	}
	visited[currUrl] = true
	resp, err := http.Get(currUrl)
	if err != nil {
		fmt.Printf("cannot load %s %e", currUrl, err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	links := ParseLinks(bytes.NewReader(body))
	currU, e := url.Parse(currUrl)
	if e != nil {
		return
	}
	domain := ""
	if currU.Scheme != "" {
		domain += currU.Scheme + "://"
	}
	domain += currU.Host
	filteredLinks := FilterLinksByDomain(domain, links)

	for _, v := range filteredLinks {
		if visited[v.Abs] {
			continue
		}
		fmt.Println(v.Abs)
		ProcessUrl(v.Abs, visited, level - 1)
		_=v
	}

}
