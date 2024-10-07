package binary_search

func BinarySearch(arr []int, target int) int {
	var lastElement = len(arr)
	firstElement := 0
	for firstElement <= lastElement {
		middleElement := (firstElement + lastElement) / 2
		if arr[middleElement] == target {
			return middleElement
		} else if arr[middleElement] > target {
			lastElement = middleElement
		} else if arr[middleElement] < target {
			firstElement = middleElement + 1
		}
	}
	return -1
}
