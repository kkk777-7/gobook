package main

import (
	"fmt"
	"gobook/ch6/ex01/intset"
)

func main() {
	var a intset.IntSet
	a.Add(10)
	a.Add(64)
	a.Add(128)
	a.Add(64)

	fmt.Printf("%b\n", a)

	fmt.Println(len(a.Words))

}
