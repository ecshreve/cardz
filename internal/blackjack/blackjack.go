package blackjack

import (
	"fmt"

	"github.com/rivo/tview"

	"github.com/ecshreve/cardz/internal/deck"
)

var app *tview.Application
var statsArea *tview.TextView
var dealerFlex *tview.Flex
var playerFlex *tview.Flex

func update(player, dealer *Player) {
	app.QueueUpdateDraw(func() {
		playerFlex.Clear()
		for _, card := range player.Hand.Cards {
			playerArea := tview.NewTextView().SetText(card.PrettyPrint()).SetTextAlign(1)
			playerFlex.AddItem(playerArea, 0, 1, false)
		}

		dealerFlex.Clear()
		for _, card := range dealer.Hand.Cards {
			dealerArea := tview.NewTextView().SetText(card.PrettyPrint()).SetTextAlign(1)
			dealerFlex.AddItem(dealerArea, 0, 1, false)
		}
	})
}

// Play contains the main game loop for this silly blackjack game.
func Play() {
	d := deck.NewDeck()
	d.ShuffleRemaining()

	player := &Player{IsDealer: false}
	dealer := &Player{IsDealer: true}
	continueGame := true

	// Loop until the player decides they don't want to play anymore.
	for continueGame {
		if len(d.Cards) < 26 {
			d.ReShuffle()
		}

		deal(player, dealer, d)
		update(player, dealer)

		player.takeTurn(d)
		if !player.Bust {
			dealer.takeTurn(d)
		}

		winner := getWinner(player, dealer)
		continueGame = winner.handleHandEnd()
	}

	fmt.Println("thanks for playing!")
}

func StartGame() {
	app = tview.NewApplication()
	tview.Styles = customCliTheme

	statsArea = tview.NewTextView().SetText("test").SetTextAlign(1)
	statsArea.SetBorder(true).SetTitle("stats").SetBorderPadding(0, 1, 1, 1)
	dealerFlex = tview.NewFlex().SetDirection(tview.FlexColumn)
	playerFlex = tview.NewFlex().SetDirection(tview.FlexColumn)
	dealerFlex.SetBorder(true).SetTitle("dealer").SetBorderPadding(0, 1, 1, 1)
	playerFlex.SetBorder(true).SetTitle("player").SetBorderPadding(1, 1, 1, 1)

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(statsArea, 0, 1, false).
			AddItem(dealerFlex, 0, 2, false).
			AddItem(playerFlex, 0, 2, false), 0, 5, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("History"), 20, 1, false)
	flex.SetBackgroundColor(tview.Styles.PrimitiveBackgroundColor)

	go Play()
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
