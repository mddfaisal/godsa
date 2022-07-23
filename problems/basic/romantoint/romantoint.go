package romantoint

import "strings"

var (
	symbols = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
)

func RomanToInt(s string) int {
	res := 0
	s = strings.ReplaceAll(s, "IV", "IIII")
	s = strings.ReplaceAll(s, "IX", "VIIII")
	s = strings.ReplaceAll(s, "XL", "XXXX")
	s = strings.ReplaceAll(s, "XC", "LXXXX")
	s = strings.ReplaceAll(s, "CD", "CCCC")
	s = strings.ReplaceAll(s, "CM", "DCCCC")
	for _, v := range []byte(s) {
		res = res + symbols[string(v)]
	}
	return res
}
