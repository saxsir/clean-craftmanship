package stack

type Stack struct {
	size int
}

func NewStack() *Stack {
	return &Stack{
		size: 0,
	}
}

var s = NewStack()

func (s *Stack) isEmpty() bool {
	return s.size == 0
}

func (s *Stack) push(_element int) {
	s.size += 1
}

func (s *Stack) pop() int {
	s.size -= 1
	return -1

}

func (s *Stack) getSize() int {
	return s.size
}
