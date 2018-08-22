package num58

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}
	strTail := len(s) - 1
	if string(s[len(s)-1]) == " " {
		for i := strTail; i >= 0; i-- {
			if string(s[i]) != " " {
				strTail = i
				break
			}
		}
	}

	for i := strTail; i >= 0; i-- {
		if string(s[i]) == " " {
			return strTail - i
		}
	}
	return strTail + 1
}

func lengthOfLastWord2(s string) int {
	if s == "" {
		return 0
	}
	strTail := len(s) - 1
	if string(s[len(s)-1]) == " " {
		for i := strTail; i >= 0; i-- {
			if string(s[i]) != " " {
				s = s[:i+1]
				break
			}
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			return len(s) - 1 - i
		}
	}
	return len(s)
}
