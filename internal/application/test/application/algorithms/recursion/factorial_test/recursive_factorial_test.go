package factorial_test

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/algorithms/recursion/factorial"
	"testing"
)

func TestRecursiveFactorial(t *testing.T) {
	result := factorial.RecursiveFactorial(10)
	expected := 3628800
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
