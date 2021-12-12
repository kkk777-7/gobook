package main

import (
	"fmt"
	"gobook/ch7/ex01/counter"
)

func main() {
	var p counter.WordCounter

	fmt.Fprintf(&p, "Hello World Test")
	fmt.Println(p)

	var p2 counter.LineCounter
	fmt.Fprintf(&p2, "Hello World Test\nHello World Test2\n")
	fmt.Println(p2)
}
