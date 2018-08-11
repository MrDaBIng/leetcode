package num28

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	if len(needle) == 0 {
		return 0
	}
	i := 0
	for i < len(haystack) {
		if needle[0] == haystack[i] {
			d := i
			j := 0
			for ; j < len(needle); j++ {
				if d > len(haystack)-1 {
					return -1
				}
				if needle[j] != haystack[d] {
					break
				}
				d++
			}
			if j == len(needle) {
				return i
			}
		}
		i++
	}
	return -1
}

//灵光一闪，推荐用这种！！！！
func strStr2(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i < len(haystack); i++ {
		if i+len(needle) > len(haystack) {
			return -1
		}
		if needle[0] == haystack[i] && needle == haystack[i:i+len(needle)] {
			return i
		}
	}
	return -1
}
