package num1

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	ret := twoSum(nums, 9)
	fmt.Println(ret)
}
