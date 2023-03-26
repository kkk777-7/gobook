package main

import (
	"strings"
)

func main() {
	echoBad()
	echoGood()
}

func echoBad() {
	var s, sep string
	data := []string{"Hello", "World", "!"}
	for _, v := range data {
		s += sep + v
		sep = " "
	}
}

func echoGood() {
	data := []string{"Hello", "World", "!"}
	strings.Join(data, " ")
}
