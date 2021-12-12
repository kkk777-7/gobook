package treesort

import "fmt"

type Tree struct {
	value       int
	left, right *Tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *Tree
	for _, v := range values {
		root = Add(root, v)
	}
	AppendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func AppendValues(values []int, t *Tree) []int {
	if t != nil {
		values = AppendValues(values, t.left)
		values = append(values, t.value)
		values = AppendValues(values, t.right)
	}
	return values
}

func Add(t *Tree, value int) *Tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(Tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = Add(t.left, value)
	} else {
		t.right = Add(t.right, value)
	}
	return t
}

func (t *Tree) String() {
	if t.left != nil {
		t.left.String()
	}
	fmt.Println(t.value)
	if t.right != nil {
		t.right.String()
	}
}
