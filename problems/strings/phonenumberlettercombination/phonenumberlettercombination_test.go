package phonenumberlettercombination_test

import (
	"godsa/problems/strings/phonenumberlettercombination"
	"reflect"
	"testing"
	"time"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output []string
	}{
		{
			name:   "tc1",
			input:  "23",
			output: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:   "tc2",
			input:  "",
			output: []string{},
		},
		{
			name:   "tc3",
			input:  "2",
			output: []string{"a", "b", "c"},
		},
		{
			name:   "tc4",
			input:  "567",
			output: []string{"jmp", "jmq", "jmr", "jms", "jnp", "jnq", "jnr", "jns", "jop", "joq", "jor", "jos", "kmp", "kmq", "kmr", "kms", "knp", "knq", "knr", "kns", "kop", "koq", "kor", "kos", "lmp", "lmq", "lmr", "lms", "lnp", "lnq", "lnr", "lns", "lop", "loq", "lor", "los"},
		},
		{
			name:  "tc5",
			input: "269",
			output: []string{
				"amw", "amx", "amy", "amz", "anw", "anx", "any", "anz", "aow", "aox", "aoy", "aoz", "bmw", "bmx", "bmy", "bmz", "bnw", "bnx", "bny", "bnz", "bow", "box", "boy", "boz", "cmw", "cmx", "cmy", "cmz", "cnw", "cnx", "cny", "cnz", "cow", "cox", "coy", "coz",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			now := time.Now()
			if got := phonenumberlettercombination.LetterCombinations(tc.input); !reflect.DeepEqual(got, tc.output) {
				t.Errorf("\ngot: %v, \nwant: %v", got, tc.output)
			}
			t.Log("Elapsed: ", time.Since(now))
		})
	}
}
