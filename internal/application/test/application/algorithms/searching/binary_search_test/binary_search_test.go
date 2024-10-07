package binary_search_test

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/algorithms/searching/binary_search"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := [...]int{1, 2, 4, 5, 8, 12}
	result := binary_search.BinarySearch(arr[:], 4)
	expected := 2
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
