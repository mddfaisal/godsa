package foursum

type pair struct {
	num1 int
	num2 int
}

func (p *pair) sum() int {
	return p.num1 + p.num2
}

func FourSum(nums []int, target int) [][]int {
	result := [][]int{}
	pairs := []pair{}
	pair_map := map[[4]int]int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			pairs = append(pairs, pair{num1: nums[i], num2: nums[j]})
		}
	}

	for i := 0; i < len(pairs); i++ {
		for j := i + 1; j < len(pairs); j++ {
			if pairs[i].sum()+pairs[j].sum() == target {
				a := sort([]int{pairs[i].num1, pairs[i].num2, pairs[j].num1, pairs[j].num2})
				pair_map[[4]int{a[0], a[1], a[2], a[3]}] = target
			}
		}
	}

	for k, _ := range pair_map {
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
