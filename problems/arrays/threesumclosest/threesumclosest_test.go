package threesumclosest_test

import (
	"godsa/problems/arrays/threesumclosest"
	"reflect"
	"testing"
	"time"
)

func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		target int
		want   [][]int
	}{
		{
			name:   "name1",
			input:  []int{-1, 0, 1, 2, -1, -4},
			target: 1,
			want:   [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
	}

	for _, tcase := range tests {
		t.Run(tcase.name, func(t *testing.T) {
			now := time.Now()
			if got := threesumclosest.ThreeSumClosest(tcase.input, tcase.target); !reflect.DeepEqual(got, tcase.want) {
				// if tcase.name == "name5" {
				// 	t.Errorf("\ngot: %v \nwant: %v", got, tcase.want)
				// }
				t.Errorf("Failed")
			}
			t.Log(tcase.name, " time: ", time.Since(now))
		})
	}
}
