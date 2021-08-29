package main

// card type
type card struct {
	suit  string
	value string
}

func newCard(suit string, value string) card {
	tempCard := card{
		suit:  suit,
		value: value,
	}
	return tempCard
}

func (c card) toString() string {
	return c.value + " of " + c.suit
}
