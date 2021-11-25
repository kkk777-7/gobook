package main

import (
	"crypto/sha256"
	"fmt"
	"gobook/ch4/ex01/hash"
)

func main() {
	sum1 := sha256.Sum256([]byte("hello world\n"))
	sum2 := sha256.Sum256([]byte("hello world2\n"))

	res := hash.PopCountDiff(sum1, sum2)
	fmt.Println(res)
}
