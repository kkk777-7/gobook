package duprmv_test

import (
	"gobook/ch4/ex05/duprmv"
	"reflect"
	"testing"
)

func TestDuprmv(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  []string
	}{
		{name: "No Duplicate", words: []string{"1", "2", "3", "4", "5"}, want: []string{"1", "2", "3", "4", "5"}},
		{name: "One Duplicate", words: []string{"1", "2", "3", "4", "4", "5"}, want: []string{"1", "2", "3", "4", "5"}},
		{name: "Two Duplicate", words: []string{"1", "2", "3", "4", "4", "4", "5"}, want: []string{"1", "2", "3", "4", "5"}},
		{name: "Three Duplicate", words: []string{"1", "2", "2", "3", "4", "4", "4", "5"}, want: []string{"1", "2", "3", "4", "5"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := duprmv.Do(tt.words)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("want: %v, get: %v", tt.want, got)
			}
		})
	}

}
