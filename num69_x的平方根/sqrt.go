package sqrt

func mySqrt(x int) int {
	if x == 1 || x == 0 {
		return x
	}
	for i := 0; i < x; i++ {
		if i*i == x {
			return i
		}
		if i*i < x && (i+1)*(i+1) > x {
			return i
		}
	}

	return -1
}
