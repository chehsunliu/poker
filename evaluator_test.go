package poker

import (
	"encoding/json"
	"testing"

	"github.com/loganjspears/joker/hand"

	_ "github.com/loganjspears/joker/hand"
	"github.com/stretchr/testify/assert"
)

func TestRankString(t *testing.T) {
	data := map[int32]string{
		398:  "Flush",
		2665: "Two Pair",
		6230: "High Card",
		6529: "High Card",
		6823: "High Card",
		2669: "Two Pair",
		4076: "Pair",
		7196: "High Card",
		7221: "High Card",
		6228: "High Card",
	}

	for rank := range data {
		assert.Equal(t, data[rank], RankString(rank))
	}
}

var data1 = map[int32]string{
	6252: `["As", "Ks", "Jc", "7h", "5d"]`, // high card
	3448: `["As", "Ac", "Jc", "7h", "5d"]`, // pair
	2497: `["As", "Ac", "Jc", "Jd", "5d"]`, // two pair
	1636: `["As", "Ac", "Ad", "Jd", "5d"]`, // three of a kind
	1600: `["As", "Ks", "Qd", "Jh", "Td"]`, // straight
	1542: `["Ts", "7s", "4s", "3s", "2s"]`, // flush
	298:  `["4s", "4c", "4d", "2s", "2h"]`, // full house
	19:   `["As", "Ac", "Ad", "Ah", "5h"]`, // four of a kind
	1:    `["As", "Ks", "Qs", "Js", "Ts"]`, // straight flush
}

var data2 = map[int32]string{
	6252: `["3d", "As", "Ks", "Jc", "7h", "5d"]`, // high card
	3448: `["3d", "As", "Ac", "Jc", "7h", "5d"]`, // pair
	2497: `["3d", "As", "Ac", "Jc", "Jd", "5d"]`, // two pair
	1636: `["3d", "As", "Ac", "Ad", "Jd", "5d"]`, // three of a kind
	1600: `["3d", "As", "Ks", "Qd", "Jh", "Td"]`, // straight
	1542: `["3d", "Ts", "7s", "4s", "3s", "2s"]`, // flush
	298:  `["3d", "4s", "4c", "4d", "2s", "2h"]`, // full house
	19:   `["3d", "As", "Ac", "Ad", "Ah", "5h"]`, // four of a kind
	1:    `["3d", "As", "Ks", "Qs", "Js", "Ts"]`, // straight flush
}

var data3 = map[int32]string{
	6252: `["2d", "3d", "As", "Ks", "Jc", "7h", "5d"]`, // high card
	3448: `["2d", "3d", "As", "Ac", "Jc", "7h", "5d"]`, // pair
	2497: `["2d", "3d", "As", "Ac", "Jc", "Jd", "5d"]`, // two pair
	1636: `["2c", "3d", "As", "Ac", "Ad", "Jd", "5d"]`, // three of a kind
	1600: `["2d", "3d", "As", "Ks", "Qd", "Jh", "Td"]`, // straight
	1542: `["2d", "3d", "Ts", "7s", "4s", "3s", "2s"]`, // flush
	298:  `["2d", "3d", "4s", "4c", "4d", "2s", "2h"]`, // full house
	19:   `["2d", "3d", "As", "Ac", "Ad", "Ah", "5h"]`, // four of a kind
	1:    `["2d", "3d", "As", "Ks", "Qs", "Js", "Ts"]`, // straight flush
}

var dataJoker1 = []string{
	`["A♠", "K♠", "J♣", "7♥", "5♦"]`, // high card
	`["A♠", "A♣", "J♣", "7♥", "5♦"]`, // pair
	`["A♠", "A♣", "J♣", "J♦", "5♦"]`, // two pair
	`["A♠", "A♣", "A♦", "J♦", "5♦"]`, // three of a kind
	`["A♠", "K♠", "Q♦", "J♥", "T♦"]`, // straight
	`["T♠", "7♠", "4♠", "3♠", "2♠"]`, // flush
	`["4♠", "4♣", "4♦", "2♠", "2♥"]`, // full house
	`["A♠", "A♣", "A♦", "A♥", "5♥"]`, // four of a kind
	`["A♠", "K♠", "Q♠", "J♠", "T♠"]`, // straight flush
}

var dataJoker2 = []string{
	`["3♦", "A♠", "K♠", "J♣", "7♥", "5♦"]`, // high card
	`["3♦", "A♠", "A♣", "J♣", "7♥", "5♦"]`, // pair
	`["3♦", "A♠", "A♣", "J♣", "J♦", "5♦"]`, // two pair
	`["3♦", "A♠", "A♣", "A♦", "J♦", "5♦"]`, // three of a kind
	`["3♦", "A♠", "K♠", "Q♦", "J♥", "T♦"]`, // straight
	`["3♦", "T♠", "7♠", "4♠", "3♠", "2♠"]`, // flush
	`["3♦", "4♠", "4♣", "4♦", "2♠", "2♥"]`, // full house
	`["3♦", "A♠", "A♣", "A♦", "A♥", "5♥"]`, // four of a kind
	`["3♦", "A♠", "K♠", "Q♠", "J♠", "T♠"]`, // straight flush
}

var dataJoker3 = []string{
	`["2♦", "3♦", "A♠", "K♠", "J♣", "7♥", "5♦"]`, // high card
	`["2♦", "3♦", "A♠", "A♣", "J♣", "7♥", "5♦"]`, // pair
	`["2♦", "3♦", "A♠", "A♣", "J♣", "J♦", "5♦"]`, // two pair
	`["2♣", "3♦", "A♠", "A♣", "A♦", "J♦", "5♦"]`, // three of a kind
	`["2♦", "3♦", "A♠", "K♠", "Q♦", "J♥", "T♦"]`, // straight
	`["2♦", "3♦", "T♠", "7♠", "4♠", "3♠", "2♠"]`, // flush
	`["2♦", "3♦", "4♠", "4♣", "4♦", "2♠", "2♥"]`, // full house
	`["2♦", "3♦", "A♠", "A♣", "A♦", "A♥", "5♥"]`, // four of a kind
	`["2♦", "3♦", "A♠", "K♠", "Q♠", "J♠", "T♠"]`, // straight flush
}

func TestFive(t *testing.T) {
	for score := range data1 {
		var cards []Card

		err := json.Unmarshal([]byte(data1[score]), &cards)
		assert.NoError(t, err)
		assert.Equal(t, score, Evaluate(cards))
	}
}

func TestSix(t *testing.T) {
	for score := range data2 {
		var cards []Card

		err := json.Unmarshal([]byte(data2[score]), &cards)
		assert.NoError(t, err)
		assert.Equal(t, score, Evaluate(cards))
	}
}

func TestSeven(t *testing.T) {
	for score := range data3 {
		var cards []Card

		err := json.Unmarshal([]byte(data3[score]), &cards)
		assert.NoError(t, err)
		assert.Equal(t, score, Evaluate(cards))
	}
}

func BenchmarkFivePoker(b *testing.B) {
	var allCards [][]Card

	for score := range data1 {
		var cards []Card

		json.Unmarshal([]byte(data1[score]), &cards)
		allCards = append(allCards, cards)
	}

	for i := 0; i < b.N; i++ {
		for _, cards := range allCards {
			Evaluate(cards)
		}
	}
}

func BenchmarkFiveJoker(b *testing.B) {
	var allCards [][]hand.Card

	for _, s := range dataJoker1 {
		var cards []hand.Card
		json.Unmarshal([]byte(s), &cards)
		allCards = append(allCards, cards)
	}

	for i := 0; i < b.N; i++ {
		for _, cards := range allCards {
			hand.New(cards)
		}
	}
}

func BenchmarkSixPoker(b *testing.B) {
	var allCards [][]Card

	for score := range data2 {
		var cards []Card

		json.Unmarshal([]byte(data2[score]), &cards)
		allCards = append(allCards, cards)
	}

	for i := 0; i < b.N; i++ {
		for _, cards := range allCards {
			Evaluate(cards)
		}
	}
}

func BenchmarkSixJoker(b *testing.B) {
	var allCards [][]hand.Card

	for _, s := range dataJoker2 {
		var cards []hand.Card
		json.Unmarshal([]byte(s), &cards)
		allCards = append(allCards, cards)
	}

	for i := 0; i < b.N; i++ {
		for _, cards := range allCards {
			hand.New(cards)
		}
	}
}

func BenchmarkSevenPoker(b *testing.B) {
	var allCards [][]Card

	for score := range data3 {
		var cards []Card

		json.Unmarshal([]byte(data3[score]), &cards)
		allCards = append(allCards, cards)
	}

	for i := 0; i < b.N; i++ {
		for _, cards := range allCards {
			Evaluate(cards)
		}
	}
}

func BenchmarkSevenJoker(b *testing.B) {
	var allCards [][]hand.Card

	for _, s := range dataJoker3 {
		var cards []hand.Card
		json.Unmarshal([]byte(s), &cards)
		allCards = append(allCards, cards)
	}

	for i := 0; i < b.N; i++ {
		for _, cards := range allCards {
			hand.New(cards)
		}
	}
}
