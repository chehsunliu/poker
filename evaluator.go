package poker

import (
	"fmt"
)

var table *LookupTable

func init() {
	table = NewLookupTable()
}

func RankClass(rank int) int {
	targets := [...]int{
		maxStraightFlush,
		maxFourOfAKind,
		maxFullHouse,
		maxFlush,
		maxStraight,
		maxThreeOfAKind,
		maxTwoPair,
		maxPair,
		maxHighCard,
	}

	if rank < 0 {
		panic(fmt.Sprintf("rank %d is less than zero", rank))
	}

	for _, target := range targets {
		if rank <= target {
			return maxToRankClass[target]
		}
	}

	panic(fmt.Sprintf("rank %d is unknown", rank))
}

func RankString(rank int) string {
	return rankClassToString[RankClass(rank)]
}

func Evaluate(cards []Card) int {
	switch len(cards) {
	case 5:
		return five(cards...)
	case 6:
		return six(cards...)
	case 7:
		return seven(cards...)
	default:
		panic("Only support 5, 6 and 7 cards.")
	}
}

func five(cards ...Card) int {
	if cards[0]&cards[1]&cards[2]&cards[3]&cards[4]&0xF000 != 0 {
		handOR := (cards[0] | cards[1] | cards[2] | cards[3] | cards[4]) >> 16
		prime := primeProductFromRankBits(int(handOR))
		return table.FlushLookup[prime]
	}

	prime := primeProductFromHand(cards)
	return table.UnsuitedLookup[prime]
}

func six(cards ...Card) int {
	minimum := maxHighCard
	targets := make([]Card, len(cards))

	for i := 0; i < len(cards); i++ {
		copy(targets, cards)
		targets := append(targets[:i], targets[i+1:]...)

		score := five(targets...)
		if score < minimum {
			minimum = score
		}
	}

	return minimum
}

func seven(cards ...Card) int {
	minimum := maxHighCard
	targets := make([]Card, len(cards))

	for i := 0; i < len(cards); i++ {
		copy(targets, cards)
		targets := append(targets[:i], targets[i+1:]...)

		score := six(targets...)
		if score < minimum {
			minimum = score
		}
	}

	return minimum
}