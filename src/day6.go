package src

import (
	"adventofcode/utils"
	"strconv"
	"strings"
)

type race struct {
	race_time    int64
	distance_max int64
}

func (r *race) getAllWaysToWin() int {
	nbr_ways := 0
	for i := 0; i < int(r.race_time); i++ {
		time_left := int(r.race_time) - i
		if i*time_left > int(r.distance_max) {
			nbr_ways++
		}
	}
	return nbr_ways
}

func (r *race) getAllWaysToWin64() int64 {
	nbr_ways := int64(0)
	for i := int64(0); i < r.race_time; i++ {
		time_left := r.race_time - i
		if i*time_left > r.distance_max {
			nbr_ways++
		}
	}
	return nbr_ways
}

func BoatRace() int {
	data := utils.ReadFromFile("inputs/day6.txt")
	res := 0
	races := parseData6(data)
	for _, race := range races {
		if res == 0 {
			res = race.getAllWaysToWin()
		} else {
			res *= race.getAllWaysToWin()
		}
	}
	return res
}

func BigBoatRace() int64 {
	data := utils.ReadFromFile("inputs/day6.txt")
	race := correctData6(data)
	return race.getAllWaysToWin64()
}

func correctData6(data []string) race {
	var r race
	values := strings.Split(data[0], ":")
	values = strings.Fields(values[1])
	total_value := ""
	for _, value := range values {
		total_value += value
	}
	n, err := strconv.ParseInt(total_value, 10, 64)
	if err != nil {
		panic(err)
	}
	r.race_time = n
	total_value = ""
	values = strings.Split(data[1], ":")
	values = strings.Fields(values[1])
	for _, value := range values {
		total_value += value
	}
	n, err = strconv.ParseInt(total_value, 10, 64)
	if err != nil {
		panic(err)
	}
	r.distance_max = n
	return r
}

func parseData6(data []string) [4]race {
	var races [4]race
	values := strings.Split(data[0], ":")
	values = strings.Fields(values[1])
	for i, value := range values {
		n, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			panic(err)
		}
		races[i].race_time = n
	}
	values = strings.Split(data[1], ":")
	values = strings.Fields(values[1])
	for i, value := range values {
		n, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			panic(err)
		}
		races[i].distance_max = n
	}
	return races
}
