package fizzbuzz

import (
	"strconv"
)

func FizzBuzz(x, y, n string) string {
	nAsInt, _ := strconv.Atoi(n)

	if nAsInt == 0 {
		return n
	}

	xAsInt, _ := strconv.Atoi(x)
	yAsInt, _ := strconv.Atoi(y)

	if nAsInt%xAsInt == 0 {
		return "F"
	}

	if nAsInt%yAsInt == 0 {
		return "B"
	}

	return n
}
