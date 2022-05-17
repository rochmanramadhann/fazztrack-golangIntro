package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type PatternValidation func(int) bool

func main() {
	var userInput string

	fmt.Print("Please input any number : ")
	// Taking user input
	fmt.Scanln(&userInput)
	var isDigit = regexp.MustCompile(`^[0-9]+$`).MatchString(userInput)

	if isDigit {
		if str, err := strconv.Atoi(userInput); err == nil {
			reverseHillPattern(str, patternLengthValidation)
		}
	} else {
		fmt.Print("Input must be a number!")
	}
}

func patternLengthValidation(length int) bool {
	if length < 3 {
		return false
	} else {
		return true
	}
}

func reverseHillPattern(length int, patternLengthValidation PatternValidation) {
	/*
			*********
			 *******
			  *****
		       ***
				*
	*/

	if patternLengthValidation(length) {
		for i := 0; i < length; i++ {
			for j := 0; j < i; j++ {
				fmt.Print(" ")
			}
			for j := i; j < length-1; j++ {
				fmt.Print("*")
			}
			for j := i; j < length; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}

		for i := 1; i <= length; i++ {
			for j := 1; j < i; j++ {
				fmt.Print(" ")
			}
			for j := i; j < length; j++ {
				fmt.Print("*")
			}
			for j := i; j <= length; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Cannot input less than 3!")
	}
}

// https://stackoverflow.com/a/47317150
