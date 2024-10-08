package fibonacci

func RecursiveFibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return RecursiveFibonacci(n-1) + RecursiveFibonacci(n-2)
	}
}
