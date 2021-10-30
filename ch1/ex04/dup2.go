package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	fileNames := make(map[string]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, fileNames, "Stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileNames, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Printf("%s\n", fileNames[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileNames map[string]string, name string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++

		if counts[input.Text()] == 1 {
			fileNames[input.Text()] = name
		} else {
			if !strings.Contains(fileNames[input.Text()], name) {
				fileNames[input.Text()] += " " + name
			}
		}
	}
	fmt.Println(fileNames)
}
