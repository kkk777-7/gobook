package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var sign string
	for i := 1; i < len(os.Args); i++ {
		num := os.Args[i]
		num, sign = IsSign(num)
		decimal := strings.Index(num, ".")
		if decimal > 0 {
			fmt.Printf("  %s\n", sign+comma(num[:decimal])+num[decimal:])
		} else {
			fmt.Printf("  %s\n", sign+comma(num))
		}
	}
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func IsSign(w string) (string, string) {
	if w[0:1] == "+" || w[0:1] == "-" {
		sign := w[:1]
		w = w[1:]
		return w, sign
	}
	return w, ""
}
