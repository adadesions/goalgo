package tree

import "errors"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type Stack struct {
	Data []Node
}

func (s *Stack) Pop() (Node, error) {
	if len(s.Data) == 0 {
		return Node{}, errors.New("Stack is empty.")
	}

	last := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return last, nil
}

func (s *Stack) Push(newData *Node) {
	s.Data = append(s.Data, *newData)
}

type Quene struct {
	Data []int
}

func (q *Quene) Enquene() (int, error) {
	if len(q.Data) == 0 {
		return -1, errors.New("Quene is empty.")
	}

	first := q.Data[0]
	q.Data = q.Data[1:]

	return first, nil
}
