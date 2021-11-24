package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	res := anagram(os.Args[1], os.Args[2])
	fmt.Printf("%s, %s: %t\n", os.Args[1], os.Args[2], res)
}

func anagram(s1 string, s2 string) bool {
	var chk bool
	long, short := difflen(s1, s2)
	buf := bytes.NewBufferString(short)
	buf2 := bytes.NewBufferString(long)
	word := buf2.Bytes()

	for _, r := range buf.Bytes() {
		idx := bytes.IndexByte(word, r)
		if idx == -1 {
			break
		}
		word = append(word[:idx], word[idx+1:]...)
	}
	if len(word) == 0 {
		chk = true
	}
	return chk
}

func difflen(s1 string, s2 string) (string, string) {
	l, s := s1, s2
	if len(s1) < len(s2) {
		l, s = s2, s1
	}
	return l, s
}
