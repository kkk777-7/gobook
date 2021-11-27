package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please input files")
		os.Exit(1)
	}

	files := os.Args[1:]
	for _, file := range files {
		counts := make(map[string]int)
		data, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()

		scanner := bufio.NewScanner(data)
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			counts[scanner.Text()]++
		}

		var names []string
		for name := range counts {
			names = append(names, name)
		}
		sort.Strings(names)

		for _, name := range names {
			fmt.Printf("%s: %v\n", name, counts[name])
		}
	}
}
