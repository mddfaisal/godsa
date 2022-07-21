package stringtointatoi_test

import (
	"godsa/problems/strings/stringtointatoi"
	"testing"
	"time"
)

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "tc1",
			input: "42",
			want:  42,
		},
		{
			name:  "tc2",
			input: "-42",
			want:  -42,
		},
		{
			name:  "tc3",
			input: "4193 with words",
			want:  4193,
		},
		{
			name:  "tc4",
			input: "words and 987",
			want:  0,
		},
		{
			name:  "tc5",
			input: "-91283472332",
			want:  -2147483648,
		},
		{
			name:  "tc6",
			input: "20000000000000000000",
			want:  2147483647,
		},
		{
			name:  "tc7",
			input: "9223372036854775808",
			want:  2147483647,
		},
		{
			name:  "tc8",
			input: "-2147483647",
			want:  -2147483647,
		},
	}

	for _, tc := range tests {
		// if tc.name != "tc7" {
		// 	continue
		// }
		t.Run(tc.name, func(t *testing.T) {
			now := time.Now()
			if got := stringtointatoi.MyAtoi(tc.input); got != tc.want {
				t.Errorf("\ngot: %v, \nwant: %v", got, tc.want)
			}
			t.Logf("%v time: %v", tc.name, time.Since(now))
		})
	}
}
