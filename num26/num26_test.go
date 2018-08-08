package num26

import (
	"fmt"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	input1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	input2 := []int{0, 1, 2, 3, 4, 5}
	type in []int
	inputs := []in{input1, input2}
	for _, in := range inputs {
		fmt.Println("input: ", in)
		fmt.Println(removeDuplicates(in))
	}
}
