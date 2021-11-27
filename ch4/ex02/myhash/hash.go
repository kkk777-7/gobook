package myhash

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func Hash(b []byte, s string) {
	switch s {
	case "256":
		sum := sha256.Sum256(b)
		fmt.Printf("%x\n", sum)
	case "384":
		sum := sha512.Sum384(b)
		fmt.Printf("%x\n", sum)
	case "512":
		sum := sha512.Sum512(b)
		fmt.Printf("%x\n", sum)
	}
}
