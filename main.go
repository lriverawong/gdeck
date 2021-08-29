package main

import "fmt"

func main() {
	cards := newDeck()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// fmt.Println(cards.toString())

	cards_db := "cards_db.csv"

	cards.saveToFile(cards_db)

	cards2 := newDeckFromFile(cards_db)
	// cards2.print()

	cards2.shuffle()
	// cards2.print()

	// -- Structs --
	single_card := newCard("Spades", "Ace")
	fmt.Println(single_card.toString())

}
