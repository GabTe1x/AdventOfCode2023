package src

import (
	"adventofcode/utils"
	"slices"
	"strconv"
	"strings"
)

type interval struct {
	src_start  int64
	dest_start int64
	length     int64
}

func Almanact() int64 {
	data := utils.ReadFromFile("inputs/day5.txt")
	seeds := parseSeeds(data[0])
	intervals := parseData10(data[1:])
	for i, seed := range seeds {
		curr_ctg := "seed"
		res := seed
		for curr_ctg != "location" {
			for key, value := range intervals {
				if checkCategory(curr_ctg, key) {
					curr_ctg = nextCategory(key)
					res = getNextValue(res, value)
				}
			}
		}
		seeds[i] = res
	}
	return slices.Min(seeds)
}

func parseSeeds(line string) []int64 {
	var res []int64
	data := strings.Fields(line)
	for i := 1; i < len(data); i++ {
		n, err := strconv.ParseInt(data[i], 10, 64)
		if err != nil {
			panic(err)
		}
		res = append(res, n)
	}
	return res
}

func parseData10(data []string) map[string][]interval {
	res := make(map[string][]interval)
	add := false
	curr := ""
	var value []interval
	for _, line := range data {
		split_data := strings.Fields(line)
		if len(split_data) == 0 {
			if len(value) != 0 && curr != "" {
				res[curr] = value
				value = []interval{}
			}
			add = true
		} else if add {
			add = false
			curr = split_data[0]
		} else {
			value = append(value, createInterval(split_data))
		}
	}
	res[curr] = value
	return res
}

func createInterval(val []string) interval {
	var res interval
	for i, el := range val {
		n, err := strconv.ParseInt(el, 10, 64)
		if err != nil {
			panic(err)
		}
		switch {
		case i == 0:
			res.dest_start = n
		case i == 1:
			res.src_start = n
		default:
			res.length = n
		}
	}
	return res
}

func checkCategory(start string, category string) bool {
	split_category := strings.Split(category, "-to-")
	return split_category[0] == start
}

func nextCategory(category string) string {
	split_category := strings.Split(category, "-to-")
	return split_category[1]
}

func getNextValue(seed int64, convert []interval) int64 {
	res := seed
	for _, v := range convert {
		if seed >= v.src_start && seed <= v.src_start+v.length-1 {
			diff := seed - v.src_start
			return v.dest_start + diff
		}
	}
	return res
}
