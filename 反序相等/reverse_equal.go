package reverse_equal

import (
	"fmt"
)

func reverseEqual() {
	res := []int{}
	for i := 9999; i/9 >= 1000; i-- {
		if i%9 == 0 {
			d := i / 9
			rev := reverse(i / 9)
			if rev == i {
				res = append(res, d)
			}
		}
	}

	fmt.Println(res)
}

func reverse(x int) int {
	rev := 0
	for x != 0 {
		pop := x % 10
		x = x / 10
		rev = 10*rev + pop
	}
	return rev
}
