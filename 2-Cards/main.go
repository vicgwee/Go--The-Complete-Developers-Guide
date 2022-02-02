package main

func main() {
	cards := newDeck()
	cards = append(cards, "Six of Spades")
	hand, remainingCards := deal(cards, 20)
	hand.print()
	remainingCards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
