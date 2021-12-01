package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "elecount: %v\n", err)
		os.Exit(1)
	}
	cnts := make(map[string]int)
	elecount(doc, cnts)
	for tag, cnt := range cnts {
		fmt.Printf("%s: %v\n", tag, cnt)
	}
}

func elecount(n *html.Node, cnt map[string]int) {
	if n.Type == html.ElementNode {
		cnt[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		elecount(c, cnt)
	}
}
