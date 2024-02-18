package utils

func Sum(a, b int) int {
	return a + b
}

func Difference(a, b int) int {
	return a - b
}

func EvenORodd(number int) string {
	if number%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}
