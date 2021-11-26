package myhash

import (
	"crypto/sha256"
	"fmt"
)

func Hash(b []byte) {

	sum := sha256.Sum256(b)
	fmt.Printf("%x\n", sum)
}
