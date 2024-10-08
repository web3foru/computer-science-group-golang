package binary_search_recursive

func RecursiveBinarySearch(arr []int, lowerLimit int, upperLimit int, target int) int {
	if upperLimit >= lowerLimit {
		middleOfArray := lowerLimit + (upperLimit-lowerLimit)/2
		if arr[middleOfArray] == target {
			return middleOfArray
		}
		if arr[middleOfArray] > target {
			return RecursiveBinarySearch(arr, lowerLimit, middleOfArray-1, target)
		}
		return RecursiveBinarySearch(arr, middleOfArray+1, upperLimit, target)
	}
	return -1
}
