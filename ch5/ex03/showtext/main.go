package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "elecount: %v\n", err)
		os.Exit(1)
	}
	searchTextNode(nil, doc)

}

func searchTextNode(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
	}
	if n.Type == html.TextNode {
		chk := stack[len(stack)-1]

		if chk != "script" && chk != "style" {
			text := strings.TrimSpace(n.Data)
			if len(text) > 0 {
				fmt.Printf("<%s>%s</%s>\n", chk, text, chk)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		searchTextNode(stack, c)
	}
}
