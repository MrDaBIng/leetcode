package num66

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i] + 1
		if i == 0 && digits[i] > 9 {
			digits = append(digits, 0)
			digits[0] = 1
		}
		if digits[i] <= 9 {
			return digits
		}
		digits[i] = 0
	}
	return digits
}
