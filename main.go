package main

func main() {
	card := getNewDeck()
	hand, remainingCards := deal(card, 5)

	hand.print()
	remainingCards.print()
}
