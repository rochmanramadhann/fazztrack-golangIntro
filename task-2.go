package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

const (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

const (
	LOW    string = "LOW"
	MEDIUM        = "MEDIUM"
	STRONG        = "STRONG"
)

func main() {
	rand.Seed(time.Now().Unix())
	var userInput string

	fmt.Print("Please input any number : ")
	// Taking user input
	fmt.Scanln(&userInput)

	var isAlphabeticalChar = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(userInput)

	if isAlphabeticalChar {
		resultLowPassword := generatePassword(userInput, LOW)
		fmt.Println("LOW Password    :", resultLowPassword)

		resultMediumPassword := generatePassword(userInput, MEDIUM)
		fmt.Println("MEDIUM Password :", resultMediumPassword)

		resultStrongPassword := generatePassword(userInput, STRONG)
		fmt.Println("STRONG Password :", resultStrongPassword)
	} else {
		fmt.Println("Input must alphabetic characters!")
	}
}

func generatePassword(input, level string) string {
	minSpecialChar := 0
	minNum := 0
	minUpperCase := 0
	passwordLength := 0

	if len(input) < 6 {
		panic("Password must greater than equal 6!")
	}

	switch level {
	case "LOW":
		minSpecialChar = 2
		minNum = 2
		minUpperCase = 2
		passwordLength = 10
	case "MEDIUM":
		minSpecialChar = 4
		minNum = 4
		minUpperCase = 4
		passwordLength = 15
	case "STRONG":
		minSpecialChar = 6
		minNum = 6
		minUpperCase = 6
		passwordLength = 25
	}

	var password strings.Builder

	// Set random special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	// Set random numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	// Set random uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	// Set random string for remaining length with all character
	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}

	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	inputResult := setUppercase(input)
	return inputResult + string(inRune)
}

func setUppercase(input string) string {
	var password strings.Builder

	for i := 0; i < len(input); i++ {
		random := rand.Intn(len(input) - 1)

		if random%2 == 1 {
			password.WriteString(strings.ToUpper(string(input[i])))
		} else {
			password.WriteString(string(input[i]))
		}
	}

	inRune := []rune(password.String())

	return string(inRune)
}
