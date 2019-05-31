package chapter

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
	}{
		{
			name: "one",
			a:    []int{1},
			b:    []int{1},
		},

		{
			name: "two",
			a:    []int{2, 1},
			b:    []int{1, 2},
		},

		{
			name: "same",
			a:    []int{2, 2, 2},
			b:    []int{2, 2, 2},
		},

		{
			name: "many",
			a:    []int{5, 2, 3, 3, 1},
			b:    []int{1, 2, 3, 3, 5},
		},

		{
			name: "exercise 1.1",
			a:    []int{31, 41, 59, 26, 41, 58},
			b:    []int{26, 31, 41, 41, 58, 59},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertionSort(tt.a)
			if !reflect.DeepEqual(tt.a, tt.b) {
				t.Errorf("wrong value %v, want %v", tt.a, tt.b)
			}
		})
	}
}
