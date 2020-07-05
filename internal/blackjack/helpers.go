package blackjack

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// deal clears the player and dealer Hands and deals the first two Cards.
func (bg *BlackjackGame) deal() {
	bg.Player.Hand = Hand{}
	bg.Dealer.Hand = Hand{}

	c, _ := bg.Deck.DealOne()
	bg.Player.Hand.addCard(*c)

	c, _ = bg.Deck.DealOne()
	bg.Dealer.Hand.addCard(*c)

	c, _ = bg.Deck.DealOne()
	bg.Player.Hand.addCard(*c)

	c, _ = bg.Deck.DealOne()
	bg.Dealer.Hand.addCard(*c)
}

// getWinner returns the winning Player, or nil in the case of a push.
func (bg BlackjackGame) getWinner() *Player {
	if bg.Player.Bust {
		return bg.Dealer
	}

	if bg.Dealer.Bust {
		return bg.Player
	}

	if bg.Player.Total > bg.Dealer.Total {
		return bg.Player
	}

	if bg.Dealer.Total > bg.Player.Total {
		return bg.Dealer
	}

	return nil
}

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
