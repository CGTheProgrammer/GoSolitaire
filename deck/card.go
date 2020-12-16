package deck

// Card definition
type Card struct {
  Suit string
  Number int
  Facing bool
}

// creates new card type
func NewCard(suit string, number int) *Card {

  c := Card{Suit: suit, Number: number}
  return &c
}
