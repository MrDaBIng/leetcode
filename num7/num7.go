package num7

import (
	"math"
)

func reverse(x int) int {
	rev := 0
	for x != 0 {
		pop := x % 10
		x = x / 10
		if rev > (math.MaxInt32-pop)/10 {
			return 0
		}
		if rev < (math.MinInt32-pop)/10 {
			return 0
		}
		rev = 10*rev + pop
	}
	return rev
}
