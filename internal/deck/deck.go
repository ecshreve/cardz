package deck

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
	1:  "Ace",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
	11: "Jack",
	12: "Queen",
	13: "King",
}

// Card is a single playing card within a Deck.
type Card struct {
	Name  string
	Suit  Suit
	Code  string
	Value int
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
func (d *Deck) DealOne() (Card, error) {
	dealt := d.Cards[0]
	d.Dealt = append(d.Dealt, dealt)
	return dealt, nil

}
