package main

import "fmt"
import "strconv"

func initSuit(suit int) string {
	switch s := suit; s {
	case 1:
		//clubs
		return "rD_"
	case 2:
		//diamonds
		return "bC_"
	case 3:
		//hearts
		return "rH_"
	case 4:
		//spades
		return "bS_"
	default:
		return "error"
	}
}

func initRoyal(j int) string {
	switch r := j; r {
	case 1:
		return "A"
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
	default:
		return strconv.Itoa(j)
	}
}

func main() {
	var deck [53]string
	var i int
	var j, suit = 1, 1
	deck[0] = "deck:\n"

	for i = 1; i <= 52; i++ {
		deck[i] = initSuit(suit) + initRoyal(j)
		if j == 13 {
			j = 1
			suit++
			println(suit)
		} else {
			j++
		}
	}

	fmt.Println("Hello World, Let's Play Solitaire!")
	fmt.Println(deck)

}
