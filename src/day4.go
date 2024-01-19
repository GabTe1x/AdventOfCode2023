package src

import (
	"adventofcode/utils"
	"strconv"
	"strings"
)

type scratchCard struct {
	id        int
	w_numbers []int
	m_numbers []int
}

func (sc *scratchCard) correspondingCards() int {
	score := 0
	for _, curr_nbr := range sc.m_numbers {
		for _, w_nbr := range sc.w_numbers {
			if w_nbr == curr_nbr {
				score++
			}
		}
	}
	return score
}

func getCardsScore(start int, end int, cards []scratchCard) int {
	score := 1
	for i := start; i <= end; i++ {
		nbr := cards[i].correspondingCards()
		score += getCardsScore(i+1, i+nbr, cards)
	}
	return score
}

func NumberOfCards() int {
	scratchCards := parseData()
	score := 0
	for i, card := range scratchCards {
		rec_call := card.correspondingCards()
		score += getCardsScore(i+1, i+rec_call, scratchCards)
	}
	return score
}

func parseData() []scratchCard {
	data := utils.ReadFromFile("inputs/day4.txt")
	var res []scratchCard
	for _, line := range data {
		var card scratchCard
		splitdata := strings.Split(line, ": ")
		card.id = getId(splitdata[0])

		splitdata = strings.Split(splitdata[1], " | ")
		card.w_numbers = getSortedList(splitdata[0])
		card.m_numbers = getSortedList(splitdata[1])
		res = append(res, card)
	}
	return res
}

func getId(line_start string) int {
	id := strings.Fields(line_start)[1]
	res, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	return res
}

func getSortedList(list string) []int {
	res := []int{}
	str_list := strings.Fields(list)
	for _, el := range str_list {
		n, err := strconv.Atoi(el)
		if err != nil {
			panic(err)
		}
		res = append(res, n)
	}
	//sort.Ints(res)
	return res
}

func (sc *scratchCard) getScore() int {
	score := 0
	for _, curr_nbr := range sc.m_numbers {
		for _, w_nbr := range sc.w_numbers {
			if w_nbr == curr_nbr {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
	}
	return score
}

func SumOfCards() int {
	scratchCards := parseData()
	score := 0
	for _, card := range scratchCards {
		score += card.getScore()
	}
	return score
}
