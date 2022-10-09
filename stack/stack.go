package stack

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) []int {
	s.items = append(s.items, item)
	return s.items
}

func (s *Stack) Pop() int {
	i := len(s.items) - 1
	item := s.items[i]
	s.items = s.items[:i]
	return item
}
