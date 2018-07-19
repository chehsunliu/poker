package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexographicallyNextBitSequence(t *testing.T) {
	var n1 int32 = 31
	n2 := lexographicallyNextBitSequence(n1)
	n3 := lexographicallyNextBitSequence(n2)
	n4 := lexographicallyNextBitSequence(n3)
	n5 := lexographicallyNextBitSequence(n4)
	n6 := lexographicallyNextBitSequence(n5)

	assert.Equal(t, 47, int(n2))
	assert.Equal(t, 55, int(n3))
	assert.Equal(t, 59, int(n4))
	assert.Equal(t, 61, int(n5))
	assert.Equal(t, 62, int(n6))
}

func TestFlushLookup(t *testing.T) {
	table := newLookupTable()

	assert.Len(t, table.flushLookup, 1287)
	data1 := map[int32]int32{
		1551891: 359,
		1469973: 435,
		128535:  775,
		3958297: 832,
		260015:  1487,
		29274:   800,
		239134:  1189,
		1357345: 954,
		601634:  628,
		4003363: 389,
		12558:   1541,
		1066533: 623,
		2750041: 1147,
		732711:  426,
		139035:  1271,
		650793:  469,
		20010:   1408,
		28490:   1141,
		46002:   797,
		3825245: 820,
		3429937: 433,
		282162:  366,
		159285:  483,
		118326:  476,
		6290999: 388,
	}
	for key := range data1 {
		assert.Equal(t, data1[key], table.flushLookup[key])
	}

	assert.Len(t, table.unsuitedLookup, 6175)
	data2 := map[int32]int32{
		1514071:  3097,
		4347:     2378,
		4350:     5689,
		725249:   3051,
		2167055:  3812,
		102675:   2697,
		4375:     140,
		299299:   5041,
		1544491:  2086,
		2068781:  6903,
		135470:   7110,
		6230319:  1685,
		69938:    3116,
		168245:   7366,
		4408:     2433,
		102718:   7264,
		5411139:  3343,
		4420:     6168,
		790855:   1927,
		1085773:  89,
		29692241: 2622,
		230115:   4367,
		11662:    2259,
		233818:   4395,
		4446:     5939,
	}
	for key := range data2 {
		assert.Equal(t, data2[key], table.unsuitedLookup[key])
	}
}
