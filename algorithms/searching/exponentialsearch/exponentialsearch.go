package exponentialsearch

import "godsa/algorithms/searching/binarysearch"

func Search(arr []int, key int) bool {
	found := false
	size := len(arr)
	i := 1
	for i < size && arr[i] <= key {
		i = i * 2
	}
	found = binarysearch.Search(arr[int(i/2):size-1], key)
	return found
}
