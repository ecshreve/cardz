package blackjack

import (
	"strconv"
	"strings"
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
var yesButton *tview.Button
var noButton *tview.Button

var continuePlaying = make(chan bool)
var dealerTurn = make(chan bool)

type BlackjackGame struct {
	Player          *Player
	Dealer          *Player
	Deck            *deck.Deck
	HasUpdate       bool
	PlayerTurn      bool
	HandComplete    bool
	ContinuePlaying bool
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
						dealerTurn <- true
					}
				})

				stayButton = tview.NewButton("STAY")
				stayButton.SetBackgroundColorActivated(tcell.ColorGreen)
				stayButton.SetBorder(true)
				stayButton.SetSelectedFunc(func() {
					bg.PlayerTurn = false
					bg.HasUpdate = true
					dealerTurn <- true
				})

				statusFlex.AddItem(hitButton, 0, 1, false)
				statusFlex.AddItem(stayButton, 0, 1, true)
				app.SetFocus(statusFlex)
			}

			if bg.HandComplete {
				winner := bg.getWinner()

				verb := "won"
				if winner == nil {
					verb = "tied"
				} else if winner.IsDealer {
					verb = "lost"
				}

				resultStr := "you " + strings.ToUpper(verb) + " do you want to play another hand?"
				gameOver := tview.NewTextView().SetText(resultStr).SetTextAlign(1)

				yesButton = tview.NewButton("yes")
				yesButton.SetBackgroundColorActivated(tcell.ColorGreen)
				yesButton.SetBorder(true)
				yesButton.SetSelectedFunc(func() {
					bg.HandComplete = false
					bg.HasUpdate = true
					continuePlaying <- true
				})

				noButton = tview.NewButton("no")
				noButton.SetBackgroundColorActivated(tcell.ColorRed)
				noButton.SetBorder(true)
				noButton.SetSelectedFunc(func() {
					continuePlaying <- false
					app.Stop()
				})

				statusFlex.Clear()
				statusFlex.AddItem(gameOver, 0, 2, false)
				statusFlex.AddItem(yesButton, 0, 1, true)
				statusFlex.AddItem(noButton, 0, 1, false)
				app.SetFocus(statusFlex)
			}
		})
		bg.HasUpdate = false
	}
}

// Play contains the main game loop for this silly blackjack game.
func Play(bg *BlackjackGame) {
	bg.Deck.ShuffleRemaining()

	for bg.ContinuePlaying {
		if len(bg.Deck.Cards) < 26 {
			bg.Deck.ReShuffle()
		}

		bg.deal()
		bg.HasUpdate = true

		bg.PlayerTurn = true

		<-dealerTurn
		for bg.Dealer.Hand.Total < 17 {
			c, _ := bg.Deck.DealOne()
			bg.Dealer.addCard(*c)
			bg.HasUpdate = true
		}

		bg.HandComplete = true
		bg.HasUpdate = true
		bg.ContinuePlaying = <-continuePlaying
	}
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
		Deck:            deck.NewDeck(),
		HasUpdate:       true,
		PlayerTurn:      false,
		HandComplete:    false,
		ContinuePlaying: true,
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
			case yesButton.HasFocus():
				app.SetFocus(noButton)
			case noButton.HasFocus():
				app.SetFocus(yesButton)
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
