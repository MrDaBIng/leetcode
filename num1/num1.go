package num1

func twoSum(nums []int, target int) []int {
	ret := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		ret[0] = i
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ret[1] = j
				return ret
			}
		}
	}
	ret = []int{}
	return ret
}
