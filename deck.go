package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Main deck holder type
type Deck struct {
	Cards []Card `json:"cards"`
}

// Needs to return a new deck
func newDeck() Deck {
	var deck Deck

	// Intelligently create the deck using slices of all suits and values
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			tempCard := Card{
				Suit:  suit,
				Value: value,
			}
			deck.Cards = append(deck.Cards, tempCard)
		}
	}
	return deck
}

func (d Deck) printJSON() {
	deckMarsh, _ := json.MarshalIndent(d, "", "  ")
	fmt.Println(string(deckMarsh))
}

func (d Deck) toString() []string {
	var strSlice []string
	for _, v := range d.Cards {
		strSlice = append(strSlice, v.toString())
	}
	return strSlice
}

func deal(d Deck, handSize int) (Deck, Deck) {
	return Deck{Cards: d.Cards[:handSize]}, Deck{Cards: d.Cards[handSize:]}
}

// Inefficient and slower because we need to marshall and hold the whole data in memeory
func (d Deck) slowSaveToFile(filename string) error {
	byteOutput, _ := json.Marshal(d)
	return ioutil.WriteFile(filename, byteOutput, 0755)
}

// Faster as it unmarshall the file, one record at a time
func (d Deck) saveToFile(filename string) {
	file, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	defer file.Close()

	encoder := json.NewEncoder(file)
	fmt.Println("[i] Writing json file.")
	encoder.Encode(d)
}

// // func newDeckFromFile(filename string) deck {
// // 	byteStr, err := ioutil.ReadFile(filename)
// // 	if err != nil {
// // 		// Option #1 - Log the error and return a call to newDeck()
// // 		// Option #2 - Log the rror and entirely quit the program <-- CHOSEN
// // 		fmt.Println("Error: ", err)
// // 		os.Exit(1)
// // 	}
// // 	// We now have byte slice that we need to convert into a deck
// // 	s := strings.Split(string(byteStr), ",")
// // 	return deck(s)
// // }

// // Shuffle the deck
// func (d deck) shuffle() {
// 	// Generate the random seed
// 	// Use the current time to generate an int64 number based on the current time
// 	source := rand.NewSource(time.Now().UnixNano())
// 	r := rand.New(source)

// 	// Works based on indices
// 	for i := range d {
// 		newPosition := r.Intn(len(d) - 1)
// 		// Swap the element at random index with current index element
// 		d[i], d[newPosition] = d[newPosition], d[i]
// 	}
// }
