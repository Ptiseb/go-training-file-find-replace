package main

import (
	"bytes"
	"fmt"
	"bufio"
	"os"
)

// src: source file name
// old/new : old/new word
// occ: number of occurencs of old
// lines: slices containing the line number where old was found
// err: error of the function
func FindReplaceFile(src, old, new string) (occ int, lines []int, err error) {
	data, err := os.ReadFile(src)

	if err != nil {
		return 0, make([]int, 0), err
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		l := scanner.Text()
		fmt.Println(l)
	}

	// scanner := 

	return 0, make([]int, 0), nil
}

// line: line of the file
// old/new : old/new word
// found: true if at last 1 occurence has been found
// res: result of the replacment (res == line if no update)
// occ: number of occurences in line
// func ProcessLine(line, old, new string) (found bool, res string, occ int)


func main() {
	fmt.Println("Hello world!")
	_, _, err := FindReplaceFile("file.txt", "", "")

	if err != nil {
		fmt.Printf("Error while reading file: %v\n", err)
		return
	}

	fmt.Println("Goodbye world.")
}