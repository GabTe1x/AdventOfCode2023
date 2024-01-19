package src

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

func SumOfEngine() int {
	data := utils.ReadFromFile("inputs/day3.txt")
	res := 0
	for i := 0; i < len(data); i++ {
		curr_nbr := ""
		add := false
		for y := 0; y < len(data[i]); y++ {
			if unicode.IsDigit(rune(data[i][y])) {
				curr_nbr += string(data[i][y])
				if !add {
					add = checkSurroundings(i, y, data)
				}
			} else {
				// if number whole
				if add {
					x, err := strconv.Atoi(curr_nbr)
					if err != nil {
						panic(err)
					}
					res += x
					add = false
				}
				curr_nbr = ""
			}
		}
		// if end of line
		if add {
			x, err := strconv.Atoi(curr_nbr)
			if err != nil {
				panic(err)
			}
			res += x
			add = false
		}
		curr_nbr = ""
	}
	return res
}

func checkSurroundings(i int, y int, data []string) bool {
	res := false
	switch {
	case i == 0:
		switch {
		case y == 0:
			if isSymbol(string(data[i][y+1])) || isSymbol(string(data[i+1][y+1])) || isSymbol(string(data[i+1][y])) {
				res = true
			}
		case y == len(data[i])-1:
			if isSymbol(string(data[i][y-1])) || isSymbol(string(data[i+1][y])) || isSymbol(string(data[i+1][y-1])) {
				res = true
			}
		default:
			if isSymbol(string(data[i][y-1])) || isSymbol(string(data[i][y+1])) || isSymbol(string(data[i+1][y+1])) || isSymbol(string(data[i+1][y])) || isSymbol(string(data[i+1][y-1])) {
				res = true
			}
		}
	case i == len(data)-1:
		switch {
		case y == 0:
			if isSymbol(string(data[i][y+1])) || isSymbol(string(data[i-1][y+1])) || isSymbol(string(data[i-1][y])) {
				res = true
			}
		case y == len(data[i])-1:
			if isSymbol(string(data[i][y-1])) || isSymbol(string(data[i-1][y])) || isSymbol(string(data[i-1][y-1])) {
				res = true
			}
		default:
			if isSymbol(string(data[i][y-1])) || isSymbol(string(data[i][y+1])) || isSymbol(string(data[i-1][y+1])) || isSymbol(string(data[i-1][y])) || isSymbol(string(data[i-1][y-1])) {
				res = true
			}
		}
	default:
		switch {
		case y == 0:
			if isSymbol(string(data[i][y+1])) || isSymbol(string(data[i+1][y+1])) || isSymbol(string(data[i+1][y])) || isSymbol(string(data[i-1][y+1])) || isSymbol(string(data[i-1][y])) {
				res = true
			}
		case y == len(data[i])-1:
			if isSymbol(string(data[i][y-1])) || isSymbol(string(data[i+1][y])) || isSymbol(string(data[i+1][y-1])) || isSymbol(string(data[i-1][y])) || isSymbol(string(data[i-1][y-1])) {
				res = true
			}
		default:
			if isSymbol(string(data[i][y-1])) || isSymbol(string(data[i][y+1])) || isSymbol(string(data[i+1][y+1])) || isSymbol(string(data[i+1][y])) || isSymbol(string(data[i+1][y-1])) || isSymbol(string(data[i-1][y+1])) || isSymbol(string(data[i-1][y])) || isSymbol(string(data[i-1][y-1])) {
				res = true
			}
		}
	}
	return res
}

func isSymbol(char string) bool {
	reg, err := regexp.Compile("[^a-zA-Z0-9.]")
	if err != nil {
		panic(err)
	}
	return reg.Match([]byte(char))
}

func isStar(char string) bool {
	reg, err := regexp.Compile("[*]")
	if err != nil {
		panic(err)
	}
	return reg.Match([]byte(char))
}

func SumOfGear() int {
	data := utils.ReadFromFile("inputs/day3.txt")
	res := 0
	for i := 0; i < len(data); i++ {
		for y := 0; y < len(data[i]); y++ {
			if isStar(string(data[i][y])) {
				fmt.Println("*", i)
				res += checkAround(i, y, data)
			}
		}
	}
	return res
}

func checkAround(i int, y int, data []string) int {
	res := 0
	// up
	var x [8]int
	x[0] = getDigit(i+1, y+1, data)
	x[1] = getDigit(i+1, y, data)
	x[2] = getDigit(i+1, y-1, data)
	// left
	x[3] = getDigit(i, y-1, data)
	// right
	x[4] = getDigit(i, y+1, data)
	// down
	x[5] = getDigit(i-1, y+1, data)
	x[6] = getDigit(i-1, y, data)
	x[7] = getDigit(i-1, y-1, data)

	valide := 0
	for _, el := range x {
		if el != 0 && valide == 0 {
			fmt.Println(el)
			res = el
			valide++
		} else if el != 0 && valide == 1 {
			if el != res {
				fmt.Println(el)
				res = res * el
				valide++
			}
		}
	}
	if valide == 2 {
		fmt.Println(res)
		return res
	}
	return 0
}

func getDigit(i int, y int, data []string) int {
	res := 0
	if i < 0 || i == len(data) || y < 0 {
		return res

	}
	if y == len(data[i]) {
		return res
	}
	if unicode.IsDigit(rune(data[i][y])) {
		left := getLeft(i, y-1, data)
		right := getRight(i, y+1, data)
		res_string := left + string(data[i][y]) + right
		res, _ = strconv.Atoi(res_string)
	}
	return res
}

func getLeft(i int, y int, data []string) string {
	res := ""
	if y < 0 {
		return res
	}
	if unicode.IsDigit(rune(data[i][y])) {
		res = getLeft(i, y-1, data) + string(data[i][y])
	}
	return res
}

func getRight(i int, y int, data []string) string {
	res := ""
	if y == len(data[i]) {
		return res
	}
	if unicode.IsDigit(rune(data[i][y])) {
		res = string(data[i][y]) + getRight(i, y+1, data)
	}
	return res
}
