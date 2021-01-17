package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main()  {
	parse( "examples/ex4.html")

}
type Link struct {
	Href string
	Text string
}
func parse(fileName string)  {
	file, readErr := os.Open(fileName)
	if readErr != nil {
		panic(readErr)
	}
	r := bufio.NewReader(file)
	doc, err := html.Parse(r)
	if err != nil {
		//...
	}

	linkNodes := searchLinkNodes(doc)
	links := buildLinks(linkNodes)
	fmt.Println(links)
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
