package insertionsort

const Path = "../../input/long.arr"

func Sort(arr []int) {
	size := len(arr)
	for i := 0; i < size-1; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
