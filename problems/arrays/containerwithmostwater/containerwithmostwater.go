package containerwithmostwater

func MaxArea(height []int) int {
	area, low, high := 0, 0, len(height)-1
	for low < high {
		h := map[bool]int{true: height[high], false: height[low]}[height[low] > height[high]]
		distance := high - low
		water := h * distance
		if water < 0 {
			water = -water
		}
		if area < water {
			area = water
		}
		if height[low] < height[high] {
			low++
		} else {
			high--
		}
	}
	return area
}
