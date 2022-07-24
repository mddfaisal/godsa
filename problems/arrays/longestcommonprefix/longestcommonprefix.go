package longestcommonprefix

func LongestCommonPrefix(strs []string) string {
	prefix, exitloop := "", false
	if len(strs) == 0 {
		return prefix
	}
	if len(strs) == 1 {
		return strs[0]
	}
	for i := 0; i < 200; i++ {
		if len(strs[0]) > 0 {
			pref := strs[0][0]
			for k, val := range strs {
				if len(val) > 0 && pref == val[0] {
					strs[k] = strs[k][1:]
				} else {
					exitloop = true
				}
			}
			if exitloop {
				break
			}
			prefix = prefix + string(pref)
		} else {
			break
		}
	}
	return prefix
}
