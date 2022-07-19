package longestsubstringwithoutrepeatingchars_test

import (
	"godsa/problems/strings/longestsubstringwithoutrepeatingchars"
	"testing"
	"time"
)

var (
	tests = []struct {
		name   string
		input  string
		output int
	}{
		{
			name:   "tc1",
			input:  "abcabcbb",
			output: 3,
		},
		{
			name:   "tc2",
			input:  "bbbbb",
			output: 1,
		},
		{
			name:   "tc3",
			input:  "pwwkew",
			output: 3,
		},
		{
			name:   "tc4",
			input:  "dfcciphajkojvsunbzsezyqiblvquvjxbobjdjjovzyrruettyzswraxexqyszyvnzgsirjeqjxkdbfwzeqyxqxcpnchpafccl",
			output: 3,
		},
	}
)

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			now := time.Now()
			if got := longestsubstringwithoutrepeatingchars.LengthOfLongestSubstring(tc.input); got != tc.output {
				t.Errorf("\ngot: %v \n want: %v", got, tc.output)
			}
			t.Logf("Tcase %v time: %v", tc.name, time.Since(now))
		})
	}
}
