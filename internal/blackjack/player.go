package blackjack

// Player represents either a human player or the CPU dealer.
type Player struct {
	Hand
	IsDealer bool
}

// takeTurn contains the main hit/stay loop for the player and dealer.
func (bg *BlackjackGame) takeTurn() {
	// Do the human player's turn.
	bg.PlayerTurn = true
	for bg.PlayerTurn {
		if bg.Player.Bust {
			bg.PlayerTurn = false
			return
		}
	}

	// Do the dealer's turn.
	for bg.Dealer.Hand.Total < 17 {
		c, _ := bg.Deck.DealOne()
		bg.Dealer.addCard(*c)
		bg.HasUpdate = true
	}
}

// handleHandEnd prints a game over message and checks if the player wants to continue.
func (p Player) handleHandEnd() bool {
	// retStr := "yay you won!"
	// if p.IsDealer {
	// 	retStr = "better luck next time!"
	// }
	// fmt.Println(retStr)
	return continueGame()
}
