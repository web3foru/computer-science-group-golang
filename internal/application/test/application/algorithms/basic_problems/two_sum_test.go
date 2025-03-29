package basic_problems

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/algorithms/basic_problems"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	t.Run("given an array of 3 elements, give the first two positions", func(T *testing.T) {
		arrayOfIntegers := []int32{11, 2, 4}
		arrayOfResult := []int32{0, 1}
		const target int32 = 13
		result := basic_problems.TwoSum(arrayOfIntegers, target)
		if reflect.DeepEqual(arrayOfResult, result[0]) == false {
			t.Error("indexes are incorrect!")
		}
	})

	t.Run("given an array of 10 elements, give 3 results", func(T *testing.T) {
		arrayOfIntegers := []int32{11, 2, 4, 7, 2, 3, 3, 6}
		const target int32 = 13
		result := basic_problems.TwoSum(arrayOfIntegers, target)
		if len(result) != 3 {
			t.Error("indexes are incorrect!")
		}
	})

	t.Run("given an array of 3 elements, give the first two positions, with a more efficient algorithm ", func(T *testing.T) {
		arrayOfIntegers := []int32{11, 2, 4}
		arrayOfResult := []int32{0, 1}
		const target int32 = 13
		result := basic_problems.TwoSumWithMap(arrayOfIntegers, target)
		if reflect.DeepEqual(arrayOfResult, result[0]) == false {
			t.Error("indexes are incorrect!")
		}
	})

	t.Run("given an array of 10 elements, give 3 results, , with a more efficient algorithm ", func(T *testing.T) {
		arrayOfIntegers := []int32{11, 2, 4, 7, 2, 3, 6, 11}
		const target int32 = 13
		result := basic_problems.TwoSumWithMap(arrayOfIntegers, target)
		if len(result) != 5 {
			t.Error("indexes are incorrect!")
		}
	})
}
