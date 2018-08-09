package num27

import "fmt"

//自解
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] == val {
			nums = []int{}
			return 0
		} else {
			return 1
		}
	}
	p := 0
	for i := 0; i < len(nums); {
		if nums[i] != val {
			i++
		} else {
			p = i
			for ; i < len(nums); i++ {
				if nums[i] != val {
					break
				}
			}
			if i == len(nums) {
				break
			}
			a := nums[p]
			nums[p] = nums[i]
			nums[i] = a
			i = p + 1
		}
	}
	if nums[p] == val && p == 0 {
		nums = []int{}
		return 0
	}
	if p != 0 {
		nums = nums[:p]
	}
	fmt.Println(nums)
	return len(nums)
}

//解法一
func removeElement1(nums []int, val int) int {
	index := -1
	for i := 0; i < len(nums); i++ {
		if val != nums[i] {
			index = index + 1
			nums[index] = nums[i]
		}
	}
	nums = nums[:index+1]
	return len(nums)
}

//解法二
func removeElement2(nums []int, val int) int {
	n := len(nums)
	i := 0
	for i < n {
		if val == nums[i] {
			nums[i] = nums[n-1]
			n-- //发生一次复制数组长度-1
		} else {
			i++
		}
	}
	nums = nums[:n]
	return len(nums)
}
