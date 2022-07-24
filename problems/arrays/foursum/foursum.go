package foursum

func FourSum(nums []int, target int) [][]int {
	sum_map, result := map[[4]int]int{}, [][]int{}
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			for k := j; k < len(nums); k++ {
				for l := k; l < len(nums); l++ {
					if i != j && j != k && k != l {
						sum := nums[i] + nums[j] + nums[k] + nums[l]
						if sum == target {
							arr := sort([]int{nums[i], nums[j], nums[k], nums[l]})
							a := [4]int{arr[0], arr[1], arr[2], arr[3]}
							sum_map[a] = sum
						}
					}
				}
			}
		}
	}
	for k, _ := range sum_map {
		a := []int{}
		for _, v := range k {
			a = append(a, v)
		}
		result = append(result, a)
	}
	return result
}

func sort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
