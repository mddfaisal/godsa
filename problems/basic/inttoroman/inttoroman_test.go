package inttoroman_test

import (
	"godsa/problems/basic/inttoroman"
	"testing"
	"time"
)

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		name    string
		integer int
		result  string
	}{
		{
			name:    "tc1",
			integer: 3,
			result:  "III",
		},
		{
			name:    "tc2",
			integer: 58,
			result:  "LVIII",
		},
		{
			name:    "tc3",
			integer: 1994,
			result:  "MCMXCIV",
		},
		{
			name:    "tc4",
			integer: 3999,
			result:  "MMMCMXCIX",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			now := time.Now()
			if got := inttoroman.IntToRoman(tc.integer); got != tc.result {
				t.Errorf("\ngot: %v, \nwant: %v", got, tc.result)
			}
			t.Logf("%v time: %v", tc.name, time.Since(now))
		})
	}
}
