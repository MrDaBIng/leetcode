package num14

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	prefix := ""
	for i := 0; ; i++ {
		s := ""
		for j := 0; j < len(strs)-1; j++ {
			if i >= len(strs[j]) {
				return prefix
			}
			if i >= len(strs[j+1]) {
				return prefix
			}

			if strs[j][i] != strs[j+1][i] {
				return prefix
			}
			s = string(strs[j][i])
		}

		prefix = prefix + s
	}
	return prefix
}
