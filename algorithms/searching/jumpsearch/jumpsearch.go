package jumpsearch

import "math"

const Path = "../../input/long.arr"

func JumpSearch(arr []int, key int) bool {
	var (
		found = false
		size  = len(arr)
		m     = int(math.Sqrt(float64(size)))
		i     = 0
	)
	for arr[m] <= key && m < size {
		i = m
		m = m + int(math.Sqrt(float64(size)))
		if m > size-1 {
			break
		}
	}

	for j := i; j < m; j++ {
		if arr[j] == key {
			found = true
			break
		}
	}

	return found
}
