package blackjack

import (
	"fmt"
	"strconv"

	"github.com/rivo/tview"

	"github.com/ecshreve/cardz/internal/deck"
)

var app *tview.Application
var statsArea *tview.TextView
var dealerFlex *tview.Flex
var playerFlex *tview.Flex
var hitButton *tview.Button
var stayButton *tview.Button
var actionButtons *tview.Form

func update(player, dealer *Player) {
	app.QueueUpdateDraw(func() {
		playerFlex.Clear()
		playerScore := tview.NewTextView().SetText(strconv.Itoa(player.Hand.Total)).SetTextAlign(1)
		playerScore.SetBorder(true)
		playerFlex.AddItem(playerScore, 10, 0, false)
		for _, card := range player.Hand.Cards {
			playerArea := tview.NewTextView().SetText(card.PrettyPrint()).SetTextAlign(1)
			playerFlex.AddItem(playerArea, 0, 1, false)
		}

		dealerFlex.Clear()
		dealerScore := tview.NewTextView().SetText(strconv.Itoa(dealer.Hand.Total)).SetTextAlign(1)
		dealerScore.SetBorder(true)
		dealerFlex.AddItem(dealerScore, 10, 1, false)
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

	actionButtons = tview.NewForm().SetButtonsAlign(tview.AlignCenter)
	dealerFlex = tview.NewFlex().SetDirection(tview.FlexColumn)
	playerFlex = tview.NewFlex().SetDirection(tview.FlexColumn)

	dealerFlex.SetBorder(true).SetTitle("dealer").SetBorderPadding(1, 1, 1, 1)
	playerFlex.SetBorder(true).SetTitle("player").SetBorderPadding(1, 1, 1, 1)
	actionButtons.SetHorizontal(true).SetBorder(true)
	actionButtons.SetItemPadding(3).SetBorderPadding(1, 0, 1, 1)

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(dealerFlex, 0, 2, false).
			AddItem(playerFlex, 0, 2, false).
			AddItem(actionButtons, 0, 1, false), 0, 5, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("stats"), 20, 1, false)
	flex.SetBackgroundColor(tview.Styles.PrimitiveBackgroundColor)

	go Play()
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
