package deck

import "math/rand"

// Deck = array of Cards
type Deck struct {
  Cards []Card
  }

// creates new Deck type of standard 52 playing cards
func NewDeck() *Deck {
  deck := Deck{}
  suits := []string{"spades", "hearts", "diamonds", "clubs"}
  numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13}

  for _, suit := range suits {
    for _, number := range numbers {
      deck.Cards = append(deck.Cards,
        Card{Suit:suit, Number:number, Facingup:false})
    }
  }

  return &deck
}

// randomizes order of Cards in Deck
//   by swapping one Card's position with another for each card place x100
func ShuffleDeck(deck *Deck) *Deck{
  cards := deck.Cards

  for z := 1; z <=100; z++ {
    for i := range cards {
      j := rand.Intn(i + 1)
      cards[i], cards[j] = cards[j], cards[i]
    }
  }

  return deck
}

// stacks function as decks without the cards being defined
func NewStack() *Deck {
  stack := Deck{}
  return &stack
}
