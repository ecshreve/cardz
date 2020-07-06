package blackjack

import (
	"fmt"

	"github.com/ecshreve/cardz/internal/deck"
)

// Hand store the Player's Cards and some other useful data.
type Hand struct {
	Cards     []deck.Card
	Total     int
	NumAces   int
	Blackjack bool
	Bust      bool
}

// Implement the Stringer interface for the Hand type.
func (h Hand) String() string {
	return fmt.Sprintf("total: %d -- bust: %v\ncards:\n%+v\n", h.Total, h.Bust, h.Cards)
}

// addCard adds a Card to the Player's hand and recalculates the Hand total, also
// updates the Bust field depending on the Total and checks the Hand for a blackjack.
func (h *Hand) addCard(c deck.Card) {
	h.Cards = append(h.Cards, c)

	// Calculate the max value of the Hand, and keep track of num Aces in Hand.
	h.Total = 0
	for _, card := range h.Cards {
		if card.Code == "A" {
			h.NumAces++
		}
		h.Total += card.Value
	}

	// Check for blackjack.
	if len(h.Cards) == 2 && h.Total == 21 {
		h.Blackjack = true
		return
	}

	// Take the Aces into account if there are any.
	for h.NumAces > 0 {
		if h.Total > 21 {
			h.Total -= 10
			h.NumAces--
		} else {
			break
		}
	}

	// Set the Hand's Bust value.
	h.Bust = h.Total > 21
}
