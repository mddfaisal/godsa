package reverseinteger_test

import (
	"godsa/problems/strings/reverseinteger"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{
			name:  "tc1",
			input: 123,
			want:  321,
		},
		{
			name:  "tc2",
			input: -123,
			want:  -321,
		},
		{
			name:  "tc3",
			input: 1534236469,
			want:  0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := reverseinteger.Reverse(tc.input); got != tc.want {
				t.Errorf("\ngot: %v, \nwant: %v", got, tc.want)
			}
		})
	}
}
