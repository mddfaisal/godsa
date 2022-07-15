package binarysearch

const Path = "../../input/long.arr"

func Search(arr []int, key int) bool {
	var (
		found = false
		left  = 0
		right = len(arr) - 1
	)
	for left <= right {
		mid := (left + (right - 1)) / 2
		if arr[mid] == key {
			found = true
			break
		}

		if key > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return found
}
