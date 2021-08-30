package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
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
	fmt.Println("--[i] deal --")
	return Deck{Cards: d.Cards[:handSize]}, Deck{Cards: d.Cards[handSize:]}
}

// Inefficient and slower because we need to marshall and hold the whole data in memeory
func (d Deck) byteSaveToFile(filename string) error {
	fmt.Println("--[i] byteSaveToFile --")
	fmt.Println("[i] Marshal data: ")
	byteOutput, _ := json.Marshal(d)
	fmt.Println("[i] Writing file: ", filename)
	return ioutil.WriteFile(filename, byteOutput, 0755)
}

// Faster as it unmarshall the file, one record at a time
func (d Deck) streamSaveToFile(filename string) {
	fmt.Println("--[i] streamSaveToFile --")
	fmt.Println("[i] Opening file: ", filename)
	file, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	defer file.Close()

	// Typically json.Encoder is use from stream and json.Marshal is used for in-memory objects
	encoder := json.NewEncoder(file)
	fmt.Println("[i] Writing json file.")
	encoder.Encode(d)
}

func byteNewDeckFromFile(filename string) Deck {
	fmt.Println("--[i] byteNewDeckFromFile --")
	fmt.Println("[i] Opening file: ", filename)
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		// Option #1 - Log the error and return a call to newDeck()
		// Option #2 - Log the rror and entirely quit the program <-- CHOSEN
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println("[i] Reading from file.")
	byteValue, _ := ioutil.ReadAll(file)

	// Create the var to hold the unmarshalled data
	var deck Deck

	fmt.Println("[i] Unmarshalling data.")
	err = json.Unmarshal(byteValue, &deck)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return deck
}

func streamNewDeckFromFile(filename string) Deck {
	fmt.Println("--[i] streamNewDeckFromFile --")
	fmt.Println("[i] Opening file: ", filename)
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		// Option #1 - Log the error and return a call to newDeck()
		// Option #2 - Log the rror and entirely quit the program <-- CHOSEN
		fmt.Println("[Error] : ", err)
		os.Exit(1)
	}

	decoder := json.NewDecoder(file)

	// Set the structure of the data
	var tempDeck Deck
	// Use - m := map[string]interface{} - when not sure of underlying structure

	err = decoder.Decode(&tempDeck)
	if err != nil {
		fmt.Println("[Error] : ", err)
	}

	return tempDeck
}

// Shuffle the deck
func (d Deck) shuffle() {
	fmt.Println("--[i] shuffle --")
	// Generate the random seed
	// Use the current time to generate an int64 number based on the current time
	fmt.Println("[i] Setting random seed.")
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	fmt.Println("[i] Iterate through cards and shuffle each.")
	// Works based on indices
	for i := range d.Cards {
		newPosition := r.Intn(len(d.Cards) - 1)
		// Swap the element at random index with current index element
		d.Cards[i], d.Cards[newPosition] = d.Cards[newPosition], d.Cards[i]
	}
}
