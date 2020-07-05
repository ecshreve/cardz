package blackjack

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"

	"github.com/ecshreve/cardz/internal/deck"
)

// deal clears the player and dealer Hands and deals the first two Cards.
func deal(player, dealer *Player, d *deck.Deck) {
	player.Hand = Hand{}
	dealer.Hand = Hand{}

	c, _ := d.DealOne()
	player.Hand.addCard(*c)

	c, _ = d.DealOne()
	dealer.Hand.addCard(*c)

	c, _ = d.DealOne()
	player.Hand.addCard(*c)

	c, _ = d.DealOne()
	dealer.Hand.addCard(*c)
}

// getWinner returns the winning Player, or nil in the case of a push.
func getWinner(player, dealer *Player) *Player {
	if player.Bust {
		return dealer
	}

	if dealer.Bust {
		return player
	}

	if player.Total > dealer.Total {
		return player
	}

	if dealer.Total > player.Total {
		return dealer
	}

	return nil
}

// continueGame prompts the user to see if they want to play another hand.
func continueGame() bool {
	fmt.Println("do you want to play another hand? (y/n) followed by 'enter'")

	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}

	return char == 'y'
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
