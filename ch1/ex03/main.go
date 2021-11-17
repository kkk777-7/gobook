package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Printf("%.10fs %s\n", time.Since(start).Seconds(), s)

	start2 := time.Now()
	s2 := strings.Join(os.Args[1:], " ")
	fmt.Printf("%.10fs %s\n", time.Since(start2).Seconds(), s2)
}
