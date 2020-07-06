package blackjack

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/samsarahq/go/oops"
)

// deal clears the player and dealer Hands, resets the Game HandComplete
// field, and deals the first two Cards if this Hand.
func (bg *Game) deal() error {
	bg.Player.Hand = Hand{}
	bg.Dealer.Hand = Hand{}
	bg.HandComplete = false

	c, err := bg.Deck.DealOne()
	if err != nil {
		return oops.Wrapf(err, "error dealing initial hands")
	}
	bg.Player.Hand.addCard(*c)

	c, err = bg.Deck.DealOne()
	if err != nil {
		return oops.Wrapf(err, "error dealing initial hands")
	}
	bg.Dealer.Hand.addCard(*c)

	c, err = bg.Deck.DealOne()
	if err != nil {
		return oops.Wrapf(err, "error dealing initial hands")
	}
	bg.Player.Hand.addCard(*c)

	c, err = bg.Deck.DealOne()
	if err != nil {
		return oops.Wrapf(err, "error dealing initial hands")
	}
	bg.Dealer.Hand.addCard(*c)

	return nil
}

// getWinner returns the winning Player, or nil in the case of a push.
func (bg Game) getWinner() *Player {
	var winner *Player
	switch {
	case bg.Player.Blackjack:
		winner = bg.Player
	case bg.Player.Bust:
		winner = bg.Dealer
	case bg.Dealer.Bust:
		winner = bg.Player
	case bg.Player.Total > bg.Dealer.Total:
		winner = bg.Player
	case bg.Dealer.Total > bg.Player.Total:
		winner = bg.Dealer
	default:
		winner = nil
	}

	bg.Winner = winner
	bg.UpdateStats()
	bg.Stats.Save()
	return winner
}

// customCliTheme holds the color theme for the tview TUI.
var customCliTheme = tview.Theme{
	PrimitiveBackgroundColor:    tcell.Color(272727),
	ContrastBackgroundColor:     tcell.Color(448488),
	MoreContrastBackgroundColor: tcell.ColorGreen,
	BorderColor:                 tcell.ColorWhite,
	TitleColor:                  tcell.ColorWhite,
	GraphicsColor:               tcell.ColorWhite,
	PrimaryTextColor:            tcell.ColorWhite,
	SecondaryTextColor:          tcell.ColorYellow,
	TertiaryTextColor:           tcell.ColorGreen,
	InverseTextColor:            tcell.Color(448488),
	ContrastSecondaryTextColor:  tcell.ColorDarkCyan,
}
