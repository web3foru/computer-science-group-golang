package summation_test

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/algorithms/recursion/summation"
	"testing"
)

func TestRecursiveSummation(t *testing.T) {
	result := summation.RecursiveSummation(10)
	expected := 55
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
