package main

import "fmt"

// card type
type Card struct {
	Suit  string `json:"suit"`
	Value string `json:"value"`
}

func newCard(suit string, value string) Card {
	tempCard := Card{
		Suit:  suit,
		Value: value,
	}
	return tempCard
}

func (c Card) String() string {
	return fmt.Sprintf("%v-%v", c.Value, c.Suit)
}

func (c Card) toString() string {
	return fmt.Sprintf("%v-%v", c.Value, c.Suit)
}
