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
	fmt.Println("Deck from Hand-Draw: \n", hand)
	fmt.Println("Remainder from deck: \n", remainingCards)

	cards_db := "cards_db.json"

	cards.streamSaveToFile(cards_db)

	cards2 := streamNewDeckFromFile(cards_db)
	fmt.Println("Deck from <streamNewDeckFromFile>: \n", cards2)

	cards2.shuffle()
	fmt.Println("Deck after <shuffle>: \n", cards2)
}
