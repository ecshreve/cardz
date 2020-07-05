package blackjack

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"

	"github.com/ecshreve/cardz/internal/deck"
)

const refreshInterval = 500 * time.Millisecond

var app *tview.Application
var dealerFlex *tview.Flex
var playerFlex *tview.Flex
var statusFlex *tview.Flex
var hitButton *tview.Button
var stayButton *tview.Button

type BlackjackGame struct {
	Player     *Player
	Dealer     *Player
	Deck       *deck.Deck
	PlayerTurn bool
	HasUpdate  bool
}

func update(bg *BlackjackGame) {
	for {
		time.Sleep(refreshInterval)
		if !bg.HasUpdate {
			continue
		}
		app.QueueUpdateDraw(func() {
			playerFlex.Clear()
			playerScore := tview.NewTextView().SetText(strconv.Itoa(bg.Player.Hand.Total)).SetTextAlign(1)
			playerScore.SetBorder(true)
			playerFlex.AddItem(playerScore, 10, 0, false)
			for _, card := range bg.Player.Hand.Cards {
				playerArea := tview.NewTextView().SetText(card.PrettyPrint()).SetTextAlign(1)
				playerFlex.AddItem(playerArea, 0, 1, false)
			}

			dealerFlex.Clear()
			dealerScore := tview.NewTextView().SetText(strconv.Itoa(bg.Dealer.Hand.Total)).SetTextAlign(1)
			dealerScore.SetBorder(true)
			dealerFlex.AddItem(dealerScore, 10, 1, false)
			for _, card := range bg.Dealer.Hand.Cards {
				dealerArea := tview.NewTextView().SetText(card.PrettyPrint()).SetTextAlign(1)
				dealerFlex.AddItem(dealerArea, 0, 1, false)
			}

			if bg.PlayerTurn {
				statusFlex.Clear()
				hitButton = tview.NewButton("HIT")
				hitButton.SetBackgroundColorActivated(tcell.ColorRed)
				hitButton.SetBorder(true)
				hitButton.SetSelectedFunc(func() {
					c, _ := bg.Deck.DealOne()
					bg.Player.addCard(*c)
					bg.HasUpdate = true
					if bg.Player.Bust {
						bg.PlayerTurn = false
					}
				})

				stayButton = tview.NewButton("STAY")
				stayButton.SetBackgroundColorActivated(tcell.ColorGreen)
				stayButton.SetBorder(true)
				stayButton.SetSelectedFunc(func() {
					bg.PlayerTurn = false
					bg.HasUpdate = true
				})

				statusFlex.AddItem(hitButton, 0, 1, false)
				statusFlex.AddItem(stayButton, 0, 1, true)
				app.SetFocus(statusFlex)
			}
		})
		bg.HasUpdate = false
	}
}

// Play contains the main game loop for this silly blackjack game.
func Play(bg *BlackjackGame) {
	bg.Deck.ShuffleRemaining()

	continueGame := true

	// Loop until the player decides they don't want to play anymore.
	for continueGame {
		if len(bg.Deck.Cards) < 26 {
			bg.Deck.ReShuffle()
		}

		bg.deal()
		bg.HasUpdate = true

		bg.takeTurn()

		winner := getWinner(bg.Player, bg.Dealer)
		continueGame = winner.handleHandEnd()
	}

	fmt.Println("thanks for playing!")
}

func StartGame() {
	bg := &BlackjackGame{
		Player: &Player{
			Hand:     Hand{},
			IsDealer: false,
		},
		Dealer: &Player{
			Hand:     Hand{},
			IsDealer: true,
		},
		Deck: deck.NewDeck(),
	}

	app = tview.NewApplication()
	tview.Styles = customCliTheme

	dealerFlex = tview.NewFlex().SetDirection(tview.FlexColumn)
	playerFlex = tview.NewFlex().SetDirection(tview.FlexColumn)
	statusFlex = tview.NewFlex().SetDirection(tview.FlexColumn)

	dealerFlex.SetBorder(true).SetTitle("dealer").SetBorderPadding(1, 1, 1, 1)
	playerFlex.SetBorder(true).SetTitle("player").SetBorderPadding(1, 1, 1, 1)
	statusFlex.SetBorder(true)

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(dealerFlex, 0, 2, false).
			AddItem(playerFlex, 0, 2, false).
			AddItem(statusFlex, 0, 1, false), 0, 5, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("stats"), 20, 1, false)
	flex.SetBackgroundColor(tview.Styles.PrimitiveBackgroundColor)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTAB {
			switch {
			case hitButton.HasFocus():
				app.SetFocus(stayButton)
			case stayButton.HasFocus():
				app.SetFocus(hitButton)
			}
		}
		return event
	})

	go update(bg)
	go Play(bg)
	if err := app.SetRoot(flex, true).SetFocus(statusFlex).Run(); err != nil {
		panic(err)
	}

}
