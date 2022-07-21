package zigzagconversion_test

import (
	"godsa/problems/strings/zigzagconversion"
	"testing"
)

func TestConversion(t *testing.T) {
	tests := []struct {
		name  string
		input string
		row   int
		want  string
	}{
		{
			name:  "tc1",
			input: "PAYPALISHIRING",
			row:   3,
			want:  "PAHNAPLSIIGYIR",
		},
		{
			name:  "tc2",
			input: "PAYPALISHIRING",
			row:   4,
			want:  "PINALSIGYAHRPI",
		},
		{
			name:  "tc3",
			input: "PAYPALISHIRING",
			row:   5,
			want:  "PHASIYIRPLIGAN",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := zigzagconversion.Convert(tc.input, tc.row); got != tc.want {
				t.Errorf("\ngot: %v, \nwant: %v", got, tc.want)
			}
		})
	}
}
