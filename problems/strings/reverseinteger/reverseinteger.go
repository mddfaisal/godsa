package reverseinteger

import "math"

func Reverse(x int) int {
	var (
		rev  = 0
		sign = map[bool]int{true: -1, false: 1}[x < 0]
		num  = map[bool]int{true: -x, false: x}[x < 0]
		ans  int
	)
	for num > 0 {
		rev = rev*10 + num%10
		num = num / 10
	}
	ans = rev * sign
	return map[bool]int{true: ans, false: 0}[ans > math.MinInt32 && ans < math.MaxInt32]
}
