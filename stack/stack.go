package stack

type Stack struct {
	Data []int
}

func (s *Stack) Pop() int {
	lastIdx := len(s.Data) - 1
	ele := s.Data[lastIdx]
	s.Data = s.Data[:lastIdx]

	return ele
}

func (s *Stack) Push(data int) []int {
	s.Data = append(s.Data, data)

	return s.Data
}

func (s *Stack) Seek() int {
	return s.Data[len(s.Data)-1]
}
