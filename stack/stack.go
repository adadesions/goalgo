package stack

type Stack struct {
	Data []int
}

func (s *Stack) pop() int {
	lastIdx := len(s.Data) - 1
	ele := s.Data[lastIdx]
	s.Data = s.Data[:lastIdx]

	return ele
}

func (s *Stack) push(data int) []int {
	s.Data = append(s.Data, data)

	return s.Data
}

func (s *Stack) seek() int {
	return s.Data[len(s.Data)-1]
}
