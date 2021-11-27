package rotate_test

import (
	"gobook/ch4/ex04/rotate"
	"testing"
)

func TestRotate(t *testing.T) {
	var chk int = 8
	a := make([]int, 10)
	for i := 0; i < 10; i++ {
		a[i] = i
	}

	anew := rotate.Do(a, chk)
	for i := 0; i < len(a)-chk; i++ {
		if anew[i] != a[i+chk] {
			t.Errorf("want: %v, get: %v", a[i+chk], anew[i])
		}
	}
	for i := 0; i < chk; i++ {
		if anew[i+len(a)-chk] != a[i] {
			t.Errorf("want: %v, get: %v", a[i], anew[i+len(a)-chk])
		}
	}

}
