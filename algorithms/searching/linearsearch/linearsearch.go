package linearsearch

const Path = "../../input/long.arr"

func Search(arr []int, key int) bool {
	found := false
	for i := 0; i < len(arr); i++ {
		if key == arr[i] {
			found = true
			break
		}
	}
	return found
}
