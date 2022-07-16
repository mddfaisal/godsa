package medianoftwosortedarray

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totlen := len(nums1) + len(nums2)
	arr := []int{}
	arr = append(arr, nums1...)
	arr = append(arr, nums2...)
	arr = mergeSort(arr)
	if totlen%2 == 0 {
		m := totlen / 2
		return checkNum((float64(arr[m]) + float64(arr[m-1])) / 2)
	}
	key := map[bool]int{true: 0, false: int(totlen / 2)}[totlen/2 == 0]
	return checkNum(float64(arr[key]))
}

func checkNum(i float64) float64 {
	if i < 0 && i > -1 {
		return 0
	}
	return i
}

func mergeSort(arr []int) []int {
	size := len(arr)
	if size == 1 {
		return arr
	}
	mid := int(size / 2)
	left := make([]int, mid)
	right := make([]int, size-mid)
	for i, val := range arr {
		if i < mid {
			left[i] = val
		} else {
			right[i-mid] = val
		}
	}
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left []int, right []int) []int {
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
