package stack

type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() {
	s.items = s.items[:len(s.items)-1]
}
