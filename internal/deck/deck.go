package deck

import (
	"fmt"
	"math/rand"
	"time"

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

var symbolMap = map[Suit]string{
	Hearts:   "♥",
	Diamonds: "♦",
	Clubs:    "♣",
	Spades:   "♠",
}

func (s Suit) String() string {
	suitNames := []string{"hearts", "diamonds", "clubs", "spades"}
	if int(s) > 3 {
		return "bad suit"
	}
	return suitNames[int(s)]
}

// CardCodes are all the possible single char codes.
var CardCodes = []string{
	"A",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"T",
	"J",
	"Q",
	"K",
}

// CardCodeToValueMap maps a Card.Code to a Card's Value.
var CardCodeToValueMap = map[string]int{
	"A": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 10,
	"Q": 10,
	"K": 10,
}

// CardCodeToNameMap maps a Card.code to a Card's string name.
var CardCodeToNameMap = map[string]string{
	"A": "ace",
	"2": "two",
	"3": "three",
	"4": "four",
	"5": "five",
	"6": "six",
	"7": "seven",
	"8": "eight",
	"9": "nine",
	"T": "ten",
	"J": "jack",
	"Q": "queen",
	"K": "king",
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

// PrettyPrint returns a string form of the card that looks like this:
// ┌─────────┐
// │ A       │
// │         │
// │    ♠    │
// │         │
// │       A │
// └─────────┘
func (c Card) PrettyPrint() string {
	r1 := "┌─────────┐\n"
	r2 := fmt.Sprintf("│ %s       │\n", c.Code)
	r3 := "│         │\n"
	r4 := fmt.Sprintf("│    %s    │\n", symbolMap[c.Suit])
	r5 := "│         │\n"
	r6 := fmt.Sprintf("│       %s │\n", c.Code)
	r7 := "└─────────┘"

	// Special handling for the 10 because it's the only one that's 2 chars.
	if c.Code == "T" {
		r2 = "│ 10      │\n"
		r6 = "│      10 │\n"
	}
	return r1 + r2 + r3 + r4 + r5 + r6 + r7
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
		for _, code := range CardCodes {
			c := Card{
				Name:  CardCodeToNameMap[code],
				Suit:  suit,
				Code:  code,
				Value: CardCodeToValueMap[code],
			}
			fmt.Println(c.PrettyPrint())
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

// DealMany returns the requested number of Cards, or an error if there aren't
// enough in the Deck.
func (d *Deck) DealMany(numToDeal int) ([]Card, error) {
	if len(d.Cards)-numToDeal < 0 {
		return nil, oops.Errorf("not enough cards in the deck, deck: %d - requested: %d", len(d.Cards), numToDeal)
	}

	cards := d.Cards[:numToDeal]
	d.Dealt = append(d.Dealt, cards...)
	d.Cards = d.Cards[numToDeal:]

	return cards, nil
}

// ShuffleRemaining is pretty self-explanatory, it shuffles the remaining Cars in
// the Deck.
func (d *Deck) ShuffleRemaining() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}

// ReShuffle puts the Dealt cards back into the Cards slice and shuffles all of
// them together.
func (d *Deck) ReShuffle() {
	d.Cards = append(d.Cards, d.Dealt...)
	d.Dealt = []Card{}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}
