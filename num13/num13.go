package num13

var ROMAN map[string]int = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	if len(s) <= 1 {
		return ROMAN[getStrByindex(s, 0)]
	}
	sum := 0
	for i := 0; i < len(s); {
		if getStrByindex(s, i) == "I" && i+1 < len(s) {
			j := i + 1
			if getStrByindex(s, j) == "V" || getStrByindex(s, j) == "X" {
				sum = sum + (ROMAN[getStrByindex(s, j)] - ROMAN["I"])
				i = j + 1
				continue
			}

		}
		if getStrByindex(s, i) == "X" && i+1 < len(s) {
			j := i + 1
			if getStrByindex(s, j) == "L" || getStrByindex(s, j) == "C" {
				sum = sum + (ROMAN[getStrByindex(s, j)] - ROMAN["X"])
				i = j + 1
				continue
			}
		}
		if getStrByindex(s, i) == "C" && i+1 < len(s) {
			j := i + 1
			if getStrByindex(s, j) == "D" || getStrByindex(s, j) == "M" {
				sum = sum + (ROMAN[getStrByindex(s, j)] - ROMAN["C"])
				i = j + 1
				continue
			}
		}
		sum = sum + ROMAN[getStrByindex(s, i)]
		i = i + 1
	}
	return sum
}

func getStrByindex(str string, index int) string {
	return string(str[index])
}
