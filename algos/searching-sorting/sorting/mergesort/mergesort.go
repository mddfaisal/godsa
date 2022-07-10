package mergesort

const Path = "../../input/long.arr"

func Sort(arr []int) []int {
	return mergesort(arr)
}

func mergesort(arr []int) []int {
	size := len(arr)
	if size == 1 {
		return arr
	}
	mid := int(size / 2)
	left := make([]int, mid)
	right := make([]int, size-mid)
	for i := 0; i < size; i++ {
		if i < mid {
			left[i] = arr[i]
		} else {
			right[i-mid] = arr[i]
		}
	}
	return merge(mergesort(left), mergesort(right))
}

func merge(left []int, right []int) []int {
	result := make([]int, len(left)+len(right))
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = result[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return result
}
