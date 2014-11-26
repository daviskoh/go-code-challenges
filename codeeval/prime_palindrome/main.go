package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// print largest prime palindrome < 1000

func reverse(s string) string {
	var buffer bytes.Buffer

	for i := len(s) - 1; i >= 0; i-- {
		buffer.WriteString(string(s[i]))
	}

	return buffer.String()
}

func isPalindrome(s string) bool {
	return s == reverse(s)
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	// loop through 2...n-1
	// if evenly divisible by i then return false
	for i := 2; i < n-1; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	// keep counting down -1
	// if isPalindrome then return num

	// start at 1000
	result := 999
	for {
		if isPalindrome(strconv.Itoa(result)) == true && isPrime(result) == true {
			fmt.Println(result)
			break
		}

		result = result - 1
	}
}
