package array_rotation_test

import (
	"testing"
)

func TestMethod1(t *testing.T) {
	var tests = []struct {
		name           string
		arr            []int
		size, elements int
	}{
		{"1", []int{1, 2, 3, 4, 5}, 5, 2},
		{"1", []int{2, 3, 4, 5, 6}, 5, 2},
		{"1", []int{4, 5, 6, 7, 8}, 5, 2},
		{"1", []int{9, 1, 2, 3, 4}, 5, 2},
		{"1", []int{6, 4, 7, 8, 3, 2}, 5, 2},
	}
	for _, tcase := range tests {
		t.Run(tcase.name, func(t *testing.T) {

		})
	}
}
