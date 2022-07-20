package longestpalindromicsubstring

func LongestPalindrom(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	longestSubstr, strBytes, size, longestSubstrSize := "", []byte(s), len(s), 0
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			// aacabdkacaa
			start, end := i, j+1
			current := strBytes[start:end]
			if len(current) > 0 {
				if current[0] == current[len(current)-1] {
					if ispalindrom(current) {
						if longestSubstrSize < len(current) {
							longestSubstrSize = len(current)
							longestSubstr = string(current)
						}
					}
				}
			}
		}
	}
	return longestSubstr
}

func ispalindrom(b []byte) bool {
	temp := make([]byte, 0)
	temp = append(temp, b...)
	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}
	return string(temp) == string(b)
}
