package num53

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum := nums[0]
	var sum int
	for i := 0; i < len(nums); i++ {
		sum = 0
		for j := i; j < len(nums); j++ {
			sum = sum + nums[j]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}
