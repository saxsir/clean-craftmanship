package stack

type Stack struct {
	empty bool
}

func NewStack() *Stack {
	return &Stack{
		empty: true,
	}
}

var s = NewStack()

func (s *Stack) isEmpty() bool {
	return s.empty
}

func (s *Stack) push(_element int) {
	s.empty = false
}

}
