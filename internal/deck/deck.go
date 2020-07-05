package deck

import (
	"fmt"

	"github.com/samsarahq/go/oops"
)

// Suit represents one of the 4 suits in a standard 52 card Deck.
type Suit int

// These are the four suits in a standard 52 card Deck, though this struct
// could probably be altered to handle different types of Cards.
const (
	Hearts Suit = iota
	Diamonds
	Clubs
	Spades
)

func (s Suit) String() string {
	suitNames := []string{"hearts", "diamonds", "clubs", "spades"}
	if int(s) > 3 {
		return "bad suit"
	}
	return suitNames[int(s)]
}

// CardValueToCodeMap maps a Card.Value to a Card's single character code.
var CardValueToCodeMap = map[int]string{
	1:  "A",
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "T",
	11: "J",
	12: "Q",
	13: "K",
}

// CardValueToNameMap maps a Card.Value to a Card's string name.
var CardValueToNameMap = map[int]string{
	1:  "ace",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "jack",
	12: "queen",
	13: "king",
}

// Card is a single playing card within a Deck.
type Card struct {
	Name string
	Suit
	Code  string
	Value int
}

func (c Card) String() string {
	return fmt.Sprintf("%s  -  %v\n", c.Name, c.Suit)
}

// Deck stores a slice of Cards remaining in the deck, as well as a slice of Cards
// that have already been removed. A Deck can have various operations performed on it.
type Deck struct {
	Cards []Card
	Dealt []Card
}

// NewDeck returns a fresh instance of a Deck.
func NewDeck() *Deck {
	d := &Deck{
		Cards: []Card{},
		Dealt: []Card{},
	}
	for _, suit := range []Suit{Hearts, Diamonds, Clubs, Spades} {
		for i := 1; i <= 13; i++ {
			c := Card{
				Name:  CardValueToNameMap[i],
				Suit:  suit,
				Code:  CardValueToCodeMap[i],
				Value: i,
			}
			d.Cards = append(d.Cards, c)
		}
	}
	return d
}

// DealOne returns the first Card in the Deck, or an error if no cards are left.
func (d *Deck) DealOne() (*Card, error) {
	if len(d.Cards) <= 0 {
		return nil, oops.Errorf("no more cards left in the deck")
	}

	dealt := d.Cards[0]
	d.Dealt = append(d.Dealt, dealt)
	d.Cards = append(d.Cards[:0], d.Cards[1:]...)

	return &dealt, nil

}
