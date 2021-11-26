package main

import (
	"bufio"
	"gobook/ch4/ex02/myhash"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	bytes, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}

	myhash.Hash(bytes[:len(bytes)-1])
}
