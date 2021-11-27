package reverse_test

import (
	"gobook/ch4/ex03/reverse"
	"testing"
)

func TestReverseDo(t *testing.T) {
	var a [reverse.Size]int
	for i := 0; i < reverse.Size; i++ {
		a[i] = i
	}

	reverse.Do(&a)
	for i := 0; i < reverse.Size; i++ {
		if a[i] != (reverse.Size - 1 - i) {
			t.Errorf("want: %v, get: %v", reverse.Size-1-i, a[i])
		}
	}
}
