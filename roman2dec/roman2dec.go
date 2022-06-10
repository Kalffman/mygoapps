package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// 0 - 3999 | "(M{0,3})" -> THOUSANDS_PATTERN | "(D?C{0,3}|CD|D|CM)" -> HUNDREDS_PATTERN | "(L?X{0,3}|XL|L|XC)" -> TENS_PATTERN | "(V?I{0,3}|IV|V|IX)" -> ONES_PATTERN
const ROMAN_NUMERAL_PATTERN = "^^(M{0,3})(D?C{0,3}|CD|D|CM)(L?X{0,3}|XL|L|XC)(V?I{0,3}|IV|V|IX)$$"

var ROMAN_NUMERALS = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func convert(romanNumeral string) int {
	// fmt.Println("\n24 -", romanNumeral)
	var rNumbers = strings.Split(romanNumeral, "")
	// fmt.Println("26 -", rNumbers)

	var dNumbers = make([]int, len(rNumbers))
	for i, r := range rNumbers {
		dNumbers[i] = ROMAN_NUMERALS[r]
	}
	// fmt.Println("32 -", dNumbers)

	var total int
	for idx, actual := range dNumbers {
		if idx == 0 {
			total = actual
		} else if total < actual {
			total = actual - total
		} else {
			total += actual
		}
	}

	// fmt.Println("45 -", total)
	return total
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("At least one roman numeral required.\nEg.: MMXXII VI X")
		os.Exit(0)
	}

	for _, romanNumber := range args[1:] {

		regx, _ := regexp.Compile(ROMAN_NUMERAL_PATTERN)

		if !regx.MatchString(romanNumber) {
			fmt.Printf("%s -> is not valid roman numeral\n", romanNumber)
			os.Exit(0)
		}

		matchers := regx.FindStringSubmatch(romanNumber)[1:]

		var total int
		for _, value := range matchers {
			total += convert(value)
		}

		fmt.Printf("%s -> %d\n", romanNumber, total)
	}
}
