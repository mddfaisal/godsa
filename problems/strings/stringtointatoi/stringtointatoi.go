package stringtointatoi

import (
	"math"
	"strconv"
	"strings"
)

func MyAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	sign := 1
	strbytes := []byte(s)
	num := 0
	if strbytes[0] == 45 || strbytes[0] == 43 {
		if strbytes[0] == 45 {
			sign = -1
		}
		strbytes = strbytes[1:]
	}
	for _, b := range strbytes {
		if b >= 48 && b <= 57 {
			i, _ := strconv.Atoi(string(b))
			num = num*10 + i
			if num >= math.MaxInt32 {
				if sign == -1 {
					if num*sign <= math.MinInt32 {
						return math.MinInt32
					}
				} else {
					return math.MaxInt32
				}
			}
		} else {
			break
		}
	}
	num = num * sign
	return num
}
