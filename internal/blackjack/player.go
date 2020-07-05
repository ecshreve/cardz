package blackjack

// Player represents either a human player or the CPU dealer.
type Player struct {
	Hand
	IsDealer bool
}
