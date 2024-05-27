package main

import "fmt"

type Deck []string
type CardSuits []string
type CardValues []string

func getNewDeck() Deck {
	cards := Deck{}
	cardSuits := CardSuits{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := CardValues{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d Deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}
