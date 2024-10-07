package summation

func RecursiveSummation(n int) int {
	if n == 0 {
		return 0
	}
	return n + RecursiveSummation(n-1)
}
