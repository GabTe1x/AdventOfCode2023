package src

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

type tuple struct {
	left  string
	right string
}

func StepsNumber() int {
	data := utils.ReadFromFile("inputs/day8.txt")
	res := 0
	instructions := data[0]
	mymap := parseData8(data[2:])
	curr := "AAA"
	for i := 0; i < len(data); i++ {
		char := instructions[i]
		if string(char) == "L" {
			curr = mymap[curr].left
		} else {
			curr = mymap[curr].right
		}
		res += 1
		if curr == "ZZZ" {
			break
		} else if i+1 == len(instructions) {
			i = -1
		}
	}
	return res
}

func StepsNumber2() int {
	data := utils.ReadFromFile("inputs/day8.txt")
	res := 0
	instructions := data[0]
	mymap := parseData8(data[2:])
	curr := getStartingPoints(mymap)
	for i := 0; i < len(data); i++ {
		char := instructions[i]
		if string(char) == "L" {
			for i := 0; i < len(curr); i++ {
				curr[i] = mymap[curr[i]].left
			}
		} else {
			for i := 0; i < len(curr); i++ {
				curr[i] = mymap[curr[i]].right
			}
		}
		res += 1
		cmpt := 0
		for i := 0; i < len(curr); i++ {
			if strings.HasSuffix(curr[i], "Z") {
				cmpt++
			}
		}
		if cmpt == len(curr) {
			fmt.Println(curr, cmpt)
			break
		} else if i == len(instructions)-1 {
			i = -1
		}
	}
	return res
}

func parseData8(data []string) map[string]tuple {
	mymap := make(map[string]tuple)
	for _, line := range data {
		key := strings.Split(line, " = ")
		m_data := strings.Replace(key[1], "(", "", 1)
		m_data = strings.Replace(m_data, ")", "", 1)
		split_data := strings.Split(m_data, ", ")
		var value tuple
		value.right = split_data[1]
		value.left = split_data[0]
		mymap[key[0]] = value
	}
	return mymap
}

func getStartingPoints(data map[string]tuple) []string {
	var res []string
	for el := range data {
		if strings.HasSuffix(el, "A") {
			res = append(res, el)
		}
	}
	fmt.Println(res[:])
	return res
}
