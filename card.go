package poker

type Card int

var (
	intRanks [13]int
	strRanks = "23456789TJQKA"
	primes   = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41}
)

var (
	charRankToIntRank = map[uint8]int{}
	charSuitToIntSuit = map[uint8]int{
		's': 1, // spades
		'h': 2, // hearts
		'd': 4, // diamonds
		'c': 8, // clubs
	}
	intSuitToCharSuit = "xshxdxxxc"
)

var (
	prettySuits = map[int]string{
		1: "\u2660", // spades
		2: "\u2764", // hearts
		4: "\u2666", // diamonds
		8: "\u2663", // clubs
	}
	prettyReds = [...]int{2, 4}
)

func init() {
	for i := 0; i < 13; i++ {
		intRanks[i] = i
	}

	for i := range strRanks {
		charRankToIntRank[strRanks[i]] = intRanks[i]
	}
}

func NewCard(s string) Card {
	rankInt := charRankToIntRank[s[0]]
	suitInt := charSuitToIntSuit[s[1]]
	rankPrime := primes[rankInt]

	bitRank := 1 << uint32(rankInt) << 16
	suit := suitInt << 12
	rank := rankInt << 8

	return Card(bitRank | suit | rank | rankPrime)
}

func (c Card) String() string {
	return string(strRanks[c.Rank()]) + string(intSuitToCharSuit[c.Suit()])
}

func (c Card) Rank() int {
	return (int(c) >> 8) & 0xF
}

func (c Card) Suit() int {
	return (int(c) >> 12) & 0xF
}

func (c Card) BitRank() int {
	return (int(c) >> 16) & 0x1FFF
}

func (c Card) Prime() int {
	return int(c) & 0x3F
}

func primeProductFromHand(cards []Card) int {
	product := 1

	for _, card := range cards {
		product *= (int(card) & 0xFF)
	}

	return product
}

func primeProductFromRankBits(rankBits int) int {
	product := 1

	for _, i := range intRanks {
		if rankBits&(1<<uint(i)) != 0 {
			product *= primes[i]
		}
	}

	return product
}
