package threesumclosest

func ThreeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	result := nums[0] + nums[1] + nums[len(nums)-1]
	nums = sort(nums)
	for i := 0; i < len(nums)-2; i++ {
		start, end := i+1, len(nums)-1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			if sum > target {
				end--
			} else {
				start++
			}
			if abs(sum-target) < abs(result-target) {
				result = sum
			}
		}
	}
	return result
}

func abs(i int) int {
	return map[bool]int{true: i, false: -i}[i >= 0]
}

func sort(arr []int) []int {
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
	return merge(sort(left), sort(right))
}

func merge(left, right []int) []int {
	result, i := make([]int, len(left)+len(right)), 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
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
