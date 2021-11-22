package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var b bytes.Buffer
	var cnt int
	n := len(s)
	if n <= 3 {
		return s
	}
	for _, w := range s {
		cnt++
		if cnt > 3 {
			b.WriteString(",")
			cnt = 1
		}
		b.WriteRune(w)
	}
	return b.String()
}
