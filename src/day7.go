package src

import (
	"adventofcode/utils"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	mtype int
	cards [5]int
	bid   int
}

func TotalWinnings() int {
	res := 0
	data := utils.ReadFromFile("inputs/day7.txt")
	hands := parseData7(data)
	sort.Slice(hands, func(i int, j int) bool {
		if hands[i].mtype == hands[j].mtype {
			for z := 0; z < 5; z++ {
				if hands[i].cards[z] != hands[j].cards[z] {
					return hands[i].cards[z] < hands[j].cards[z]
				}
			}
		} else {
			return hands[i].mtype < hands[j].mtype
		}
		return false
	})
	for i, hand := range hands {
		res += (i + 1) * hand.bid
	}
	return res
}

/*
So, 33332 and 2AAAA are both four of a kind hands, but 33332 is stronger because
its first card is stronger.Similarly, 77888 and 77788 are both a full house, but
77888 is stronger because its third card is stronger (and  both hands  have  the
same first and second card).
*/

func parseData7(data []string) []hand {
	var hands []hand
	for _, line := range data {
		var curr_hand hand
		values := strings.Fields(line)
		bid, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}
		curr_hand.bid = bid
		curr_hand.cards = getCards(values[0])
		curr_hand.mtype = getType(curr_hand.cards)
		hands = append(hands, curr_hand)
	}
	return hands
}

func getCards(cards string) [5]int {
	var res [5]int
	for i, char := range cards {
		res[i] = getCard(string(char))
	}
	return res
}

func getCard(card string) int {
	n, err := strconv.Atoi(card)
	if err != nil {
		switch {
		case card == "T":
			n = 10
		case card == "J":
			n = 11
		case card == "Q":
			n = 12
		case card == "K":
			n = 13
		case card == "A":
			n = 14
		}
	}
	return n
}

func getType(cards [5]int) int {
	res := 0
	pair := 0
	triple := 0
	var previouscard []int
out:
	for i := 0; i < 5; i++ {
		nbr := 1
		if !slices.Contains(previouscard, cards[i]) {
			for j := i + 1; j < 5; j++ {
				if cards[i] == cards[j] {
					nbr++
				}
			}
			switch {
			case nbr == 5:
				res = 6
				break out
			case nbr == 4:
				res = 5
				break out
			case nbr == 3:
				triple++
			case nbr == 2:
				pair++
			default:
			}
			previouscard = append(previouscard, cards[i])
		}
	}
	if pair == 2 {
		res = 2
	} else if triple == 1 && pair == 1 {
		res = 4
	} else if triple == 1 {
		res = 3
	} else if pair == 1 {
		res = 1
	}
	return res
}

// PART 2

func TotalWinnings2() int {
	res := 0
	data := utils.ReadFromFile("inputs/day7.txt")
	hands := parseData7_2(data)
	sort.Slice(hands, func(i int, j int) bool {
		if hands[i].mtype == hands[j].mtype {
			for z := 0; z < 5; z++ {
				if hands[i].cards[z] != hands[j].cards[z] {
					return hands[i].cards[z] < hands[j].cards[z]
				}
			}
		} else {
			return hands[i].mtype < hands[j].mtype
		}
		return false
	})
	for i, hand := range hands {
		res += (i + 1) * hand.bid
	}
	return res
}

func parseData7_2(data []string) []hand {
	var hands []hand
	for _, line := range data {
		var curr_hand hand
		values := strings.Fields(line)
		bid, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}
		curr_hand.bid = bid
		curr_hand.cards = getCards2(values[0])
		curr_hand.mtype = getType2(curr_hand.cards)
		hands = append(hands, curr_hand)
	}
	return hands
}

func getCards2(cards string) [5]int {
	var res [5]int
	for i, char := range cards {
		res[i] = getCard2(string(char))
	}
	return res
}

func getCard2(card string) int {
	n, err := strconv.Atoi(card)
	if err != nil {
		switch {
		case card == "T":
			n = 10
		case card == "J":
			n = 1
		case card == "Q":
			n = 12
		case card == "K":
			n = 13
		case card == "A":
			n = 14
		}
	}
	return n
}

func getType2(cards [5]int) int {
	res := 0
	pair := 0
	triple := 0
	var previouscard []int
	for i := 0; i < 5; i++ {
		nbr := 1
		if !slices.Contains(previouscard, cards[i]) {
			for j := i + 1; j < 5; j++ {
				if cards[i] == cards[j] {
					nbr++
				}
			}
			switch {
			case nbr == 5:
				return 6
			case nbr == 4:
				if nbrJ(cards) != 0 {
					return 6
				} else {
					return 5
				}
			case nbr == 3:
				triple++
			case nbr == 2:
				pair++
			default:
			}
			previouscard = append(previouscard, cards[i])
		}
	}
	if pair == 2 {
		if nbrJ(cards) == 2 {
			res = 5
		} else if nbrJ(cards) == 1 {
			res = 4
		} else {
			res = 2
		}
	} else if triple == 1 && pair == 1 {
		if nbrJ(cards) == 2 || nbrJ(cards) == 3 {
			res = 6
		} else {
			res = 4
		}
	} else if triple == 1 {
		if nbrJ(cards) != 0 {
			res = 5
		} else {
			res = 3
		}
	} else if pair == 1 {
		if nbrJ(cards) != 0 {
			res = 3
		} else {
			res = 1
		}
	} else if nbrJ(cards) == 1 && res == 0 {
		res = 1
	}
	return res
}

func nbrJ(cards [5]int) int {
	nbr := 0
	for _, card := range cards {
		if card == 1 {
			nbr++
		}
	}
	return nbr
}
