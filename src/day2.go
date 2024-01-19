package src

import (
	"adventofcode/utils"
	"strconv"
	"strings"
)

type dice_bag struct {
	id        int
	red_max   int
	blue_max  int
	green_max int
}

func parseFile() []dice_bag {
	data := utils.ReadFromFile("inputs/day2.txt")
	var games = []dice_bag{}
	for _, line := range data {
		dice_game := parseLine(line)
		games = append(games, dice_game)
	}
	return games
}

func parseLine(line string) dice_bag {

	split_id := strings.Split(line, ":")
	split_name_id := strings.Split(split_id[0], " ")
	id, err := strconv.Atoi(split_name_id[1])
	if err != nil {
		panic(err)
	}
	temp := dice_bag{id: id, blue_max: 0, red_max: 0, green_max: 0}

	rounds := strings.Split(split_id[1], ";")

	for _, round := range rounds {
		res := strings.Split(round, ",")
		for _, dice_throw := range res {
			dice_split := strings.Split(dice_throw, " ")
			dice_nbr, err := strconv.Atoi(dice_split[1])
			if err != nil {
				panic(err)
			}
			switch {
			case dice_split[2] == "red":
				if dice_nbr > temp.red_max {
					temp.red_max = dice_nbr
				}
			case dice_split[2] == "blue":
				if dice_nbr > temp.blue_max {
					temp.blue_max = dice_nbr
				}
			case dice_split[2] == "green":
				if dice_nbr > temp.green_max {
					temp.green_max = dice_nbr
				}
			default:
			}
		}
	}

	return temp
}

func SumIdPossibleGames(red int, green int, blue int) int {
	games := parseFile()
	res := 0
	for _, dice_bag := range games {
		if dice_bag.blue_max <= blue && dice_bag.red_max <= red && dice_bag.green_max <= green {
			res += dice_bag.id
		}
	}
	return res
}

func SumPowerOfGames() int {
	res := 0
	games := parseFile()
	for _, dice_bag := range games {
		res += (dice_bag.blue_max * dice_bag.red_max * dice_bag.green_max)
	}
	return res
}
