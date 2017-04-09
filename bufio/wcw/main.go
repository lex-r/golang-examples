package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	countWords := 0
	var input io.Reader
	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Printf("Cannot open file: %s\n", err)
			return
		}
		input = file
	} else {
		input = os.Stdin
	}

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		countWords++
	}

	fmt.Println(countWords)
}
