package poker

const (
	maxStraightFlush = 10
	maxFourOfAKind   = 166
	maxFullHouse     = 322
	maxFlush         = 1599
	maxStraight      = 1609
	maxThreeOfAKind  = 2467
	maxTwoPair       = 3325
	maxPair          = 6185
	maxHighCard      = 7462
)

var maxToRankClass = map[int32]int32{
	maxStraightFlush: 1,
	maxFourOfAKind:   2,
	maxFullHouse:     3,
	maxFlush:         4,
	maxStraight:      5,
	maxThreeOfAKind:  6,
	maxTwoPair:       7,
	maxPair:          8,
	maxHighCard:      9,
}

var rankClassToString = map[int32]string{
	1: "Straight Flush",
	2: "Four of a Kind",
	3: "Full House",
	4: "Flush",
	5: "Straight",
	6: "Three of a Kind",
	7: "Two Pair",
	8: "Pair",
	9: "High Card",
}

type lookupTable struct {
	flushLookup    map[int32]int32
	unsuitedLookup map[int32]int32
}

func newLookupTable() *lookupTable {
	table := &lookupTable{}
	table.flushLookup = map[int32]int32{}
	table.unsuitedLookup = map[int32]int32{}

	table.flushes()
	table.multiples()

	return table
}

func (table *lookupTable) flushes() {
	// straight flushes in rank order
	straightFlushes := []int32{
		7936, // 0b1111100000000 royal flush
		3968, // 0b111110000000
		1984, // 0b11111000000
		992,  // 0b1111100000
		496,  // 0b111110000
		248,  // 0b11111000
		124,  // 0b1111100
		62,   // 0b111110
		31,   // 0b11111
		4111, // 0b1000000001111 5 high
	}

	var flushes []int32
	var flush int32 = 31 // 0b11111

	for i := 0; i < 1277+len(straightFlushes)-1; i++ {
		flush = lexographicallyNextBitSequence(flush)

		notSF := true
		for _, sf := range straightFlushes {
			if flush^sf == 0 {
				notSF = false
			}
		}

		if notSF {
			flushes = append(flushes, flush)
		}
	}

	for i, j := 0, len(flushes)-1; i < j; i, j = i+1, j-1 {
		flushes[i], flushes[j] = flushes[j], flushes[i]
	}

	var rank int32 = 1
	for _, sf := range straightFlushes {
		primeProduct := primeProductFromRankBits(sf)
		table.flushLookup[primeProduct] = rank
		rank++
	}

	rank = maxFullHouse + 1
	for _, f := range flushes {
		primeProduct := primeProductFromRankBits(f)
		table.flushLookup[primeProduct] = rank
		rank++
	}

	table.straightAndHighCards(straightFlushes, flushes)
}

func (table *lookupTable) straightAndHighCards(straights, highcards []int32) {
	var rank int32 = maxFlush + 1

	for _, s := range straights {
		primeProduct := primeProductFromRankBits(s)
		table.unsuitedLookup[primeProduct] = rank
		rank++
	}

	rank = maxPair + 1
	for _, h := range highcards {
		primeProduct := primeProductFromRankBits(h)
		table.unsuitedLookup[primeProduct] = rank
		rank++
	}
}

func (table *lookupTable) multiples() {
	backwardRanks := make([]int32, len(intRanks))
	for i := range intRanks {
		backwardRanks[13-i-1] = intRanks[i]
	}

	// 1) Four of a Kind
	var rank int32 = maxStraightFlush + 1

	for _, i := range backwardRanks {
		kickers := make([]int32, len(backwardRanks))
		copy(kickers, backwardRanks)

		for j := 0; j < len(kickers); j++ {
			if kickers[j] == i {
				kickers = append(kickers[:j], kickers[j+1:]...)
				break
			}
		}

		for _, k := range kickers {
			product := primes[i] * primes[i] * primes[i] * primes[i] * primes[k]
			table.unsuitedLookup[product] = rank
			rank++
		}
	}

	// 2) Full House
	rank = maxFourOfAKind + 1

	for _, i := range backwardRanks {
		pairRanks := make([]int32, len(backwardRanks))
		copy(pairRanks, backwardRanks)

		for j := 0; j < len(pairRanks); j++ {
			if pairRanks[j] == i {
				pairRanks = append(pairRanks[:j], pairRanks[j+1:]...)
				break
			}
		}

		for _, pr := range pairRanks {
			product := primes[i] * primes[i] * primes[i] * primes[pr] * primes[pr]
			table.unsuitedLookup[product] = rank
			rank++
		}
	}

	// 3) Three of a Kind
	rank = maxStraight + 1

	for _, i := range backwardRanks {
		kickers := make([]int32, len(backwardRanks))
		copy(kickers, backwardRanks)

		for j := 0; j < len(kickers); j++ {
			if kickers[j] == i {
				kickers = append(kickers[:j], kickers[j+1:]...)
				break
			}
		}

		for j := 0; j < len(kickers)-1; j++ {
			for k := j + 1; k < len(kickers); k++ {
				c1, c2 := kickers[j], kickers[k]
				product := primes[i] * primes[i] * primes[i] * primes[c1] * primes[c2]
				table.unsuitedLookup[product] = rank
				rank++
			}
		}
	}

	// 4) Two Pair
	rank = maxThreeOfAKind + 1

	for i := 0; i < len(backwardRanks)-1; i++ {
		for j := i + 1; j < len(backwardRanks); j++ {
			pair1, pair2 := backwardRanks[i], backwardRanks[j]

			kickers := make([]int32, len(backwardRanks))
			copy(kickers, backwardRanks)

			for k := 0; k < len(kickers); k++ {
				if kickers[k] == pair1 {
					kickers = append(kickers[:k], kickers[k+1:]...)
					break
				}
			}

			for k := 0; k < len(kickers); k++ {
				if kickers[k] == pair2 {
					kickers = append(kickers[:k], kickers[k+1:]...)
					break
				}
			}

			for _, kicker := range kickers {
				product := primes[pair1] * primes[pair1] * primes[pair2] * primes[pair2] * primes[kicker]
				table.unsuitedLookup[product] = rank
				rank++
			}
		}
	}

	// 5) Pair
	rank = maxTwoPair + 1

	for _, pairRank := range backwardRanks {
		kickers := make([]int32, len(backwardRanks))
		copy(kickers, backwardRanks)

		for k := 0; k < len(kickers); k++ {
			if kickers[k] == pairRank {
				kickers = append(kickers[:k], kickers[k+1:]...)
				break
			}
		}

		for i := 0; i < len(kickers)-2; i++ {
			for j := i + 1; j < len(kickers)-1; j++ {
				for k := j + 1; k < len(kickers); k++ {
					k1, k2, k3 := kickers[i], kickers[j], kickers[k]
					product := primes[pairRank] * primes[pairRank] * primes[k1] * primes[k2] * primes[k3]
					table.unsuitedLookup[product] = rank
					rank++
				}
			}
		}
	}
}

// LexographicallyNextBitSequence calculates the next permutation of
// bits in a lexicographical sense. The algorithm comes from
// https://graphics.stanford.edu/~seander/bithacks.html#NextBitPermutation.
func lexographicallyNextBitSequence(bits int32) int32 {
	t := (bits | (bits - 1)) + 1
	return t | ((((t & -t) / (bits & -bits)) >> 1) - 1)
}
