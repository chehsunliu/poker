package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRankIntegers(t *testing.T) {
	for i := 0; i < 13; i++ {
		assert.Equal(t, int32(i), intRanks[i])
	}
}

func TestNewCard(t *testing.T) {
	assert.Equal(t, Card(268446761), NewCard("Ah"))
	assert.Equal(t, Card(134224677), NewCard("Ks"))
}

func TestString(t *testing.T) {
	assert.Equal(t, "3s", NewCard("3s").String())
}

func TestBitRank(t *testing.T) {
	assert.Equal(t, int32(2048), NewCard("Ks").BitRank())
}

func TestPrimeProductFromHand(t *testing.T) {
	assert.Equal(t, int32(20387), primeProductFromHand([]Card{8398611, 134236965, 33564957}))
	assert.Equal(t, int32(61161), primeProductFromHand([]Card{8398611, 134236965, 33564957, 135427}))
	assert.Equal(t, int32(183483), primeProductFromHand([]Card{8398611, 134236965, 33564957, 135427, 139523}))
}

func TestPrimeProductFromRankBits(t *testing.T) {
	assert.Equal(t, int32(2), primeProductFromRankBits(1))
	assert.Equal(t, int32(42), primeProductFromRankBits(11))
	assert.Equal(t, int32(110), primeProductFromRankBits(21))
	assert.Equal(t, int32(2310), primeProductFromRankBits(31))
	assert.Equal(t, int32(4290), primeProductFromRankBits(55))
	assert.Equal(t, int32(1785), primeProductFromRankBits(78))
	assert.Equal(t, int32(1326), primeProductFromRankBits(99))
	assert.Equal(t, int32(34034), primeProductFromRankBits(121))
	assert.Equal(t, int32(30107), primeProductFromRankBits(344))
}
