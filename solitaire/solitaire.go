package main

import "fmt"
import "gogames/deck"

func SetGameBoard(gamecards *deck.Deck) *[]deck.Deck {
	gameboard := []deck.Deck{}

	slot1 := deck.NewStack()
	gameboard[0] = *slot1

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
