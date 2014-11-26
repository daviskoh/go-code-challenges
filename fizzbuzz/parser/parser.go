package parser

import (
	"strconv"
	"strings"
)

type formatter func(x, y, n string) string

// convert line to array
// loop through array
// map to fizzbuzz num
// return array.join
func Parser(line string, fn formatter) string {
	var result string

	lineArr := strings.Split(line, " ")
	num, _ := strconv.Atoi(lineArr[2])

	for i := 1; i <= num; i++ {
		iString := strconv.Itoa(i)

		result += fn(lineArr[0], lineArr[1], iString) + " "
	}

	return result[0 : len(result)-1]
}
