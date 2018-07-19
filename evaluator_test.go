package poker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRankString(t *testing.T) {
	data := map[int]string{
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

func TestFive(t *testing.T) {
	data := map[int][]Card{
		398:  {268446761, 134228773, 33564957, 139523, 533255},
		2665: {134228773, 134224677, 2131213, 2102541, 69634},
		6230: {33589533, 16812055, 268454953, 134253349, 8394515},
		6529: {557831, 4228625, 268442665, 33589533, 139523},
		6823: {16787479, 135427, 134224677, 33560861, 295429},
		2669: {16787479, 134224677, 1057803, 134228773, 1082379},
		4076: {134253349, 33573149, 33589533, 1053707, 529159},
		7196: {1065995, 67144223, 266757, 73730, 4204049},
		7221: {81922, 4212241, 33589533, 8406803, 16795671},
		6228: {268471337, 67127839, 98306, 134253349, 295429},
	}

	for score := range data {
		assert.Equal(t, score, Evaluate(data[score]))
	}
}

func TestSix(t *testing.T) {
	data := map[int][]Card{
		2559: {557831, 268454953, 8398611, 529159, 16783383, 268446761},
		6359: {164099, 67115551, 268454953, 33564957, 8423187, 2114829},
		4261: {134224677, 1053707, 16783383, 67127839, 16812055, 33560861},
		3548: {8423187, 134253349, 134228773, 2131213, 67115551, 268442665},
		6689: {541447, 8394515, 67119647, 295429, 134253349, 33560861},
		4592: {8406803, 69634, 4199953, 533255, 8394515, 16795671},
		6218: {268471337, 67115551, 147715, 134253349, 69634, 2102541},
		6366: {268454953, 557831, 67127839, 4199953, 33560861, 1065995},
		5488: {16787479, 73730, 529159, 541447, 295429, 147715},
		2782: {1065995, 557831, 67119647, 1053707, 2106637, 67115551},
	}

	for score := range data {
		assert.Equal(t, score, Evaluate(data[score]))
	}
}

func TestSeven(t *testing.T) {
	data := map[int][]Card{
		5198: {139523, 295429, 67144223, 1065995, 1082379, 73730, 16783383},
		2652: {4204049, 279045, 135427, 134236965, 134228773, 4199953, 270853},
		5854: {139523, 16787479, 8394515, 295429, 67127839, 147715, 4204049},
		4528: {533255, 8423187, 139523, 67119647, 2114829, 33589533, 8394515},
		4187: {33564957, 147715, 81922, 270853, 1053707, 33560861, 2131213},
		2350: {98306, 8394515, 279045, 139523, 135427, 147715, 134236965},
		3022: {4228625, 81922, 16787479, 139523, 8406803, 4212241, 8398611},
		3205: {279045, 69634, 268446761, 73730, 2102541, 2114829, 33564957},
		5900: {533255, 164099, 98306, 2102541, 139523, 33573149, 1082379},
		3153: {4228625, 73730, 33573149, 81922, 4212241, 135427, 557831},
	}

	for score := range data {
		assert.Equal(t, score, Evaluate(data[score]))
	}
}
