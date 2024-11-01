package parenthesis_validation

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/stack/double_linked_list"
)

func ValidateIfParenthesisAreBalanced(aStringOfParenthesisBracketsAndOrCurlyBraces string) bool {
	auxiliaryStack := double_linked_list.CreateNewStackAsDoubleLinkedList()
	stringRepresentedAsArray := []rune(aStringOfParenthesisBracketsAndOrCurlyBraces)

	for i := 0; i < len(stringRepresentedAsArray); i++ {
		currentCharacter := stringRepresentedAsArray[i]
		if auxiliaryStack.Empty() && isCharacterThatCloses(currentCharacter) {
			return false
		} else if isACharacterThatOpens(currentCharacter) {
			auxiliaryStack.Push(currentCharacter)
		} else {
			auxiliaryStack.Pop()
		}
	}

	return auxiliaryStack.Empty()
}

func isACharacterThatOpens(character rune) bool {
	charactersThatOpens := []string{"(", "[", "{"}
	for i := 0; i < len(charactersThatOpens); i++ {
		if charactersThatOpens[i] == string(character) {
			return true
		}
	}
	return false
}

func isCharacterThatCloses(character rune) bool {
	charactersThatCloses := []string{")", "]", "}"}

	for i := 0; i < len(charactersThatCloses); i++ {
		if charactersThatCloses[i] == string(character) {
			return true
		}
	}
	return false
}
