package romantoint_test

import (
	"godsa/problems/basic/romantoint"
	"testing"
	"time"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name   string
		roman  string
		result int
	}{
		{
			name:   "tc1",
			roman:  "III",
			result: 3,
		},
		{
			name:   "tc2",
			roman:  "LVIII",
			result: 58,
		},
		{
			name:   "tc3",
			roman:  "MCMXCIV",
			result: 1994,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			now := time.Now()
			if got := romantoint.RomanToInt(tc.roman); got != tc.result {
				t.Errorf("\ngot: %v, \nwant: %v", got, tc.result)
			}
			t.Logf("%v time: %v", tc.name, time.Since(now))
		})
	}
}
