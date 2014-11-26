package main

import (
	"bufio"
	"fmt"
	// "github.com/daviskoh/go-code-challenges/fizzbuzz"
	"os"
)

func main() {
	file, err := os.Open("data")
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}

	fileReader := bufio.NewReader(file)

	for {
		line, _ := fileReader.ReadString('\n')

		if line == "" {
			break
		}

		fmt.Print(line)
	}
	fmt.Println()

	file.Close()
}
