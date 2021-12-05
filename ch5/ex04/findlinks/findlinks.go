package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	link := make(map[string][]string)
	visit(link, doc)

	for name, urls := range link {
		for _, url := range urls {
			fmt.Printf("%s: %s\n", name, url)
		}
	}
}

func visit(links map[string][]string, n *html.Node) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			for _, a := range n.Attr {
				if a.Key == "href" {
					links[n.Data] = append(links[n.Data], a.Val)
				}
			}
		case "img":
			for _, img := range n.Attr {
				if img.Key == "src" {
					links[n.Data] = append(links[n.Data], img.Val)
				}
			}
		case "script":
			for _, sc := range n.Attr {
				if sc.Key == "src" {
					links[n.Data] = append(links[n.Data], sc.Val)
				}
			}
		case "link":
			for _, ss := range n.Attr {
				if ss.Key == "href" {
					links[n.Data] = append(links[n.Data], ss.Val)
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(links, c)
	}
}
