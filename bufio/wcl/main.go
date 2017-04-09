package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func main() {
	var input io.Reader
	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Printf("Cannot open file: %v\n", err)
			return
		}
		input = file
		defer file.Close()
	} else {
		input = os.Stdin
	}
	scanner := bufio.NewScanner(input)
	linesNumber := 0
	for scanner.Scan() {
		scanner.Text()
		linesNumber++
	}
	fmt.Println(linesNumber)	
}

