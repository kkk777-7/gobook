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
	word, bit := x/64, uint(x%64)
	return word < len(s.Words) && s.Words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.Words) {
		s.Words = append(s.Words, 0)
	}
	s.Words[word] |= 1 << bit // s.Words[Word] = s.Words[Word] | 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.Words {
		if i < len(s.Words) {
			s.Words[i] |= tword
		} else {
			s.Words = append(s.Words, tword)
		}
	}
}

// IntersectWith sets s to the intersect of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.Words {
		if i < len(s.Words) {
			s.Words[i] &= tword
		}
	}
}

// SymmetricDifference sets s to the SymmetricDifference of s and t.
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.Words {
		if i < len(s.Words) {
			s.Words[i] ^= tword
		} else {
			s.Words = append(s.Words, tword)
		}
	}
}

// DifferenceWith sets s to the difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.Words {
		if i < len(s.Words) {
			s.Words[i] &^= tword
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.Words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
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

// Len returns the number of elements.
func (s *IntSet) Len() int {
	num := len(s.Words)
	return num
}

// Remove removes x from the set.
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.Words[word] &= 0 << bit
}

// Clear removes all elements from the set.
func (s *IntSet) Clear() {
	for i := range s.Words {
		for j := 0; j < 64; j++ {
			s.Words[i] &= 0 << uint(j)
		}
	}
}

// Copy returns a copy of the set.
func (s *IntSet) Copy() *IntSet {
	res := s
	return res
}

// Addall adds the non-negative values args to the set.
func (s *IntSet) AddAll(args ...int) {
	for _, arg := range args {
		s.Add(arg)
	}
}

// Elems returns slice of elements of Intset
func (s *IntSet) Elems() []int {
	var res []int
	for num := range s.Words {
		for i := 0; i < 64; i++ {
			if s.Words[num]&(1<<uint(i)) != 0 {
				res = append(res, num*64+i)
			}
		}
	}
	return res
}
