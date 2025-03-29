package basic_problems

func LengthOfLongestString(inputString string) string {
	runes := []rune(inputString)
	lastSeen := make(map[rune]int)
	start := 0
	maxLength := 0
	maxStart := 0

	for i, char := range runes {
		if last, found := lastSeen[char]; found && last >= start {
			start = last + 1
		}
		lastSeen[char] = i

		currentLength := i - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
			maxStart = start
		}
	}
	longestSubString := string(runes[maxStart : maxStart+maxLength])
	print(longestSubString)
	return longestSubString
}
