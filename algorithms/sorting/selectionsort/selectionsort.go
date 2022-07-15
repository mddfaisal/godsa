package selectionsort

func Sort(arr []int) []int {
	size := len(arr)
	for i := 0; i < size-1; i++ {
		minElementIndex := i
		for j := i + 1; j < size; j++ {
			if arr[j] < arr[minElementIndex] {
				minElementIndex = j
			}
		}
		arr[minElementIndex], arr[i] = arr[i], arr[minElementIndex]
	}
	return arr
}
