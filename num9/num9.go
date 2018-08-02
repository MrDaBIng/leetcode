package num9

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	res := 0
	n := x
	for n != 0 {
		pop := n % 10
		n = n / 10
		res = 10*res + pop
	}
	return x == res
}
