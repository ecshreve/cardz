package blackjack

import (
	"fmt"

	"github.com/ecshreve/cardz/internal/deck"
)

// Hand store the Player's Cards and some other useful data.
type Hand struct {
	Cards []deck.Card
	Total int
	//HasAce    bool
	//Blackjack bool
	Bust bool
}

// Implement the Stringer interface for the Hand type.
func (h Hand) String() string {
	return fmt.Sprintf("total: %d -- bust: %v\ncards:\n%+v\n", h.Total, h.Bust, h.Cards)
}

// addCard adds a Card to the Player's hand and recalculates the Hand total, also
// updates the Bust field depending on the Total.
func (h *Hand) addCard(c deck.Card) {
	h.Cards = append(h.Cards, c)
	h.Total += c.Value
	h.Bust = h.Total > 21
}
