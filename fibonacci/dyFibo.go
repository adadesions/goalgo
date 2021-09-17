package fibonacci

func DyFibo(n int) int {
	f := []int{1, 1}

	for i := 2; i <= n; i++ {
		f = append(f, f[i-1]+f[i-2])
	}

	return f[n]
}
