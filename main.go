package main

import "fmt"

func main() {
	cards := newDeck()

	// Different printing varieties
	// fmt.Println(cards)
	// cards.printJSON()
	// fmt.Println(cards.toString())

	// Pass by value/copy, the original still remains intact
	hand, remainingCards := deal(cards, 5)
	fmt.Println(hand, "\n\n", remainingCards)

	cards_db := "cards_db.json"

	cards.saveToFile(cards_db)

	// cards2 := newDeckFromFile(cards_db)
	// // cards2.print()

	// cards2.shuffle()
	// cards2.print()

	// -- Structs --
	// single_card := newCard("Spades", "Ace")
	// fmt.Println(single_card)

}
