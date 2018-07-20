# Poker

[![GoDoc](https://godoc.org/github.com/chehsunliu/poker?status.svg)](https://godoc.org/github.com/chehsunliu/poker)

This package is ported from the Python library [worldveil/deuces](https://github.com/worldveil/deuces).

## Installation

Use `go get` to install poker:

```sh
$ go get github.com/chehsunliu/poker
```

## Example

Evaluate five cards:

```go
package main

import (
	"fmt"

	"github.com/chehsunliu/poker"
)

func main() {
	board := []poker.Card{
		poker.NewCard("Ah"),
		poker.NewCard("Kd"),
		poker.NewCard("Jc"),
	}
	hand := []poker.Card{
		poker.NewCard("Qs"),
		poker.NewCard("Th"),
	}
	fmt.Println(board, hand)

	rank := poker.Evaluate(append(board, hand...))
	fmt.Println(rank)
	fmt.Println(poker.RankString(rank))
}
```

```sh
$ go run ./main.go
[Ah Kd Jc] [Qs Th]
1600
Straight
```

Evaluate seven cards:

```go
package main

import (
	"fmt"

	"github.com/chehsunliu/poker"
)

func main() {
	deck := poker.NewDeck()
	hand := deck.Draw(7)
	fmt.Println(hand)

	rank := poker.Evaluate(hand)
	fmt.Println(rank)
	fmt.Println(poker.RankString(rank))
}
```

```sh
$ go run ./main.go
[4c Qh Ad 9c 9s 3h 4d]
3062
Two Pair
```