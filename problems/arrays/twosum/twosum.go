package twosum

func TwoSum(n []int, target int) []int {
	for i := 0; i < len(n)-1; i++ {
		for j := i + 1; j < len(n); j++ {
			if n[i]+n[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
