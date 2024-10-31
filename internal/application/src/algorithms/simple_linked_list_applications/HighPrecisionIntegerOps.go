package simple_linked_list_applications

type HighPrecisionIntegerOps interface {
	SetUpHighPrecisionInteger(number int) HighPrecisionInteger
	Sum(number1 HighPrecisionInteger, number2 HighPrecisionInteger) HighPrecisionInteger
	Difference(number1 HighPrecisionInteger, number2 HighPrecisionInteger) HighPrecisionInteger
}
