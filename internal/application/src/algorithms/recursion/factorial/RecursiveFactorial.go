package factorial

func RecursiveFactorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * RecursiveFactorial(n-1)
	}
}
