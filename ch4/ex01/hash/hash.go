package hash

import (
	"crypto/sha256"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountDiff(s1, s2 [sha256.Size]byte) int {
	var cnt int
	for i := 0; i < sha256.Size; i++ {
		diff := s1[i] ^ s2[i]
		cntdiff := pc[diff]
		cnt += int(cntdiff)
	}
	return cnt
}
