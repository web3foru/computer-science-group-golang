package fibonacci_test

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/algorithms/recursion/fibonacci"
	"testing"
)

func TestRecursiveSummation(t *testing.T) {
	var result = fibonacci.RecursiveFibonacci(10)
	const expected int = 55
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
