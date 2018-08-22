package num67

import "fmt"

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	res := ""
	carry := 0
	var s int
	for i >= 0 || j >= 0 {
		if i < 0 {
			s, carry = sumDigit(toNum(b[j]), 0, carry)
			res = fmt.Sprintf("%d"+res, s)
			j--
			continue
		}
		if j < 0 {
			s, carry = sumDigit(toNum(a[i]), 0, carry)
			res = fmt.Sprintf("%d"+res, s)
			i--
			continue
		}
		s, carry = sumDigit(toNum(a[i]), toNum(b[j]), carry)
		res = fmt.Sprintf("%d"+res, s)
		i--
		j--
	}
	if carry == 1 {
		res = fmt.Sprintf("%d"+res, 1)
	}
	return res
}

func toNum(b byte) int {
	if string(b) == "0" {
		return 0
	}
	return 1
}

func sumDigit(a, b, carry int) (s, c int) {
	s = a + b + carry
	if s >= 2 {
		s = s - 2
		c = 1
	} else {
		c = 0
	}
	return
}
