package main

func main() {
	cards := newDeck()
	cards.saveToFile("cards.txt")
	cards = newDeckFromFile("cards.txt")
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
