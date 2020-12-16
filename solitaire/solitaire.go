package main

import "fmt"
import "gogames/deck"

// the gameboard is an array of stacks
// The Tableau    : seven piles that make up the main table
// The Foundation: four piles on which a whole suit or sequence must be built up
func SetGameBoard(gamecards *deck.Deck) *[]deck.Deck {
	gameboard := []deck.Deck{}

	// The Tableau
	slot0 := deck.NewStack()
	slot1 := deck.NewStack()
	slot2 := deck.NewStack()
	slot3 := deck.NewStack()
	slot4 := deck.NewStack()
	slot5 := deck.NewStack()
	slot6 := deck.NewStack()
	// The Foundation
	slot7 := deck.NewStack()
	slot8 := deck.NewStack()
	slot9 := deck.NewStack()
	slot10 := deck.NewStack()

	// append stacks to gameboard
	gameboard = append(gameboard, *slot0, *slot1, *slot2, *slot3, *slot4, *slot5, *slot6, *slot7, *slot8, *slot9, *slot10)

	return &gameboard
}


func main() {
	// create deck
	gamecards := deck.NewDeck()

	// opening message
	fmt.Println("Hello World, Let's Play Solitaire!")

	// setup game
	deck.ShuffleDeck(gamecards)
	gameboard := SetGameBoard(gamecards)
	fmt.Println(gameboard)
}
