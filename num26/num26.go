package num26

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	p := 0
	for i := 1; i < len(nums); {
		if nums[i] == nums[p] {
			i++
		} else {
			p = p + 1
			nums[p] = nums[i]
			i++
		}
	}
	nums = nums[:p+1]
	fmt.Println(nums)
	return len(nums)
}
