package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var depth int

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)
	return nil
}

func forEachNode(n *html.Node, startElement, endElement func(n *html.Node)) {
	if startElement != nil {
		startElement(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, startElement, endElement)
	}

	if endElement != nil {
		endElement(n)
	}
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		nodeprint(n)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func nodeprint(n *html.Node) {
	if n.Type == html.ElementNode {
		if len(n.Attr) > 0 {
			a := n.Attr
			fmt.Printf(" %s='%s'>\n", a[0].Key, a[0].Val)
		} else {
			fmt.Printf(">\n")
		}
	} else {
		fmt.Printf(">\n")
	}
}
