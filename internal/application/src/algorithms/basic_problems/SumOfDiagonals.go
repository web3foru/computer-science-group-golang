package basic_problems

import "math"

func AbsoluteDifferenceOfDiagonals(arr [][]int32) int32 {

	var sumOfPrimaryDiagonal int32 = 0
	var sumOfSecondaryDiagonal int32 = 0
	var auxIndex = len(arr) - 1
	for i := 0; i < len(arr); i++ {
		sumOfPrimaryDiagonal += arr[i][i]
		sumOfSecondaryDiagonal += arr[i][auxIndex]
		auxIndex--
	}
	return int32(math.Abs(float64(sumOfPrimaryDiagonal) - float64(sumOfSecondaryDiagonal)))
}
