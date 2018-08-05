package num20

import (
	"container/list"
	"fmt"
)

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	if len(s) == 0 {
		return true
	}
	stack := NewStack()
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "[" {
			stack.Push("]")
		} else if string(s[i]) == "{" {
			stack.Push("}")
		} else if string(s[i]) == "(" {
			stack.Push(")")
		} else {
			if stack.IsEmpty() || string(s[i]) != fmt.Sprint(stack.Top()) {
				return false
			}
			stack.Pop()
		}
	}
	return stack.IsEmpty()
}

//golang stack的简单实现
type Stack struct {
	*list.List
}

func NewStack() Stack {
	return Stack{
		List: list.New(),
	}
}

func (s Stack) Push(v interface{}) {
	s.List.PushBack(v)
}

func (s Stack) Pop() interface{} {
	e := s.Back()
	s.Remove(e)
	return e.Value
}

func (s Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s Stack) Top() interface{} {
	return s.Back().Value
}
