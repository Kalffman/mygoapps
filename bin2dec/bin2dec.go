package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func SBinToDec(sBin string) (int, error) {
	regx, _ := regexp.Compile("^[0-1]{8}$")

	if !regx.MatchString(sBin) {
		return -1, errors.New("is not a binary string")
	}

	binArr := strings.Split(sBin, "")

	var total int

	for idx, sIdx := range binArr {
		iIdx, _ := strconv.Atoi(sIdx)

		total += iIdx * calcExpoBase2(7-idx)
	}

	return total, nil
}

func calcExpoBase2(expo int) int {
	return int(math.Pow(2, float64(expo)))
}

func main() {

	numbers := os.Args[1:]

	for _, num := range numbers {
		result, err := SBinToDec(num)

		if err != nil {
			fmt.Printf("%s -> %s\n", num, err)
		} else {
			fmt.Printf("%s -> %d\n", num, result)
		}

	}
}
