package main

import (
	"bufio"
	"flag"
	"fmt"
	"gobook/ch4/ex02/myhash"
	"log"
	"os"
)

var f string

func init() {
	flag.StringVar(&f, "H", "256", `type of hash "256" or "384" or "512"`)
}

func main() {
	flag.Parse()
	switch f {
	case "256", "384", "512":
	default:
		fmt.Printf("invalid flag: %s\n", f)
		flag.PrintDefaults()
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	bytes, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}
	myhash.Hash(bytes[:len(bytes)-1], f)
}
