package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Main deck holder type
type deck []string

// Needs to return a new deck
func newDeck() deck {
	cards := deck{}

	// Intelligently create the deck using slices of all suits and values
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// convert deck to a string
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteStr, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - Log the error and return a call to newDeck()
		// Option #2 - Log the rror and entirely quit the program <-- CHOSEN
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// We now have byte slice that we need to convert into a deck
	s := strings.Split(string(byteStr), ",")
	return deck(s)
}

// Shuffle the deck
func (d deck) shuffle() {
	// Generate the random seed
	// Use the current time to generate an int64 number based on the current time
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Works based on indices
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// Swap the element at random index with current index element
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
