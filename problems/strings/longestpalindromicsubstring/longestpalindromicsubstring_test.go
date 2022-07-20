package longestpalindromicsubstring_test

import (
	"godsa/problems/strings/longestpalindromicsubstring"
	"testing"
	"time"
)

func TestLongestPalindrom(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "tc1",
			input: "babad",
			want:  "bab",
		},
		{
			name:  "tc2",
			input: "cbbd",
			want:  "bb",
		},
		{
			name:  "tc3",
			input: "ac",
			want:  "a",
		},
		{
			name:  "tc4",
			input: "bb",
			want:  "bb",
		},
		{
			name:  "tc5",
			input: "jglknendplocymmvwtoxvebkekzfdhykknufqdkntnqvgfbahsljkobhbxkvyictzkqjqydczuxjkgecdyhixdttxfqmgksrkyvopwprsgoszftuhawflzjyuyrujrxluhzjvbflxgcovilthvuihzttzithnsqbdxtafxrfrblulsakrahulwthhbjcslceewxfxtavljpimaqqlcbrdgtgjryjytgxljxtravwdlnrrauxplempnbfeusgtqzjtzshwieutxdytlrrqvyemlyzolhbkzhyfyttevqnfvmpqjngcnazmaagwihxrhmcibyfkccyrqwnzlzqeuenhwlzhbxqxerfifzncimwqsfatudjihtumrtjtggzleovihifxufvwqeimbxvzlxwcsknksogsbwwdlwulnetdysvsfkonggeedtshxqkgbhoscjgpiel",
			want:  "sknks",
		},
		{
			name:  "tc6",
			input: "qartysknksbhhok",
			want:  "sknks",
		},
		{
			name:  "tc7",
			input: "aacabdkacaa",
			want:  "aca",
		},
	}

	for _, tc := range tests {
		// if tc.name != "tc7" {
		// 	continue
		// }
		t.Run(tc.name, func(t *testing.T) {
			now := time.Now()
			if got := longestpalindromicsubstring.LongestPalindrom(tc.input); got != tc.want {
				t.Errorf("\ngot: %v \n want: %v", got, tc.want)
			}
			t.Logf("Tcase %v time: %v", tc.name, time.Since(now))
		})
	}
}
