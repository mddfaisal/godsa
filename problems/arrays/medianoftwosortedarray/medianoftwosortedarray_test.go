package medianoftwosortedarray_test

import (
	"godsa/problems/arrays/medianoftwosortedarray"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name   string
		i1     []int
		i2     []int
		result int
	}{
		{
			name:   "name1",
			i1:     []int{1, 3},
			i2:     []int{2},
			result: 2,
		},
		{
			name:   "name2",
			i1:     []int{0, 0, 0, 0, 0},
			i2:     []int{-1, 0, 0, 0, 0, 0, 1},
			result: 0,
		},
		{
			name:   "name3",
			i1:     []int{},
			i2:     []int{1},
			result: 1,
		},
		{
			name:   "name4",
			i1:     []int{},
			i2:     []int{1, 2, 3, 4, 5},
			result: 3,
		},
	}

	for _, tcase := range tests {
		t.Run(tcase.name, func(t *testing.T) {
			if got := medianoftwosortedarray.FindMedianSortedArrays(tcase.i1, tcase.i2); got != float64(tcase.result) {
				t.Errorf("want: %v, found: %v", tcase.result, got)
			}
		})
	}
}
