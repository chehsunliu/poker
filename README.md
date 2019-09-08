# Poker

[![CircleCI](https://circleci.com/gh/chehsunliu/poker/tree/master.svg?style=shield&circle-token=abebd63b852ce8ecdcdf3f7e597be743d07402e4)](https://circleci.com/gh/chehsunliu/poker/tree/master) [![GoDoc](https://godoc.org/github.com/chehsunliu/poker?status.svg)](https://godoc.org/github.com/chehsunliu/poker) [![codecov](https://codecov.io/gh/chehsunliu/poker/branch/master/graph/badge.svg)](https://codecov.io/gh/chehsunliu/poker)

Poker is ported from the Python library [worldveil/deuces](https://github.com/worldveil/deuces).

## Installation

Use `go get` to install Poker:

```sh
$ go get github.com/chehsunliu/poker
```

## Usage

Support 5-, 6-, and 7-card evalutions:

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
[Kd 4h Qh 3s 8s 5h Jd]
6695
High Card

$ go run ./main.go
[4c Qh Ad 9c 9s 3h 4d]
3062
Two Pair

$ go run ./main.go
[Jh Qd Kd Qs 7d As Qh]
1742
Three of a Kind
```

## Performance

Compared with [notnil/joker](https://github.com/notnil/joker), Poker is 160x faster on 5-card evaluation, and drops to 40x faster on 7-card evaluation.

```sh
go test -bench=. -benchtime 5s
goos: darwin
goarch: amd64
pkg: github.com/chehsunliu/poker
BenchmarkFivePoker-4    	23396181	       253 ns/op
BenchmarkFiveJoker-4    	  141036	     41662 ns/op
BenchmarkSixPoker-4     	 3037298	      1949 ns/op
BenchmarkSixJoker-4     	   28158	    211533 ns/op
BenchmarkSevenPoker-4   	  356448	     16357 ns/op
BenchmarkSevenJoker-4   	    7143	    759394 ns/op
PASS
ok  	github.com/chehsunliu/poker	40.111s
```
