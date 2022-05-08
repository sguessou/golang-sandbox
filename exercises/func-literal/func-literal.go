//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

func handleStrings(lines []string) {
	stringsInfo := map[string]int{
		"letters":     0,
		"digits":      0,
		"spaces":      0,
		"punctuation": 0,
	}

	makeReport := func(r rune) {
		switch {
		case unicode.IsDigit(r):
			stringsInfo["digits"] += 1
		case unicode.IsLetter(r):
			stringsInfo["letters"] += 1
		case unicode.IsPunct(r):
			stringsInfo["punctuation"] += 1
		case unicode.IsSpace(r):
			stringsInfo["spaces"] += 1
		default:
			fmt.Println("Unrecognized character...")
		}
	}
	for _, line := range lines {
		for _, r := range line {
			makeReport(r)
		}
	}

	printReport := func(report map[string]int) {
		fmt.Printf("There are:\n%v letters\n%v digits\n%v spaces\nand %v punctuation marks in these lines of text.",
			report["letters"], report["digits"], report["spaces"], report["punctuation"])
	}

	printReport(stringsInfo)
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	handleStrings(lines)
}
