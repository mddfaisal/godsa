package threesum

func ThreeSum(nums []int) [][]int {
	nums = mergeSort(nums)
	arrMap := make(map[[3]int][]int)
	for i := 0; i < len(nums); i++ {
		low, high := i+1, len(nums)-1
		for low < high {
			if nums[i]+nums[low]+nums[high] > 0 {
				high--
			} else if nums[i]+nums[low]+nums[high] < 0 {
				low++
			} else {
				arrMap[sort([3]int{nums[i], nums[low], nums[high]})] = []int{nums[i], nums[low], nums[high]}
				low++
				high--
			}
		}
	}
	arr := [][]int{}
	for _, v := range arrMap {
		arr = append(arr, v)
	}
	return arr
}

func mergeSort(arr []int) []int {
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
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i := 0
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

func sort(arr [3]int) [3]int {
	a, b, c := arr[0], arr[1], arr[2]
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
		if a > b {
			a, b = b, a
		}
	}
	return [3]int{a, b, c}
}
