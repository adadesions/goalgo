package sorting

func SelectSort(data []int) []int {
	var swap func(a, b *int)
	swap = func(a, b *int) {
		temp := *a
		*a = *b
		*b = temp
	}

	for i := 0; i < len(data); i++ {
		minIdx := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[minIdx] {
				minIdx = j
			}
		}
		swap(&data[i], &data[minIdx])
	}
	return data
}
