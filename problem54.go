// Poker hands

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var p = fmt.Println

type Card struct {
	Rank, Suit rune
	Value      int
}

func newCard(token string) *Card {
	valuesMap := map[rune]int{
		'2': 1,
		'3': 2,
		'4': 3,
		'5': 4,
		'6': 5,
		'7': 6,
		'8': 7,
		'9': 8,
		'T': 9,
		'J': 10,
		'Q': 11,
		'K': 12,
		'A': 13,
	}
	tokens := []rune(token)
	return &Card{tokens[0], tokens[1], valuesMap[tokens[0]]}
}

func (card Card) String() string {
	return fmt.Sprintf("%s%s", string(card.Rank), string(card.Suit))
}

// returns -1, 0 or 1: the compare result among card and card2
func (card *Card) cmp(card2 *Card) (result int) {
	result = card.Value - card2.Value
	if result > 0 {
		result = 1
	} else if result < 0 {
		result = -1
	}
	return
}

func (hand Hand) String() string {
	return fmt.Sprintf(
		"Cards: [%s %s %s %s %s] Value: %d",
		hand.Cards[0].String(),
		hand.Cards[1].String(),
		hand.Cards[2].String(),
		hand.Cards[3].String(),
		hand.Cards[4].String(),
		hand.Value,
	)
}

type Hand struct {
	Cards   [5]*Card
	Counter map[int]int
	Value   int
}

// given a string with the 5 cards, it will return a sorted array of cards
func parseHand(cards string) *Hand {
	hand := Hand{}
	newCards := make([]*Card, 5)
	for i, v := range strings.Split(cards, " ") {
		newCards[i] = newCard(v)
	}
	sort.Slice(newCards, func(i, j int) bool {
		return newCards[i].Value > newCards[j].Value
	})
	// the [:] forces the copy from a slice to an array
	copy(hand.Cards[:], newCards)
	// map the card by their values and qty,
	// ex: JC JD QC QD TD = 10: 2, 11: 2, 9: 1

	hand.setCounter()
	hand.setValue()

	return &hand
}

func (hand *Hand) setCounter() {
	hand.Counter = make(map[int]int)
	for _, card := range hand.Cards {
		_, found := hand.Counter[card.Value]
		if found {
			hand.Counter[card.Value]++
		} else {
			hand.Counter[card.Value] = 1
		}
	}
}

func parsePlayersHands(line string) []*Hand {
	hands := make([]*Hand, 2)
	hands[0] = parseHand(line[0:14])
	hands[1] = parseHand(line[15:29])
	return hands
}

func (hand *Hand) isFlush() bool {
	suit := hand.Cards[0].Suit
	for _, card := range hand.Cards {
		if card.Suit != suit {
			return false
		}
	}
	return true
}

func (hand *Hand) isStraight() bool {
	// Cards are sorted by their values in DESC order
	initialValue := hand.Cards[0].Value
	for index, card := range hand.Cards {
		expectedValue := initialValue - index
		if expectedValue != card.Value {
			return false
		}
	}
	return true
}

func (hand *Hand) isFour() bool {
	for _, v := range hand.Counter {
		if v == 4 {
			return true
		}
	}
	return false
}

func (hand *Hand) isFullHouse() bool {
	return hand.isThreeOfAKind() && hand.isPair()
}

func (hand *Hand) isThreeOfAKind() bool {
	for _, v := range hand.Counter {
		if v == 3 {
			return true
		}
	}
	return false
}

func (hand *Hand) isTwoPairs() bool {
	pairCounter := 0
	for _, v := range hand.Counter {
		if v == 2 {
			pairCounter++
		}
	}
	return pairCounter == 2
}

func (hand *Hand) isPair() bool {
	for _, v := range hand.Counter {
		if v == 2 {
			return true
		}
	}
	return false
}

func (hand *Hand) isRoyalStraightFlush() bool {
	startWithAce := hand.Cards[0].Value == 13
	return startWithAce && hand.isFlush() && hand.isStraight()
}

func (hand *Hand) setValue() {
	result := 0
	if hand.isRoyalStraightFlush() {
		result = 9
	} else if hand.isStraight() && hand.isFlush() {
		result = 8
	} else if hand.isFour() {
		result = 7
	} else if hand.isFullHouse() {
		result = 6
	} else if hand.isFlush() {
		result = 5
	} else if hand.isStraight() {
		result = 4
	} else if hand.isThreeOfAKind() {
		result = 3
	} else if hand.isTwoPairs() {
		result = 2
	} else if hand.isPair() {
		result = 1
	}
	hand.Value = result
}

func highCardIndex(hands []*Hand) int {
	for i := 0; i < 5; i++ {
		if hands[0].Cards[i].Value > hands[1].Cards[i].Value {
			return 0
		} else if hands[1].Cards[i].Value > hands[0].Cards[i].Value {
			return 1
		}
	}
	return -1 // a draw
}

func pairIndex(hands []*Hand) int {
	player1PairRank, player2PairRank := -1, -1
	for rank, count := range hands[0].Counter {
		if count == 2 {
			player1PairRank = rank
		}
	}
	for rank, count := range hands[1].Counter {
		if count == 2 {
			player2PairRank = rank
		}
	}
	if player1PairRank > player2PairRank {
		return 0
	} else if player2PairRank > player1PairRank {
		return 1
	}
	return highCardIndex(hands)
}

func twoPairsIndex(hands []*Hand) int {
	player1PairRanks := []int{}
	player2PairRanks := []int{}
	for rank, count := range hands[0].Counter {
		if count == 2 {
			player1PairRanks = append(player1PairRanks, rank)
		}
	}
	for rank, count := range hands[1].Counter {
		if count == 2 {
			player2PairRanks = append(player2PairRanks, rank)
		}
	}
	sort.Slice(player1PairRanks, func(i, j int) bool {
		return player1PairRanks[i] > player1PairRanks[j]
	})
	sort.Slice(player2PairRanks, func(i, j int) bool {
		return player2PairRanks[i] > player2PairRanks[j]
	})
	for i := 0; i < 2; i++ {
		if player1PairRanks[i] > player2PairRanks[i] {
			return i
		}
	}
	return highCardIndex(hands)
}

func decideWinner(hands []*Hand) int {
	if hands[0].Value == 0 {
		return highCardIndex(hands)
	} else if hands[0].Value == 1 {
		return pairIndex(hands)
	} else if hands[0].Value == 2 {
		return twoPairsIndex(hands) // NOT REACHED IN GAMES FILE
	}
	// The only hand draws present in the games file are solved with highcards
	// or single pairs; the other draw deciders were not necessary
	return -1
}

func readFile() [][]*Hand {
	hands := make([][]*Hand, 0)
	file, _ := os.Open("p054_poker.txt")
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		currentGame := parsePlayersHands(line)
		hands = append(hands, currentGame)
	}
	return hands
}

func main() {
	p("Problem 54")
	result := readFile()
	total := 0
	for _, game := range result {
		player1 := game[0]
		player2 := game[1]
		if player1.Value > player2.Value {
			total++
		} else if player2.Value == player1.Value {
			winnerIndex := decideWinner(game)
			total += (winnerIndex ^ 1)
			// p("player", (winnerIndex + 1), "wins")
			if winnerIndex == -1 {
				panic("DRAW GAME")
			}
		}
	}
	p(total)
}
