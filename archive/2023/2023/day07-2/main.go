package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/merijnf/advent-of-code-go/utilities"
)

type handKind int

const (
	highCard     handKind = iota
	onePair      handKind = iota
	twoPair      handKind = iota
	threeOfAKind handKind = iota
	fullHouse    handKind = iota
	fourOfAKind  handKind = iota
	fiveOfAKind  handKind = iota
)

type hand struct {
	cards []rune
	bid   int
	kind  handKind
}

func main() {
	lines := utilities.LoadInputStringLines(2023, 7)
	hands := make([]hand, len(lines))
	for i, line := range lines {
		hands[i] = parseHands(line)
	}
	sort.Slice(hands, func(i, j int) bool {
		hand1 := hands[i]
		hand2 := hands[j]
		if hand1.kind == hand2.kind {
			for k := 0; k < 5; k++ {
				card1 := determineCardValue(hand1.cards[k])
				card2 := determineCardValue(hand2.cards[k])
				if card1 < card2 {
					return true
				} else if card1 > card2 {
					return false
				}
			}
		}
		return hand1.kind < hand2.kind
	})
	sum := 0
	for index, hand := range hands {
		fmt.Println(string(hand.cards) + " bid: " + strconv.Itoa(hand.bid) + " kind: " + strconv.Itoa(int(hand.kind)))
		sum = sum + ((index + 1) * hand.bid)
	}
	fmt.Println(sum)
}

func parseHands(line string) hand {
	parts := strings.Fields(line)
	bid, error := strconv.Atoi(parts[1])
	utilities.CheckError(error)
	cards := []rune(parts[0])
	return hand{cards: cards, bid: bid, kind: determineHandKind(cards)}
}

func determineCardValue(card rune) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 1
	case 'T':
		return 10
	case '9':
		return 9
	case '8':
		return 8
	case '7':
		return 7
	case '6':
		return 6
	case '5':
		return 5
	case '4':
		return 4
	case '3':
		return 3
	case '2':
		return 2
	}
	return 0
}

func determineHandKind(cardsIn []rune) handKind {
	cards := make([]rune, len(cardsIn))
	copy(cards, cardsIn)
	sort.Slice(cards, func(i, j int) bool {
		return cards[i] < cards[j]
	})

	jokers := 0

	prev := cards[0]
	count := 0
	counts := make([]int, 0)
	for _, card := range cards {
		if prev == card {
			count++
		} else {
			if prev == 'J' {
				jokers = count
			} else {
				counts = append(counts, count)
			}
			count = 1
		}
		prev = card

	}
	if prev == 'J' {
		jokers = count
	} else {
		counts = append(counts, count)
	}
	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})
	if jokers == 5 {
		return fiveOfAKind
	}
	counts[0] = counts[0] + jokers
	if reflect.DeepEqual(counts, []int{5}) {
		return fiveOfAKind
	} else if reflect.DeepEqual(counts, []int{4, 1}) {
		return fourOfAKind
	} else if reflect.DeepEqual(counts, []int{3, 2}) {
		return fullHouse
	} else if reflect.DeepEqual(counts, []int{3, 1, 1}) {
		return threeOfAKind
	} else if reflect.DeepEqual(counts, []int{2, 2, 1}) {
		return twoPair
	} else if reflect.DeepEqual(counts, []int{2, 1, 1, 1}) {
		return onePair
	} else {
		return highCard
	}
}
