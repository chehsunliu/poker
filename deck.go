package poker

type Deck struct {
	cards []Card
}

func NewDeck() *Deck {
	deck := &Deck{}
	return deck
}
