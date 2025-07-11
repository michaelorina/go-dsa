package sort

type StackNode struct {
	Value int
	Next  *StackNode
}

type Stack struct {
	top  *StackNode
	size int
}

func (s *Stack) Push(value int) {
	newNode := &StackNode{Value: value, Next: s.top}
	s.top = newNode
	s.size++
}

func (s *Stack) Pop() (int, bool) {
	if s.top == nil {
		return 0, false
	}
	val := s.top.Value
	s.top = s.top.Next
	s.size--
	return val, true
}

func (s *Stack) Peek() (int, bool) {
	if s.top == nil {
		return 0, false
	}
	return s.top.Value, true
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Len() int {
	return s.size
}
