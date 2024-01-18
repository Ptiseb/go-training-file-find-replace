package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
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

	i := 1

	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		l := scanner.Text()
		found, res, o := ProcessLine(l, old+" ", new+" ")

		if found {
			occ += o
			lines = append(lines, i)
		}

		fmt.Println(res)

		i++
	}

	return occ, lines, nil
}

// line: line of the file
// old/new : old/new word
// found: true if at last 1 occurence has been found
// res: result of the replacment (res == line if no update)
// occ: number of occurences in line
func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	// c := strings.Contains(line, old) // true or false
	// cnt = strings.Count(line, old)   // number
	// res := strings.Replace(line, old, new, 0) // res

	if !strings.Contains(line, old) {
		return false, line, 0
	}

	return true, strings.Replace(line, old, new, -1), strings.Count(line, old)
}

func main() {
	fmt.Println("Hello world!")
	occ, lines, err := FindReplaceFile("wikigo.txt", "Go", "rust")

	if err != nil {
		fmt.Printf("Error while reading file: %v\n", err)
		return
	}

	fmt.Printf("Number of occurences of go: %d\n", occ)
	fmt.Printf("Number of lines: %d\n", lines)

	fmt.Println("Goodbye world.")
}
