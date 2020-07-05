package blackjack

import (
	"github.com/ecshreve/cardz/internal/deck"
	"github.com/rivo/tview"
)

// Player represents either a human player or the CPU dealer.
type Player struct {
	Hand
	IsDealer bool
}

// dealerTurn takes care of the logic for the Dealer's hitting/standing.
func (p *Player) dealerTurn(d *deck.Deck) {
	for p.Hand.Total < 17 {
		c, _ := d.DealOne()
		p.addCard(*c)

		app.QueueUpdateDraw(func() {
			dealerFlex.Clear()
			for _, card := range p.Hand.Cards {
				dealerArea := tview.NewTextView().SetText(card.PrettyPrint()).SetTextAlign(1)
				dealerFlex.AddItem(dealerArea, 0, 1, false)
			}
		})
	}
}

// takeTurn contains the main hit/stand loop for the player.
func (p *Player) takeTurn(d *deck.Deck) {
	if p.IsDealer {
		p.dealerTurn(d)
		return
	}

	turnDone := false
	for !turnDone {
		app.QueueUpdateDraw(func() {
			actionButtons.Clear(true)
			actionButtons.AddButton("HIT", func() {
				c, _ := d.DealOne()
				p.addCard(*c)
				playerFlex.Clear()
				for _, card := range p.Hand.Cards {
					playerArea := tview.NewTextView().SetText(card.PrettyPrint()).SetTextAlign(1)
					playerFlex.AddItem(playerArea, 0, 1, false)
				}
				if p.Bust {
					turnDone = true
					return
				}
			})
			actionButtons.AddButton("STAY", func() {
				turnDone = true
				return
			})
			app.SetFocus(actionButtons).Run()
		})
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
