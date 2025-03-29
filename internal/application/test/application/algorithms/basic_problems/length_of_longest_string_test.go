package basic_problems

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/algorithms/basic_problems"
	"testing"
)

func TestLengthOfLongestString(t *testing.T) {

	t.Run("giving the string \"abcabcbb\", the output must be abc", func(T *testing.T) {
		testString := "abcabcbb"
		result := basic_problems.LengthOfLongestString(testString)
		if len(result) != 3 {
			t.Error("The result must be 3")
		}
	})

	t.Run("giving the string \"aywowklmdfki\", the output must be 1", func(T *testing.T) {
		testString := "aywowklmdfki"
		result := basic_problems.LengthOfLongestString(testString)
		if result != "owklmdf" {
			t.Error("The result must be 3")
		}
	})
}
