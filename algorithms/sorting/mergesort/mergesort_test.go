package mergesort_test

import (
	"godsa/algorithms/sorting/mergesort"
	"reflect"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "tc1",
			input: []int{1, 5, 3, 8, 6, 2, 9, 7, 4},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			now := time.Now()
			if got := mergesort.Sort(tc.input); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\ngot: %v, \nwant: %v", got, tc.want)
			}
			t.Log("Elapsed: ", time.Since(now))
		})
	}
}
