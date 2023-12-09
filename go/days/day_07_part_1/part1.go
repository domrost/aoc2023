package day_07_part_1

import (
	"cmp"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// Part1 solves the first part of the exercise
func Run(input []string) string {

	handTypes := typeSet{fiveOfAKind: []hand{}, fourOfAKind: []hand{}, fullHouse: []hand{}, threeOfAKind: []hand{}, twoPair: []hand{}, onePair: []hand{}, highCard: []hand{}}

	for _, line := range input {
		hand := parseToHand(line)

		typify(hand, &handTypes)
	}

	compFunc := func(a, b hand) int {
		for i := 0; i < len(a.cards); i++ {
			res := cmp.Compare(a.cards[i], b.cards[i])
			if res != 0 {
				return res
			}
		}
		return 0
	}

	slices.SortFunc(handTypes.highCard, compFunc)
	slices.SortFunc(handTypes.onePair, compFunc)
	slices.SortFunc(handTypes.twoPair, compFunc)
	slices.SortFunc(handTypes.threeOfAKind, compFunc)
	slices.SortFunc(handTypes.fullHouse, compFunc)
	slices.SortFunc(handTypes.fourOfAKind, compFunc)
	slices.SortFunc(handTypes.fiveOfAKind, compFunc)

	ranking := []hand{}
	ranking = append(ranking, handTypes.highCard...)
	ranking = append(ranking, handTypes.onePair...)
	ranking = append(ranking, handTypes.twoPair...)
	ranking = append(ranking, handTypes.threeOfAKind...)
	ranking = append(ranking, handTypes.fullHouse...)
	ranking = append(ranking, handTypes.fourOfAKind...)
	ranking = append(ranking, handTypes.fiveOfAKind...)

	winnings := 0
	for i := 0; i < len(ranking); i++ {
		winnings += ranking[i].bid * (i + 1)
	}
	return fmt.Sprint(winnings)
}

func typify(hand hand, typeSet *typeSet) {
	cards := slices.Clone(hand.cards)
	slices.Sort(cards)

	sortedString := ""
	for _, card := range cards {
		sortedString += card.string()
	}

	// Because Go does not implement back references, the checks become super ugly

	// Check five of a kind
	re := regexp.MustCompile(`2{5,5}|3{5,5}|4{5,5}|5{5,5}|6{5,5}|7{5,5}|8{5,5}|9{5,5}|T{5,5}|J{5,5}|Q{5,5}|K{5,5}|A{5,5}`)
	if re.Match([]byte(sortedString)) {
		typeSet.fiveOfAKind = append(typeSet.fiveOfAKind, hand)
		return
	}

	// Check four of a kind
	re = regexp.MustCompile(`2{4,4}|3{4,4}|4{4,4}|5{4,4}|6{4,4}|7{4,4}|8{4,4}|9{4,4}|T{4,4}|J{4,4}|Q{4,4}|K{4,4}|A{4,4}`)
	if re.Match([]byte(sortedString)) {
		typeSet.fourOfAKind = append(typeSet.fourOfAKind, hand)
		return
	}

	// Check full house
	re = regexp.MustCompile(`((2{3,3}|3{3,3}|4{3,3}|5{3,3}|6{3,3}|7{3,3}|8{3,3}|9{3,3}|T{3,3}|J{3,3}|Q{3,3}|K{3,3}|A{3,3})(2{2,2}|3{2,2}|4{2,2}|5{2,2}|6{2,2}|7{2,2}|8{2,2}|9{2,2}|T{2,2}|J{2,2}|Q{2,2}|K{2,2}|A{2,2}))|((2{2,2}|3{2,2}|4{2,2}|5{2,2}|6{2,2}|7{2,2}|8{2,2}|9{2,2}|T{2,2}|J{2,2}|Q{2,2}|K{2,2}|A{2,2})(2{3,3}|3{3,3}|4{3,3}|5{3,3}|6{3,3}|7{3,3}|8{3,3}|9{3,3}|T{3,3}|J{3,3}|Q{3,3}|K{3,3}|A{3,3}))`)
	if re.Match([]byte(sortedString)) {
		typeSet.fullHouse = append(typeSet.fullHouse, hand)
		return
	}

	// Check three of a kind
	re = regexp.MustCompile(`2{3,3}|3{3,3}|4{3,3}|5{3,3}|6{3,3}|7{3,3}|8{3,3}|9{3,3}|T{3,3}|J{3,3}|Q{3,3}|K{3,3}|A{3,3}`)
	if re.Match([]byte(sortedString)) {
		typeSet.threeOfAKind = append(typeSet.threeOfAKind, hand)
		return
	}

	// Check two pair
	re = regexp.MustCompile(`.?(2{2,2}|3{2,2}|4{2,2}|5{2,2}|6{2,2}|7{2,2}|8{2,2}|9{2,2}|T{2,2}|J{2,2}|Q{2,2}|K{2,2}|A{2,2}).?(2{2,2}|3{2,2}|4{2,2}|5{2,2}|6{2,2}|7{2,2}|8{2,2}|9{2,2}|T{2,2}|J{2,2}|Q{2,2}|K{2,2}|A{2,2})`)
	if re.Match([]byte(sortedString)) {
		typeSet.twoPair = append(typeSet.twoPair, hand)
		return
	}

	// Check one pair
	re = regexp.MustCompile(`2{2,2}|3{2,2}|4{2,2}|5{2,2}|6{2,2}|7{2,2}|8{2,2}|9{2,2}|T{2,2}|J{2,2}|Q{2,2}|K{2,2}|A{2,2}`)
	if re.Match([]byte(sortedString)) {
		typeSet.onePair = append(typeSet.onePair, hand)
		return
	}

	typeSet.highCard = append(typeSet.highCard, hand)
}

func parseToHand(line string) hand {
	handString, bidString, _ := strings.Cut(line, " ")
	bid, _ := strconv.Atoi(bidString)
	hand := hand{cards: []card{}, bid: bid}

	for _, cardString := range strings.Split(handString, "") {
		card, _ := getCardFromString(cardString)
		hand.cards = append(hand.cards, card)
	}
	return hand
}
