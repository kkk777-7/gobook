package intset

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	Words []uint64
}

// Has reports Whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	Word, bit := x/64, uint(x%64)
	return Word < len(s.Words) && s.Words[Word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	Word, bit := x/64, uint(x%64)
	for Word >= len(s.Words) {
		fmt.Printf("%b\n", s.Words)
		s.Words = append(s.Words, 0)
		fmt.Printf("%b\n", s.Words)
	}
	s.Words[Word] |= 1 << bit // s.Words[Word] = s.Words[Word] | 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tWord := range t.Words {
		if i < len(s.Words) {
			s.Words[i] |= tWord
		} else {
			s.Words = append(s.Words, tWord)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, Word := range s.Words {
		if Word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if Word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
