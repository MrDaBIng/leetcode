package num20

import (
	"fmt"
	"testing"
)

func Test_reverse(t *testing.T) {
	fmt.Println(isValid("({})"))
}

func Test_Stack(t *testing.T) {
	s:=NewStack()
	s.Push("a")
	s.Push("b")
	s.Push("c")
	fmt.Println(s.Pop())
	fmt.Println(s.Top())
}
