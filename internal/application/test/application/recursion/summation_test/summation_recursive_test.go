package summation_test

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/recursion/summation"
	"testing"
)

func TestSum(t *testing.T) {
	result := summation.RecursiveSummation(10)
	expected := 55
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
