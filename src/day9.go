package src

import (
	"adventofcode/utils"
	"strconv"
	"strings"
)

func OasisReportSum() int {
	data := utils.ReadFromFile("inputs/day9.txt")
	reports := parseData9(data)
	res := 0
	for _, ints := range reports {
		res += predictLastReport(ints)
	}
	return res
}

func OasisReportStartSum() int {
	data := utils.ReadFromFile("inputs/day9.txt")
	reports := parseData9(data)
	res := 0
	for _, ints := range reports {
		res += predictFirstReport(ints)
	}
	return res
}

func parseData9(data []string) [][]int {
	var res [][]int
	for _, line := range data {
		temp := parseLine9(line)
		res = append(res, temp)
	}
	return res
}

func parseLine9(line string) []int {
	var res []int
	values := strings.Fields(line)
	for _, value := range values {
		n, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		res = append(res, n)
	}
	return res
}

func predictLastReport(values []int) int {
	res := 0
	if sum(values) == 0 {
		return res
	}
	var new_values []int
	for i := 0; i < len(values)-1; i++ {
		new_values = append(new_values, values[i+1]-values[i])
	}
	res = values[len(values)-1] + predictLastReport(new_values)
	return res
}

func predictFirstReport(values []int) int {
	res := 0
	if sum(values) == 0 {
		return res
	}
	var new_values []int
	for i := 0; i < len(values)-1; i++ {
		new_values = append(new_values, values[i+1]-values[i])
	}
	res = values[0] + -predictFirstReport(new_values)
	return res
}

func sum(array []int) int {
	res := 0
	for _, val := range array {
		res += Abs(val)
	}
	return res
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
