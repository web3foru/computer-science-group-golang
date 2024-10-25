package applications

import (
	"github.com/web3foru/computer-science-group-golang/internal/application/src/data_structures/stack/applications/parenthesis_validation"
	"testing"
)

func TestParenthesisValidator(t *testing.T) {
	t.Run("must validate a single opened and closed parenthesis", func(t *testing.T) {
		expression := "()"

		result := parenthesis_validation.ValidateIfParenthesisAreBalanced(expression)
		if result != true {
			t.Errorf("result should be true")
		}
	})

	t.Run("must validate multiple closed parenthesis", func(t *testing.T) {
		expression := "(((((((((())))))))))"

		result := parenthesis_validation.ValidateIfParenthesisAreBalanced(expression)
		if result != true {
			t.Errorf("result should be true")
		}
	})

	t.Run("must validate another multiple closed parenthesis", func(t *testing.T) {
		expression := "()()()(())((()))"

		result := parenthesis_validation.ValidateIfParenthesisAreBalanced(expression)
		if result != true {
			t.Errorf("result should be true")
		}
	})

	t.Run("must validate an incorrect string", func(t *testing.T) {
		expression := ")))))))"

		result := parenthesis_validation.ValidateIfParenthesisAreBalanced(expression)
		if result != false {
			t.Errorf("result should be true")
		}
	})
}
