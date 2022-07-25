package phonenumberlettercombination

import "fmt"

var (
	symbols = map[string]string{
		"0": "",
		"1": "",
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	result = make([]string, 0)
)

func LetterCombinations(digits string) []string {
	result = make([]string, 0)
	if len(digits) > 0 {
		char_list := make([]string, 0)
		for _, v := range digits {
			char_list = append(char_list, symbols[string(v)])
		}
		fmt.Println(char_list)
		combinations(char_list, 0, "")
	}
	return result
}

func combinations(c []string, current int, temp string) {
	if current == len(c) {
		result = append(result, temp)
		return
	}
	for _, v := range c[current] {
		temp = temp + string(v)
		combinations(c, current+1, temp)
		temp = temp[0 : len(temp)-1]
	}
}
