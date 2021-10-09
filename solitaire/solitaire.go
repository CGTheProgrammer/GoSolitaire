package main

import "fmt"
import "gogames/deck"

// the gameboard is an array of stacks
// The Deck       : the
// The Tableau    : seven piles that make up the main table
// The Foundation : four piles on which a whole suit or sequence must be built up
// The Play Slot  : the cards from the deck that are playable
func SetGameBoard() *[]deck.Deck {
	gameboard := []deck.Deck{}

	// The Deck
	slot0 := deck.NewDeck()
	// The Tableau
	slot1 := deck.NewStack()
	slot2 := deck.NewStack()
	slot3 := deck.NewStack()
	slot4 := deck.NewStack()
	slot5 := deck.NewStack()
	slot6 := deck.NewStack()
	slot7 := deck.NewStack()
	// The Foundation
	slot8 := deck.NewStack()
	slot9 := deck.NewStack()
	slot10 := deck.NewStack()
	slot11 := deck.NewStack()
	// Play Slot
	slot12 := deck.NewStack()

	// append stacks to gameboard
	gameboard = append(gameboard,
		*slot0,
		*slot1,
		*slot2,
		*slot3,
		*slot4,
		*slot5,
		*slot6,
		*slot7,
		*slot8,
		*slot9,
		*slot10,
		*slot11,
		*slot12)

	return &gameboard
}


func main() {
	// opening message
	fmt.Println("Hello World, Let's Play Solitaire!")

	// setup game
	gameboard := SetGameBoard()
	fmt.Println(gameboard)
}
