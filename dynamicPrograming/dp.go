package dp

func Fibo(n int) int {
	// Tabulation Technique
	f := []int{1, 1}

	for i := 2; i <= n; i++ {
		f = append(f, f[i-1]+f[i-2])
	}

	return f[n]
}

func Fibo2(n int) int {
	// Memoization Technique
	f := [100]int{}

	if n == 0 || n == 1 {
		return 1
	}

	if f[n] != 0 {
		return f[n]
	}

	f[n] = Fibo2(n-1) + Fibo2(n-2)

	return f[n]
}

func fac(n int) int {
	// Tabulation Technique
	f := []int{1, 1}

	for i := 2; i <= n; i++ {
		f = append(f, f[i-1]*i)
	}

	return f[n]
}

func fac2(n int) int {
	// Memoizaiton Technique
	mem := [1000]int{}
	if n == 0 {
		return 1
	}

	if mem[n] != 0 {
		return mem[n]
	}

	mem[n] = n * fac2(n-1)
	return mem[n]
}

func Min(nums ...int) int {
	min := nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
	}

	return min
}

func UglyNumber(n int) int {
	u := [1000]int{}

	u[0] = 1

	p2, p3, p5 := 0, 0, 0

	next2 := 2
	next3 := 3
	next5 := 5

	for i := 1; i < n; i++ {
		u[i] = Min(next2, next3, next5)

		if u[i] == next2 {
			p2 = p2 + 1
			next2 = u[p2] * 2
		}

		if u[i] == next3 {
			p3 = p3 + 1
			next3 = u[p3] * 3
		}

		if u[i] == next5 {
			p5 = p5 + 1
			next5 = u[p5] * 5
		}
	}

	return u[n-1]

}

func search(left, right, target int, data []int) bool {
	var mid int = (left + right) / 2

	if left > right {
		return false
	}

	if data[mid] == target {
		return true
	}

	if mid < target {
		left = mid + 1
	} else {
		right = mid - 1
	}

	return search(left, right, target, data)
}

func BST(data []int, target int) bool {
	left := 0
	right := len(data) - 1

	return search(left, right, target, data)
}
