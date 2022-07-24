package foursum_test

import (
	"godsa/problems/arrays/foursum"
	"testing"
	"time"
)

func TestFourSum(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		result [][]int
	}{
		{
			name:   "tc1",
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
			result: [][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			name:   "tc2",
			nums:   []int{2, 2, 2, 2, 2},
			target: 8,
			result: [][]int{
				{2, 2, 2, 2},
			},
		},
		{
			name:   "tc3",
			nums:   []int{-5, 5, 4, -3, 0, 0, 4, -2},
			target: 4,
			result: [][]int{
				{-5, 0, 4, 5},
				{-3, -2, 4, 5},
			},
		},
		{
			name:   "tc4",
			nums:   []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			target: 8,
			result: [][]int{
				{2, 2, 2, 2},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			now := time.Now()
			// if got := foursum.FourSum(tc.nums, tc.target); !reflect.DeepEqual(got, tc.result) {
			// 	t.Errorf("\ngot: %v, \nwant: %v", got, tc.result)
			// }
			foursum.FourSum(tc.nums, tc.target)
			t.Log("Elapsed: ", time.Since(now))
		})
	}
}
