package day07_part1

import "errors"

type card int

const (
	card2 card = iota
	card3
	card4
	card5
	card6
	card7
	card8
	card9
	cardT
	cardJ
	cardQ
	cardK
	cardA
)

func (c card) string() string {
	switch c {
	case card2:
		return "2"
	case card3:
		return "3"
	case card4:
		return "4"
	case card5:
		return "5"
	case card6:
		return "6"
	case card7:
		return "7"
	case card8:
		return "8"
	case card9:
		return "9"
	case cardT:
		return "T"
	case cardJ:
		return "J"
	case cardQ:
		return "Q"
	case cardK:
		return "K"
	case cardA:
		return "A"
	}
	return ""
}

func getCardFromString(s string) (card, error) {
	switch s {
	case "2":
		return card2, nil
	case "3":
		return card3, nil
	case "4":
		return card4, nil
	case "5":
		return card5, nil
	case "6":
		return card6, nil
	case "7":
		return card7, nil
	case "8":
		return card8, nil
	case "9":
		return card9, nil
	case "T":
		return cardT, nil
	case "J":
		return cardJ, nil
	case "Q":
		return cardQ, nil
	case "K":
		return cardK, nil
	case "A":
		return cardA, nil

	}
	return -1, errors.New("invalid card string")
}

type hand struct {
	cards []card
	bid   int
}

type typeSet struct {
	fiveOfAKind  []hand
	fourOfAKind  []hand
	fullHouse    []hand
	threeOfAKind []hand
	twoPair      []hand
	onePair      []hand
	highCard     []hand
}
