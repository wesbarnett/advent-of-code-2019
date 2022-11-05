package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func validatePasswordPart2(password int) bool {

	passwordStr := strconv.Itoa(password)
	currRepeatLength := 1
	repeatedDigit := false
	for i := 1; i < len(passwordStr); i++ {
		curr, err := strconv.Atoi(string(passwordStr[i]))
		if err != nil {
			log.Fatal(err)
		}
		prev, err := strconv.Atoi(string(passwordStr[i-1]))
		if err != nil {
			log.Fatal(err)
		}
		// Decreases so invalid
		if curr < prev {
			return false
		}
		// Required repeating digit
		if curr == prev {
			currRepeatLength += 1
		} else {
			if currRepeatLength == 2 {
				repeatedDigit = true
			}
			currRepeatLength = 1
		}
	}
	// Handle if repeating sequence is at the end
	if currRepeatLength == 2 {
		repeatedDigit = true
	}
	return repeatedDigit

}

func validatePassword(password int) bool {

	passwordStr := strconv.Itoa(password)
	repeatedDigit := false
	for i := 1; i < len(passwordStr); i++ {
		curr, err := strconv.Atoi(string(passwordStr[i]))
		if err != nil {
			log.Fatal(err)
		}
		prev, err := strconv.Atoi(string(passwordStr[i-1]))
		if err != nil {
			log.Fatal(err)
		}
		// Decreases so invalid
		if curr < prev {
			return false
		}
		// Required repeating digit
		if curr == prev {
			repeatedDigit = true
		}
	}
	return repeatedDigit

}

func main() {
	var file string
	flag.StringVar(&file, "infile", "input", "Input file")
	flag.Parse()

	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.Trim(string(content), "\n"), "-")
	lowerBound, err := strconv.Atoi(input[0])
	if err != nil {
		log.Fatal(err)
	}
	upperBound, err := strconv.Atoi(input[1])
	if err != nil {
		log.Fatal(err)
	}

	validPasswordCount := 0
	for password := lowerBound; password <= upperBound; password++ {
		if validatePassword(password) {
			validPasswordCount += 1
		}
	}
	fmt.Println(validPasswordCount)

	validPasswordCountPart2 := 0
	for password := lowerBound; password <= upperBound; password++ {
		if validatePasswordPart2(password) {
			validPasswordCountPart2 += 1
		}
	}
	fmt.Println(validPasswordCountPart2)

}
