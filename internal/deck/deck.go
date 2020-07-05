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

// Card is a single playing card within a Deck.
type Card struct {
	Name  string
	Suit  Suit
	Char  string
	Value int
}

// Deck stores a slice of Cards remaining in the deck, as well as a slice of Cards
// that have already been removed. A Deck can have various operations performed on it.
type Deck struct {
	Cards []Card
	Dealt []Card
}
