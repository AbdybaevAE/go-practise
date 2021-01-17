package main

import (
	"bufio"
	"golang.org/x/net/html"
	"io"
	"net/url"
	"os"
)

type Link struct {
	Abs string
	Href string
	Text string
}


func FilterLinksByDomain(domain string, links []Link) []Link {
	items := make([]Link, 0)
	rootLink, err := url.Parse(domain)
	if err != nil {
		return []Link{}
	}
	rootHost := rootLink.Host
	for _, link:= range links {
		if len(link.Href) == 0 {
			continue
		}
		_, err := url.ParseRequestURI(link.Href)
		if err == nil {
			currLink, e := url.Parse(link.Href)
			if e != nil || currLink.Host != rootHost   {
				continue
			}
			link.Abs = link.Href
			items = append(items, link)
		} else {
			if link.Href[0] == '/' {
				link.Abs = domain + link.Href
			} else {
				link.Abs = domain + "/" + link.Href
			}

		}
	}
	return items
}
func ParseLinks(r io.Reader) []Link {
	doc, err := html.Parse(r)
	if err != nil {
		//...
	}

	linkNodes := searchLinkNodes(doc)
	return buildLinks(linkNodes)
}
func parse(fileName string)([]Link, error)  {
	file, readErr := os.Open(fileName)
	if readErr != nil {
		return nil, readErr
	}
	return ParseLinks(bufio.NewReader(file)),nil


}
func buildLinks(nodes []*html.Node) []Link {
	links := make([]Link, len(nodes))
	for i, node := range nodes {
		links[i] = buildLink(node)
	}
	return links
}
func buildLink(node *html.Node) Link {
	href := ""
	for _, v := range node.Attr {
		if v.Key == "href" {
			href = v.Val
		}
	}
	return Link{Href: href, Text: searchText(node)}

}
func searchText(node *html.Node) string {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if node.Type == html.TextNode {
			return node.Data
		}
	}
	return ""
}
func searchLinkNodes(node *html.Node)[]*html.Node {
	if node.Type == html.ElementNode && node.Data == "a" {
		return []*html.Node { node }
	}
	linkNodes := make([]*html.Node, 0)
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		linkNodes = append(linkNodes, searchLinkNodes(c)...)
	}
	return linkNodes
}

