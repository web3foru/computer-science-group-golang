package basic_problems

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/algorithms/basic_problems"
	"testing"
)

func TestAbsoluteSumOfDigonal(t *testing.T) {
	t.Run("given a 3*3 matrix calculate", func(T *testing.T) {
		matrix := [][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}
		sum := basic_problems.AbsoluteDifferenceOfDiagonals(matrix)
		if sum != 15 {
			t.Error("The suum must be an absolute value")
		}
	})
}
