package longestcommonprefix_test

import (
	"godsa/problems/arrays/longestcommonprefix"
	"testing"
	"time"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name  string
		array []string
		want  string
	}{
		{
			name:  "tc1",
			array: []string{"flower", "flow", "flight"},
			want:  "fl",
		},
		{
			name:  "tc2",
			array: []string{"dog", "racecar", "car"},
			want:  "",
		},
		{
			name:  "tc3",
			array: []string{"a"},
			want:  "a",
		},
		{
			name:  "tc4",
			array: []string{"flower", "flower", "flower", "flower"},
			want:  "flower",
		},
		{
			name:  "tc5",
			array: []string{"flower", "flower", "flower", "flo"},
			want:  "flo",
		},
	}

	for _, tc := range tests {
		// if tc.name != "tc4" {
		// 	continue
		// }
		t.Run(tc.name, func(t *testing.T) {
			now := time.Now()
			if got := longestcommonprefix.LongestCommonPrefix(tc.array); got != tc.want {
				t.Errorf("\ngot: %v, \nwant: %v", got, tc.want)
			}
			t.Logf("%v time: %v", tc.name, time.Since(now))
		})
	}
}
