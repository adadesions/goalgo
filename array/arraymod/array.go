package arraymod

type methods interface {
	NewArrayList() *ArrayList
	Add(newEle int)
	Remove(target int)
	Update(target, newValue int)
}

type ArrayList struct {
	data []int
	cap  int
	size int
}

func NewArrayList() *ArrayList {
	return &ArrayList{}
}

func Add(data []int, newEle int) []int {
	newObj := []int{}
	copy(newObj, data)
	newObj = append(data, newEle)
	return newObj
}
