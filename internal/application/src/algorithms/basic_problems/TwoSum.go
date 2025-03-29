package basic_problems

func TwoSum(arr []int32, target int32) [][]int32 {
	indexes := make([][]int32, 0)

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			targetSum := arr[i] + arr[j]
			if targetSum == target {
				temporalArray := []int32{int32(i), int32(j)}
				indexes = append(indexes, temporalArray)
			}
		}
	}
	return indexes
}

func TwoSumWithMap(arr []int32, target int32) [][]int32 {
	indexes := make([][]int32, 0)
	visited := make(map[int32][]int32)
	for position, value := range arr {
		mapPointer := target - value
		if prevPositions, exists := visited[mapPointer]; exists {
			for _, prevPos := range prevPositions {
				indexes = append(indexes, []int32{prevPos, int32(position)})
			}

		}
		visited[value] = append(visited[value], int32(position))
	}
	return indexes
}

func convertArrayToMap(arr []int32) map[int32]int32 {
	resultMap := make(map[int32]int32)
	for position, value := range arr {
		resultMap[value] = int32(position)
	}
	return resultMap
}
