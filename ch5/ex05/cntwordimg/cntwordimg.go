package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	words, images := countWordsAndImages(doc)
	fmt.Printf("words: %v\nimages: %v\n", words, images)
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	scanner := bufio.NewScanner(strings.NewReader(n.Data))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}
	return words, images
}
