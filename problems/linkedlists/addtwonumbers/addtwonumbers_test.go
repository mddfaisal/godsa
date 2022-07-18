package addtwonumbers_test

import (
	"fmt"
	"godsa/problems/linkedlists/addtwonumbers"
	"testing"
	"time"
)

var (
	tests = []struct {
		name string
		i    []int
		k    []int
		want []int
	}{
		{
			name: "name1",
			i:    []int{9, 9, 9, 9, 9, 9, 9},
			k:    []int{9, 9, 9, 9},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
		{
			name: "name2",
			i:    []int{2, 4, 3},
			k:    []int{5, 6, 4},
			want: []int{7, 0, 8},
		},
	}
)

func TestAddTwoNumbers(t *testing.T) {
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			now := time.Now()
			result := addtwonumbers.AddTwoNumbers(addtwonumbers.CreateList(tc.i), addtwonumbers.CreateList(tc.k))
			t.Logf("tcase %v time: %v", tc.name, time.Since(now))
			fmt.Println("")
			addtwonumbers.Display(result)
			fmt.Println("")
		})
	}
}
