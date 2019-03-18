package num70

import "fmt"

func climbStairs(n int) int {
	step := map[int]int{
		1: 1,
		2: 2,
	}

	for i := 3; i <= n; i++ {
		step[i] = step[i-1] + step[i-2]
	}

	return step[n]
}

var resultMap = map[int]int{
	0: 1,
	1: 1,
	2: 2,
}

func climbStairs2(n int) int {
	if _, ok := resultMap[n]; ok {
		return resultMap[n]
	}

	ret := climbStairs2(n-1) + climbStairs2(n-2)
	resultMap[n] = ret
	fmt.Println(resultMap)
	return ret
}
