package longestsubstringwithoutrepeatingchars

import "strings"

func LengthOfLongestSubstring(s string) int {
	temp, result , strBytes := "", -1, []byte(s)
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1{
		return 1
	}
	for _, b := range strBytes {
		current := string(b)
		if strings.Contains(temp, current) {
			arr := strings.Split(temp, current)
			temp = arr[len(arr)-1]
		}
		temp = temp + current
		if result < len(temp) {
			result = len(temp)
		}
	}
	return result
}