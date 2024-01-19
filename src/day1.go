package src

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var numbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func SumOfCalibrationValue() int {
	data := utils.ReadFromFile("inputs/day1.txt")
	calib_value := 0
	for _, line := range data {
		var curr = []string{}
		for _, char := range line {
			if unicode.IsDigit(char) {
				curr = append(curr, string(char))
			}
		}
		size := len(curr)
		if size >= 1 {
			n, err := strconv.Atoi(curr[0] + curr[size-1])
			if err != nil {
				panic(err)
			}
			calib_value += n
		}
	}
	return calib_value
}

func replaceFirstOccurence(line string) string {
	cp_line := ""
	for _, char := range line {
		cp_line += string(char)
		cp_line = replaceMultipleWords(cp_line)
	}
	return cp_line
}

func replaceLastOccurence(line string) string {
	cp_line := ""
	for i := len(line) - 1; i >= 0; i-- {
		cp_line = string(line[i]) + cp_line
		cp_line = replaceMultipleWords(cp_line)
	}
	return cp_line
}

func replaceMultipleWords(line string) string {
	for i := 0; i < len(numbers); i++ {
		line = strings.ReplaceAll(line, numbers[i], strconv.Itoa(i+1))
	}
	return line
}

func SumOfCalibrationValueWithString() int {
	data := utils.ReadFromFile("inputs/day1.txt")
	calib_value := 0
	for _, line := range data {
		first := getCurrFirst(line)
		last := getCurrLast(line)
		size := len(last)
		if size >= 1 {
			n, err := strconv.Atoi(first[0] + last[size-1])
			if err != nil {
				panic(err)
			}
			fmt.Println(n)
			calib_value += n
		}
	}
	return calib_value
}

func getCurrLast(line string) []string {
	var curr = []string{}
	line = replaceLastOccurence(line)
	for _, char := range line {
		if unicode.IsDigit(char) {
			curr = append(curr, string(char))
		}
	}
	return curr
}

func getCurrFirst(line string) []string {
	var curr = []string{}
	line = replaceFirstOccurence(line)
	for _, char := range line {
		if unicode.IsDigit(char) {
			curr = append(curr, string(char))
		}
	}
	return curr
}
