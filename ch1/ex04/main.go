package main

import (
	"bufio"
	"fmt"
	"os"
)

type chk struct {
	cnt   int
	files []string
}

func main() {
	counts := make(map[string]chk)

	args := os.Args[1:]

	if len(args) == 0 {
		countLines(os.Stdin, counts, "stdin")
	} else {
		for _, arg := range args {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, c := range counts {
		if c.cnt > 1 {
			fmt.Printf("%d\t%s\n", c.cnt, line)
			fmt.Printf("%v\n", c.files)
		}
	}
}

func countLines(f *os.File, counts map[string]chk, name string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if _, ok := counts[input.Text()]; !ok {
			c := chk{cnt: 1, files: []string{name}}
			counts[input.Text()] = c
		} else {
			v, _ := counts[input.Text()]
			v.cnt++
			if !contain(name, v.files) {
				v.files = append(v.files, name)
			}
			counts[input.Text()] = v
		}
	}
}

func contain(name string, files []string) bool {
	for _, v := range files {
		if name == v {
			return true
		}
	}
	return false
}
